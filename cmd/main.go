package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"canty/config"
	"canty/internal/core/entities"
	"canty/internal/core/services"
	"canty/internal/infrastructures/monitoring"
	"canty/internal/modules/ai_video"
	"canty/internal/modules/analysis"
	"canty/internal/modules/audio"
	"canty/internal/modules/downloader"
	processor "canty/internal/modules/processing"
	"canty/internal/modules/seo"
	"canty/internal/modules/uploader"

	"gopkg.in/yaml.v3"
)

// processedVideo используется для передачи обработанного видео с информацией о владельце (username)
type processedVideo struct {
	username string
	video    *entities.Video
}

type videoProcessingParams struct {
	analysisService *analysis.VideoAnalysisService
	downloader      downloader.VideoDownloader
	processor       processor.VideoProcessor
	audioGenerator  *audio.AudioGenerator
	seoGenerator    seo.SeoGenerator
	videoService    *services.VideoService
	videoGenerator  *ai_video.VideoGenerator
	config          config.Config
}

func main() {
	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)

	// Загрузка конфигурации
	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	var cfg config.Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Ошибка парсинга YAML: %v", err)
	}
	fmt.Printf("config: %+v\n", cfg.YtAccounts[0].Credentials)

	//// Инициализация базы данных через фабрику
	//dbFactory := &databases.DatabaseFactory{}
	//database, err := dbFactory.CreateDatabase(cfg.DBConfig.Type, cfg)
	//if err != nil {
	//	log.Fatalf("Error creating database: %v", err)
	//}
	//log.Println("Database connected successfully")
	//
	//var videoRepo entities.VideoRepository
	//switch cfg.DBConfig.Type {
	//case "mongo":
	//	cl, err := database.Connect()
	//	if err != nil {
	//		log.Fatalf("Error connecting to MongoDB: %v", err)
	//	}
	//	mongoClient, ok := cl.(*mongo.Client)
	//	if !ok {
	//		log.Fatalf("Failed to cast to *mongo.Client")
	//	}
	//	videoRepo = repoMongo.NewMongoVideoRepository(mongoClient, cfg.DBConfig.Mongo.DBName, "videos")
	//case "postgres":
	//	_, err := database.Connect()
	//	if err != nil {
	//		log.Fatalf("Error connecting to PostgreSQL: %v", err)
	//	}
	//	videoRepo = postgresql.NewPostgresVideoRepository(nil)//todo
	//default:
	//	log.Fatalf("Unsupported database type for repositories: %+v", cfg.DBConfig)
	//}
	log.Println("Initialized repositories successfully")

	// Инициализация сервисов и модулей
	videoService := services.NewVideoService(nil)
	videoUploader, err := uploader.NewVideoUploader(ctx, cfg) // внутри него создаются YClients по AccountConfig
	if err != nil {
		log.Fatalf("Error initialazing videoUploader: %v", err)
	}

	videoAnalysisService := analysis.NewVideoAnalysisService(videoUploader.YClients, cfg) // future receive -> tiktok,etc
	videoDownloader := downloader.VideoDownloader{}
	videoProcessor := processor.VideoProcessor{}
	audioGenerator, err := audio.NewAudioGenerator(ctx)
	if err != nil {
		log.Fatalf("Error creating audio generator: %v", err)
	}
	seoGenerator := seo.SeoGenerator{}
	videoGenerator := ai_video.NewVideoGenerator(cfg.App.VeedAPIKey)

	params := &videoProcessingParams{
		analysisService: videoAnalysisService,
		downloader:      videoDownloader,
		processor:       videoProcessor,
		audioGenerator:  audioGenerator,
		seoGenerator:    seoGenerator,
		videoService:    videoService,
		videoGenerator:  videoGenerator,
		config:          cfg,
	}

	// Запуск мониторинга
	monitoring.StartPerformanceMonitoring()
	go monitoring.StartPrometheusMetrics()

	// Обработка видео – результаты собираем в канал processedVideo
	processedVideosCh := make(chan processedVideo, 100)
	var wg sync.WaitGroup
	processVideos(&wg, params, processedVideosCh)
	go func() {
		wg.Wait()
		close(processedVideosCh)
	}()

	// Агрегация: создаём мапу username -> *entities.Video
	processedVideosMap := make(map[string]*entities.Video)
	for pv := range processedVideosCh {
		processedVideosMap[pv.username] = pv.video
	}

	// Загрузка видео: для каждого username получаем соответствующий YouTube-клиент и загружаем видео
	uploadVideos(videoUploader, processedVideosMap)

	// Завершаем работу приложения
	shutDown(ctx)
}

