package uploader

import (
	"fmt"
	"log"
	"os"

	"canty/config"
	"canty/internal/core/entities"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	youtube "google.golang.org/api/youtube/v3"
)

type VideoUploader struct {
	YtClients []*youtube.Service
	//TkClients []*api.TikTokClient
}

func NewVideoUploader(config config.Config) *VideoUploader {
	return &VideoUploader{
		YtClients: initializeYouTubeClients(config.Youtube.Accounts),
		//TkClients: initializeTikTokClients(config.TikTok.Accounts),
	}
}

func initializeYouTubeClients(accounts []config.AccountConfig) []*youtube.Service {
	ytClients := make([]*youtube.Service, 0)
	for _, account := range accounts {
		client, err := youtube.NewService(context.Background(), option.WithAPIKey(account.ApiKey))
		if err != nil {
			fmt.Printf("Error creating YouTube service: %v", err)
			return nil
		}
		ytClients = append(ytClients, client)
	}
	return ytClients
}

//func initializeTikTokClients(accounts []config.TikTokAccount) []*api.TikTokClient {
//	tkClients := make([]*api.TikTokClient, 0)
//for _, account := range accounts {
//	client, err := api.NewTikTokClient(account.AccessToken)
//	if err != nil {
//		log.Fatalf("Error creating TikTok client: %v", err)
//	}
//	tkClients = append(tkClients, client)
//}
//return tkClients
//}

func (vu *VideoUploader) UploadToYouTube(newVideo *entities.Video, description string, hashtags []string) error {
	for _, ytClient := range vu.YtClients {
		err := vu.uploadToYouTube(ytClient, newVideo, description, hashtags)
		if err != nil {
			return err
		}
	}
	return nil
}

//func (vu *VideoUploader) UploadVideoToTikTok(newVideo *entities.Video, description string, hashtags []string) error {
//	for _, tkClient := range vu.TkClients {
//		err := vu.uploadToTikTok(tkClient, newVideo, description, hashtags)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (vu *VideoUploader) uploadToYouTube(ytClient *youtube.Service, video *entities.Video, description string, tags []string) error {
	file, err := os.Open(video.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open video file: %w", err)
	}
	defer file.Close()

	call := ytClient.Videos.Insert([]string{"snippet", "status"}, &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       video.Title,
			Description: description,
			Tags:        tags,
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "public"},
	})

	response, err := call.Media(file).Do()
	if err != nil {
		return fmt.Errorf("failed to upload video: %w", err)
	}
	log.Printf("Video uploaded successfully: %s", response.Id)
	return nil
}

func (vu *VideoUploader) uploadTikTokVideo(filePath, title, description string, tags []string) error {
	// Реализуйте логику для загрузки видео на TikTok
	return nil
}
