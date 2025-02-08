package analysis

import (
	"fmt"

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

// getPopularYouTubeVideos запрашивает у каждого YouTube-клиента список популярных видео.
// Для каждого клиента вызывается метод API, и результаты собираются в мапу: username -> []entities.Video.
func (vas *VideoAnalysisService) getPopularYouTubeVideos() (map[string][]entities.Video, error) {
	result := make(map[string][]entities.Video)

	// Обходим всех клиентов
	for _, ytClient := range vas.ytClients {
		// Формируем запрос: получаем популярные видео по заданной категории
		call := ytClient.Client.Videos.List([]string{"snippet"}).
			Chart("mostPopular").
			VideoCategoryId(ytClient.Category).
			MaxResults(1) // Можно вынести в конфигурацию
		response, err := call.Do()
		if err != nil {
			return nil, fmt.Errorf("error making API call to YouTube for user %s: %w", ytClient.UserName, err)
		}

		videos := make([]entities.Video, 0, len(response.Items))
		for _, item := range response.Items {
			videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id)
			videos = append(videos, entities.Video{
				Title:       item.Snippet.Title,
				Description: item.Snippet.Description,
				URL:         videoURL,
				Tags:        item.Snippet.Tags,
			})
		}
		result[ytClient.UserName] = videos
	}

	return result, nil
}
