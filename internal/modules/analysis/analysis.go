package analysis

import (
	"fmt"
	"time"

	"canty/config"
	"canty/internal/core/entities"
)

const (
	YouTube = "youtube"
	TikTok  = "tiktok"
)

// VideoAnalysisService осуществляет анализ видео для различных платформ.
type VideoAnalysisService struct {
	ytClients []entities.YClient
	// TODO: Добавить поддержку TikTok, например: tkClients *api.TikTokClient
}

// NewVideoAnalysisService создает новый экземпляр VideoAnalysisService.
// При этом формируется срез YouTube-клиентов (YClients) на основе AccountConfig,
// используя переданный мапу клиентов, где ключ – username.
func NewVideoAnalysisService(ytClients map[string]*entities.YClient, cfg config.Config) *VideoAnalysisService {
	clients := make([]entities.YClient, 0, len(cfg.Youtube.Accounts))
	for _, account := range cfg.Youtube.Accounts {
		if cl, ok := ytClients[account.Username]; ok {
			clients = append(clients, entities.YClient{
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
		// Формируем запрос: получаем популярные видео по заданной категории
		searchCall := ytClient.Client.Search.List([]string{"snippet"}).
			RegionCode("US"). //todo
			PublishedAfter(time.Now().Add(-7 * 24 * time.Hour).Format(time.RFC3339)).
			VideoCategoryId(ytClient.Category).
			Type("video").
			MaxResults(1) //todo

		searchResponse, err := searchCall.Do()
		if err != nil {
			return nil, err
		}

		// Собираем все ID видео
		var videoIDs []string
		for _, item := range searchResponse.Items {
			videoIDs = append(videoIDs, item.Id.VideoId)
		}

		// Запрашиваем полную информацию о видео, включая теги, через Videos.list
		videosCall := ytClient.Client.Videos.List([]string{"snippet"}).
			Id(videoIDs...)
		videosResponse, err := videosCall.Do()
		if err != nil {
			return nil, err
		}
		videos := make([]entities.Video, 0, len(videosResponse.Items))

		// Обрабатываем полученные видео, включая теги
		for _, item := range videosResponse.Items {

			videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id)
			videos = append(videos, entities.Video{
				Title:       item.Snippet.Title,       //todo generate new title by chatGPT?
				Description: item.Snippet.Description, //todo generate new Description by chatGPT?
				URL:         videoURL,
				Tags:        item.Snippet.Tags,
			})
		}
		result[ytClient.UserName] = videos
	}

	return result, nil
}