// processVideos запускает параллельную обработку для платформ (например, YouTube и TikTok)
func processVideos(wg *sync.WaitGroup, params *videoProcessingParams, out chan<- processedVideo) {
	wg.Add(2)
	go analyzeAndProcessVideos(wg, "youtube", params, out)
	go analyzeAndProcessVideos(wg, "tiktok", params, out)
}

// analyzeAndProcessVideos обрабатывает видео для указанной платформы.
// Для каждого аккаунта (username) выбирается первое популярное видео и обрабатывается.
func analyzeAndProcessVideos(wg *sync.WaitGroup, platform string, params *videoProcessingParams, out chan<- processedVideo) {
	defer wg.Done()

	videosByAccount, err := params.analysisService.GetPopularVideos(platform)
	if err != nil {
		log.Printf("Error getting popular videos for %s: %v", platform, err)
		return
	}

	fmt.Println("GetPopularVideos count", len(videosByAccount))

	for username, videos := range videosByAccount {
		if len(videos) == 0 {
			continue
		}

		// Выбираем первое видео для данного аккаунта
		video := videos[0]

		// Скачивание видео
		downloadedVideo, err := params.downloader.DownloadVideo(platform, video.URL, params.config.App.DownloadPath)
		if err != nil {
			log.Printf("Error downloading video for user %s: %v", username, err)
			continue
		}
		video.Content = downloadedVideo.Content
		video.FilePath = downloadedVideo.FilePath

		// Генерация аудио
		var audioBt []byte

		if video.Subtitles != "" {
			audioBt, err = params.audioGenerator.GenerateAudio(context.Background(), video.Subtitles, params.config.App.OutputPath+"/audio.mp3")
			if err != nil {
				log.Printf("Error generating audio for user %s: %v", username, err)
			}
		}

		// Обработка видео (например, наложение аудио, текста и т.д.)
		newGeneratedVideo, err := params.processor.ProcessVideo(&video, params.config.App.OutputPath, video.Subtitles, audioBt)
		if err != nil {
			log.Printf("Error processing video for user %s: %v", username, err)
			continue
		}

		// Генерация описания и хэштегов
		newGeneratedVideo.Description = params.seoGenerator.GenerateDescription(video.Description)
		newGeneratedVideo.Tags = params.seoGenerator.GenerateHashtags(video.Tags)

		// Отправляем результат в канал вместе с username
		out <- processedVideo{username: username, video: newGeneratedVideo}
	}
}

// uploadVideos принимает мапу с видео, где ключ – username, и для каждого пользователя использует соответствующий YouTube-клиент для загрузки видео.
func uploadVideos(uploader *uploader.VideoUploader, processedVideos map[string]*entities.Video) {
	var wg sync.WaitGroup
	for username, video := range processedVideos {
		yClient, ok := uploader.YClients[username]
		if !ok {
			log.Printf("No YouTube client for username %s", username)
			continue
		}
		wg.Add(1)
		go func(username string, video *entities.Video, yClient *entities.YClient) {
			defer wg.Done()
			if err := uploader.UploadToYouTube(video); err != nil {
				log.Printf("Error uploading video to YouTube: %v", err)
			}

		}(username, video, yClient)
	}
	wg.Wait()
}

func setupGracefulShutdown(cancel context.CancelFunc) {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		fmt.Println("Received shutdown signal")
		cancel()
	}()
}

func shutDown(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Shutting down gracefully")
		os.Exit(0)
	case <-time.After(10 * time.Second):
		fmt.Println("Shutdown timeout, exiting")
		os.Exit(1)
	}
}
