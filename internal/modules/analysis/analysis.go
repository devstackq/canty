package analysis

import (
	"fmt"
	"log"

	"canty/internal/core/entities"

	"google.golang.org/api/youtube/v3"
)

const (
	YouTube = "youtube"
	TikTok  = "tiktok"
)

type VideoAnalysisService struct {
	//ytClient map[string]*youtube.Service
	ytClient *youtube.Service
	//tkClient *api.TikTokClient todo
}

func NewVideoAnalysisService(ytClient *youtube.Service) *VideoAnalysisService { //2 param tkClient *api.TikTokClient)
	return &VideoAnalysisService{
		ytClient: ytClient,
		//tkClient: tkClient,
	}
}

func (vas *VideoAnalysisService) GetPopularVideos(platform string, category string) ([]entities.Video, error) {
	switch platform {
	case YouTube:
		return vas.getPopularYouTubeVideos(category)
	case TikTok:
		//return vas.getPopularTikTokVideos(account, category)
	default:
		return nil, fmt.Errorf("unsupported platform")
	}
	return nil, fmt.Errorf("unsupported platform")
}

func (vas *VideoAnalysisService) getPopularYouTubeVideos(category string) ([]entities.Video, error) {

	call := vas.ytClient.Videos.List([]string{"snippet"}).
		Chart("mostPopular").
		VideoCategoryId(category).
		MaxResults(5)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making API call to YouTube: %v", err)
	}

	//todo ref
	var video = make([]entities.Video, 0, len(response.Items))

	for _, item := range response.Items {
		videoURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.Id)

		video = append(video, entities.Video{
			Title:       item.Snippet.Title,
			Description: item.Snippet.Description,
			URL:         videoURL,
			Tags:        item.Snippet.Tags,
		})
	}

	return video, nil
}

//func (vas *VideoAnalysisService) getPopularTikTokVideos(account string, category string) ([]*TikTokVideo, error) {
//	return fmt.Errorf("not implemented")
//}
