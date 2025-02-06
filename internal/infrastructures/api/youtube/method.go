package youtube

import (
	"google.golang.org/api/youtube/v3"
)

func GetPopularVideos(client *youtube.Service, category string) ([]*youtube.Video, error) {
	call := client.Videos.List([]string{"snippet"}).Chart("mostPopular").VideoCategoryId(category)
	response, err := call.Do()
	if err != nil {
		return nil, err
	}
	return response.Items, nil
}
