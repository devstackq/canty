package youtube

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func NewYouTubeClient(apiKey string) *youtube.Service {
	client, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}
	return client
}
