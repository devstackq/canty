package analysis

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"canty/config"
	"canty/internal/core/entities"

	"google.golang.org/api/youtube/v3"
)

const (
	YouTube = "youtube"
	TikTok  = "tiktok"
)

// VideoAnalysisService осуществляет анализ видео для различных платформ.
type VideoAnalysisService struct {
	ytClients []*entities.YClient
	// TODO: tik tok
}

func NewVideoAnalysisService(ytClients map[string]*entities.YClient, cfg config.Config) *VideoAnalysisService {
	clients := make([]*entities.YClient, 0, len(cfg.YtAccounts))

	for _, account := range cfg.YtAccounts {
		if cl, ok := ytClients[account.Username]; ok {
			clients = append(clients, &entities.YClient{
				// Если в мапе хранится готовый клиент, можно использовать его напрямую:
				Client:   cl.Client,
				Category: account.Category,
				UserName: account.Username, // Используем имя пользователя, а не категорию
			})
		}
	}

	return &VideoAnalysisService{
		ytClients: clients,
	}
}

// GetPopularVideos возвращает популярные видео для указанной платформы.
// Результат – мапа, где ключ – username, а значение – срез видео.
func (vas *VideoAnalysisService) GetPopularVideos(platform string) (map[string][]entities.Video, error) {
	switch platform {
	case YouTube:
		return vas.getPopularYouTubeVideos()
	case TikTok:
		// TODO: Реализовать получение популярных видео для TikTok
		return nil, fmt.Errorf("TikTok functionality not implemented")
	default:
		return nil, fmt.Errorf("unsupported platform: %s", platform)
	}
}

/*
Film & Animation – ID: 1
Autos & Vehicles – ID: 2
Music – ID: 10
Pets & Animals – ID: 15
Sports – ID: 17
Short Movies – ID: 18
Travel & Events – ID: 19
Gaming – ID: 20
People & Blogs – ID: 22
Comedy – ID: 23
Entertainment – ID: 24
News & Politics – ID: 25
Howto & Style – ID: 26
Education – ID: 27
Science & Technology – ID: 28
Nonprofits & Activism – ID: 29
Movies – ID: 30
Anime/Animation – ID: 31
*/

// getPopularYouTubeVideos запрашивает у каждого YouTube-клиента список популярных видео.
// Для каждого клиента вызывается метод API, и результаты собираются в мапу: username -> []entities.Video.
func (vas *VideoAnalysisService) getPopularYouTubeVideos() (map[string][]entities.Video, error) {
	result := make(map[string][]entities.Video)

	// Обходим всех клиентов
	for _, ytClient := range vas.ytClients {
		if ytClient.Client == nil {
			log.Printf("ytClient.Client is nil for user %s", ytClient.UserName)
			continue
		}

		// 1. Выполняем запрос поиска видео за последние 7 дней по заданной категории
		searchCall := ytClient.Client.Search.List([]string{"snippet"}).
			RegionCode("US"). // get from config
			PublishedAfter(time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339)).
			VideoCategoryId(ytClient.Category).
			Type("video").
			Q("a").
			MaxResults(1) // get from config

		searchResponse, err := searchCall.Do()
		if err != nil {
			return nil, err
		}

		if len(searchResponse.Items) == 0 {
			return nil, err
		}

		var videos = make([]entities.Video, 0, len(searchResponse.Items)) //todo - now we get 1 video

		// Из результатов получаем идентификатор видео
		searchItem := searchResponse.Items[0]
		videoID := searchItem.Id.VideoId

		// 2. Получаем подробную информацию о видео (включая теги) через Videos.list
		videosCall := ytClient.Client.Videos.List([]string{"snippet"}).Id(videoID)
		videosResponse, err := videosCall.Do()
		if err != nil {
			return nil, err
		}
		if len(videosResponse.Items) == 0 {
			return nil, errors.New("no videos found")
		}

		videoDetails := videosResponse.Items[0]

		// 3. Пытаемся получить субтитры для видео через Captions.list
		var subtitles string
		captionsCall := ytClient.Client.Captions.List([]string{"snippet"}, videoID)
		captionsResponse, err := captionsCall.Do()
		if err != nil {
			return nil, err
		}

		if len(captionsResponse.Items) > 0 {
			// Выбираем дорожку, предпочтительно на английском языке
			var chosenCaption *youtube.Caption
			for _, caption := range captionsResponse.Items {
				if caption.Snippet.Language == "en" { // язык можно вынести в конфигурацию
					chosenCaption = caption
					break
				}
			}
			if chosenCaption == nil {
				// Если нет английской дорожки, берем первую доступную
				chosenCaption = captionsResponse.Items[0]
			}

			// 4. Загружаем субтитры с выбранной дорожки в формате SRT
			downloadCall := ytClient.Client.Captions.Download(chosenCaption.Id).Tfmt("srt")
			resp, err := downloadCall.Download()
			if err != nil {
				return nil, err
				continue
			}
			defer resp.Body.Close()

			data, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}

			subtitles = string(data)
		}

		// Формируем URL видео
		videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

		// Собираем объект Video с тегами и субтитрами
		video := entities.Video{
			Title:       videoDetails.Snippet.Title,       // todo gen new Title? by Gpt
			Description: videoDetails.Snippet.Description, // todo gen new Description? by GTP
			URL:         videoURL,
			Tags:        videoDetails.Snippet.Tags,
			Subtitles:   subtitles,
		}

		videos = append(videos, video)

		result[ytClient.UserName] = videos
	}

	return result, nil
}
