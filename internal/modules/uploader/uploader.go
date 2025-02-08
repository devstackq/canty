package uploader

import (
	"context"
	"fmt"
	"log"

	"canty/config"
	"canty/internal/core/entities"

	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

// VideoUploader отвечает за инициализацию и организацию загрузки видео
type VideoUploader struct {
	YClients map[string]*entities.YClient
	// TClients []*api.TikTokClient // если потребуется
}

// NewVideoUploader создает VideoUploader и инициализирует YouTube клиентов
func NewVideoUploader(cfg config.Config) *VideoUploader {
	ctx := context.Background()
	return &VideoUploader{
		YClients: initializeYouTubeClients(ctx, cfg.Youtube.Accounts),
		// TClients: initializeTikTokClients(cfg.TikTok.Accounts),
	}
}

// initializeYouTubeClients создает клиентов для каждого аккаунта на YouTube
func initializeYouTubeClients(ctx context.Context, accounts []config.AccountConfig) map[string]*entities.YClient {
	yClients := make(map[string]*entities.YClient, len(accounts))
	for _, account := range accounts {
		client, err := youtube.NewService(ctx, option.WithAPIKey(account.ApiKey))
		if err != nil {
			log.Printf("Error creating YouTube service for %s: %v", account.Username, err)
			continue
		}
		yClients[account.Username] = &entities.YClient{
			Client:   client,
			Category: account.Category,
			UserName: account.Username,
		}
	}
	return yClients
}

// UploadToYouTube – метод высокого уровня, который для каждого нужного пользователя вызывает загрузку видео
// Здесь происходит поиск соответствующего клиента и вызов метода UploadVideo, инкапсулированного в entities.YClient.
func (vu *VideoUploader) UploadToYouTube(newVideo *entities.Video) error {
	// Например, если вы хотите загрузить видео для каждого клиента, либо выбрать определенного:
	for username, yClient := range vu.YClients {
		// Можно добавить логику, выбирающую конкретного клиента по какому-либо критерию
		log.Printf("Uploading video for user %s", username)
		if err := yClient.UploadVideo(newVideo); err != nil {
			return fmt.Errorf("failed to upload video for %s: %w", username, err)
		}
	}
	return nil
}

// Дополнительные функции (например, для TikTok) можно добавить аналогично
