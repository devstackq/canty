package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"canty/config"
	"canty/internal/core/entities"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

// VideoUploader отвечает за инициализацию и организацию загрузки видео
type VideoUploader struct {
	YClients map[string]*entities.YClient
	// TClients []*api.TikTokClient // если потребуется
}

// NewVideoUploader создает VideoUploader и инициализирует YouTube клиентов
func NewVideoUploader(ctx context.Context, cfg config.Config) (*VideoUploader, error) {
	ytClients, err := initializeYouTubeClients(ctx, cfg.YtAccounts)
	if err != nil {
		return nil, err
	}

	return &VideoUploader{
		YClients: ytClients,
		// TClients: initializeTikTokClients(cfg.TikTok.Accounts),
	}, nil
}

// initializeYouTubeClients создает клиентов для каждого аккаунта на YouTube
func initializeYouTubeClients(ctx context.Context, accounts []config.YouTubeAccount) (map[string]*entities.YClient, error) {
	yClients := make(map[string]*entities.YClient, len(accounts))
	for _, ytAccount := range accounts {
		if ytAccount.Credentials != nil {
			// Формируем JSON-данные из полей credentials
			credsData := map[string]interface{}{
				"installed": map[string]interface{}{
					"client_id":                   ytAccount.Credentials.Installed.ClientID,
					"project_id":                  ytAccount.Credentials.Installed.ProjectID,
					"auth_uri":                    ytAccount.Credentials.Installed.AuthURI,
					"token_uri":                   ytAccount.Credentials.Installed.TokenURI,
					"auth_provider_x509_cert_url": ytAccount.Credentials.Installed.AuthProviderCertURL,
					"client_secret":               ytAccount.Credentials.Installed.ClientSecret,
					"redirect_uris":               ytAccount.Credentials.Installed.RedirectURIs,
				},
			}
			jsonBytes, err := json.Marshal(credsData)
			if err != nil {
				return nil, err
			}
			creds, err := google.CredentialsFromJSON(ctx, jsonBytes, youtube.YoutubeForceSslScope)
			if err != nil {
				return nil, err
			}
			client, err := youtube.NewService(ctx, option.WithCredentials(creds))
			if err != nil {
				return nil, err
			}

			yClients[ytAccount.Username] = &entities.YClient{
				Client:   client,
				Category: ytAccount.Category,
				UserName: ytAccount.Username,
			}
		}
	}

	return yClients, nil
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
