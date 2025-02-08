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
	"canty/internal/infrastructures/databases"
	repoMongo "canty/internal/infrastructures/databases/mongo"
	"canty/internal/infrastructures/databases/postgresql"
	"canty/internal/infrastructures/monitoring"
	"canty/internal/modules/ai_video"
	"canty/internal/modules/analysis"
	"canty/internal/modules/audio"
	"canty/internal/modules/downloader"
	processor "canty/internal/modules/processing"
	"canty/internal/modules/seo"
	"canty/internal/modules/uploader"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/yaml.v3"
)

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

	//Handle graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	setupGracefulShutdown(cancel)
	// Load configuration

	data, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}

	var config config.Config

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("Ошибка парсинга YAML: %v", err)
	}

	fmt.Printf("config: %+v\n", config)

	// Initialize database using factory
	dbFactory := &databases.DatabaseFactory{}
	database, err := dbFactory.CreateDatabase(config.DBConfig.Type, config)
	if err != nil {
		//log.Fatalf("Error creating database: %v", err)//TODO temporary for testing
	}
	//defer database.Close()
	log.Println("Database connected successfully")

	// MongoDB Connection Handling
	var videoRepo entities.VideoRepository
	//var adRepo entities.AdvertisementRepository

	switch config.DBConfig.Type {
	case "mongo":
		cl, err := database.Connect()
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}
		mongoClient, ok := cl.(*mongo.Client)
		if !ok {
			log.Fatalf("Failed to cast to *mongo.Client")
		}
		videoRepo = repoMongo.NewMongoVideoRepository(mongoClient, config.DBConfig.Mongo.DBName, "videos")
		//adRepo = repoMongo.NewMongoAdvertisementRepository(mongoClient, config.DBConfig.Mongo.DBName, "advertisements")
	case "postgres":
		_, err := database.Connect()
		if err != nil {
			//TODO temporary for testing
			//log.Fatalf("Error connecting to PostgreSQL: %v", err)
		}
		//pgClient := psqlDB.(*sql.DB)
		//defer pgClient.Close()

		videoRepo = postgresql.NewPostgresVideoRepository(nil) //todo
		//adRepo = postgresql.NewPostgresAdvertisementRepository(nil)
	default:
		log.Fatalf("Unsupported database type for repositories: %+v", config.DBConfig)
	}

	log.Println("Initialize repositories successfully")

	// Initialize services
	videoService := services.NewVideoService(videoRepo)
	//adService := services.NewAdvertisementService(adRepo)

	//todo Temp, move func?
	////ads monetize
	//clientURL := "your_infura_or_other_client_url"
	//contractAddress := "your_contract_address"
	//
	//adInserter, err := ads.NewSmartContractAdInserter(clientURL, contractAddress)
	//if err != nil {
	//	log.Fatalf("Error creating ad inserter: %v", err)
	//}

	//adText := "Your ad text"
	//adImage := "Your ad image URL"
	//payment := big.NewInt(1000000000000000000) // 1 ether in wei

	//err = adInserter.PlaceAd(adText, adImage, payment)
	//if err != nil {
	//	log.Fatalf("Error placing ad: %v", err)
	//}

	// Create sample advertisement
	//sampleAd := &entities.Advertisement{
	//	ID:      "ad1",
	//	Title:   config.App.AdText,
	//	Content: config.App.AdImage,
	//	//URL:     config.App.,
	//}
	//if err = adService.CreateAd(sampleAd); err != nil {
	//	log.Fatalf("Error creating advertisement: %v", err)
	//}

	var wg sync.WaitGroup

	videoUploader := uploader.NewVideoUploader(config)
	//we get 1 account, and get most popular video 1 day, by category
	videoAnalysisService := analysis.NewVideoAnalysisService(videoUploader.YtClients[0]) //todo config - set, how much return videos
	videoDownloader := downloader.VideoDownloader{}
	videoProcessor := processor.VideoProcessor{}
	audioGenerator, err := audio.NewAudioGenerator(ctx)
	if err != nil {
		log.Fatalf("Error creating audio generator: %v", err)
	}
	seoGenerator := seo.SeoGenerator{}
	videoGenerator := ai_video.NewVideoGenerator(config.App.VeedAPIKey) // Инициализация нового модуля

	params := &videoProcessingParams{
		analysisService: videoAnalysisService,
		downloader:      videoDownloader,
		processor:       videoProcessor,
		audioGenerator:  audioGenerator,
		seoGenerator:    seoGenerator,
		videoService:    videoService,
		videoGenerator:  videoGenerator,
		config:          config,
	}

	//// Initialize deployment manager
	//services := []deployment.Service{
	//	{
	//		Name:       "PostgreSQL",
	//		StartCmd:   "docker-compose up -d db",
	//		StopCmd:    "docker-compose stop db",
	//		HealthCmd:  "docker-compose exec db pg_isready",
	//		RestartCmd: "docker-compose restart db",
	//	},
	//	{
	//		Name:       "MongoDB",
	//		StartCmd:   "docker-compose up -d mongo",
	//		StopCmd:    "docker-compose stop mongo",
	//		HealthCmd:  "docker-compose exec mongo mongo --eval 'db.runCommand({ ping: 1 })'",
	//		RestartCmd: "docker-compose restart mongo",
	//	},
	//	{
	//		Name:       "Prometheus",
	//		StartCmd:   "docker-compose up -d prometheus",
	//		StopCmd:    "docker-compose stop prometheus",
	//		HealthCmd:  "docker-compose exec prometheus curl -f http://localhost:9090/-/healthy",
	//		RestartCmd: "docker-compose restart prometheus",
	//	},
	//	{
	//		Name:       "MyApp",
	//		StartCmd:   "docker-compose up -d app",
	//		StopCmd:    "docker-compose stop app",
	//		HealthCmd:  "docker-compose exec app curl -f http://localhost:8080/health",
	//		RestartCmd: "docker-compose restart app",
	//	},
	//}
	//
	//deploymentManager := deployment.NewDeploymentManager(services)
	//
	//// Start services
	//deploymentManager.StartServices()

	monitoring.StartPerformanceMonitoring()

	// Run monitoring
	go monitoring.StartPrometheusMetrics()

	// Business logic to analyze, download, process, generate audio, SEO
	var processedVideos []*entities.Video

	// Process YouTube and TikTok videos
	processVideos(&wg, params, &processedVideos)

	wg.Wait()

	// Upload processed videos
	uploadVideos(videoUploader, processedVideos)

	//Shut down gracefully after successful video upload
	shutDown(ctx)
}

func processVideos(wg *sync.WaitGroup, params *videoProcessingParams, processedVideos *[]*entities.Video) {
	wg.Add(2) // Параллельная обработка для YouTube и TikTok
	go analyzeAndProcessVideos(wg, "youtube", params, processedVideos)
	go analyzeAndProcessVideos(wg, "tiktok", params, processedVideos)
}

func analyzeAndProcessVideos(wg *sync.WaitGroup, platform string, params *videoProcessingParams, processedVideos *[]*entities.Video) {
	defer wg.Done()

	popularVideos, err := params.analysisService.GetPopularVideos(platform, params.config.App.VideoCategory)
	if err != nil {
		log.Fatalf("Error getting popular videos: %v", err)
	}
	var descriptions []string

	for _, video := range popularVideos {

		descriptions = append(descriptions, video.Description)
		// Download video
		downloadedVideo, err := params.downloader.DownloadVideo(platform, video.URL, params.config.App.DownloadPath)
		if err != nil {
			log.Fatalf("Error downloading video: %v", err)
			continue
		}

		video.Content = downloadedVideo.Content
		video.FilePath = downloadedVideo.FilePath

		// Generate audio
		audioText := "This is a generated audio description."
		audioFile, err := params.audioGenerator.GenerateAudio(audioText, params.config.App.OutputPath+"/audio.mp3")
		if err != nil {
			log.Fatalf("Error generating audio: %v", err)
		}

		// Generate description and hashtags
		description := params.seoGenerator.GenerateDescription(video.Description)
		hashtags := params.seoGenerator.GenerateHashtags(video.Tags)

		// Save video info in DB
		//if err = params.videoService.SaveVideo(&video); err != nil {//TODO temp
		//	log.Fatalf("Error creating video: %v", err)
		//}

		// Process video
		newGeneratedVideo, err := params.processor.ProcessVideo(&video, params.config.App.OutputPath, audioText, audioFile)
		if err != nil {
			log.Fatalf("Error processing video: %v", err)
		}

		newGeneratedVideo.Description = description
		newGeneratedVideo.Tags = hashtags
		*processedVideos = append(*processedVideos, newGeneratedVideo)
	}

	// Create new videos based on descriptions
	createVideosFromDescriptions(params.videoGenerator, params.videoService, params.config, descriptions, processedVideos)
}

func createVideosFromDescriptions(
	videoGenerator *ai_video.VideoGenerator,
	videoService *services.VideoService,
	config config.Config,
	descriptions []string,
	processedVideos *[]*entities.Video,
) {
	for _, description := range descriptions {
		outputPath := fmt.Sprintf("%s/video_%d.mp4", config.App.OutputPath, time.Now().UnixNano())

		// Generate video from description
		content, err := videoGenerator.Generate(description, outputPath)
		if err != nil {
			log.Fatalf("Error creating video: %v", err)
			continue
		}

		// Create new video entity
		newVideo := &entities.Video{
			ID:          fmt.Sprintf("video_%d", time.Now().UnixNano()),
			Title:       "Generated Video", //add title
			Description: description,
			URL:         outputPath,
		}

		// Save video info in DB
		//if err = videoService.SaveVideo(newVideo); err != nil {//TODO temp
		//	log.Fatalf("Error saving video in DB: %v", err)
		//	continue
		//}

		newVideo.Content = content

		*processedVideos = append(*processedVideos, newVideo)
	}
}

func uploadVideos(uploader *uploader.VideoUploader, processedVideos []*entities.Video) {

	var wg sync.WaitGroup
	for _, newVideo := range processedVideos {
		wg.Add(1)
		go func(video *entities.Video) {
			defer wg.Done()
			description := video.Description
			hashtags := video.Tags

			// Upload video to YouTube
			if err := uploader.UploadToYouTube(video, description, hashtags); err != nil {
				log.Printf("Error uploading video to YouTube: %v", err)
			}

			//todo Upload video to TikTok
			//if err := videoUploader.UploadVideoToTikTok(video, description, hashtags); err != nil {
			//	log.Printf("Error uploading video to TikTok: %v", err)
			//}
		}(newVideo)
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
