package entities

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/api/youtube/v3"
)

type VideoRepository interface {
	Create(video *Video) error
	Read(id string) (*Video, error)
	Update(video *Video) error
	Delete(id string) error
}

type Video struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Content     []byte   `json:"content"`
	FilePath    string   `json:"file_path"`
	Format      string   `json:"format"`
	Duration    int      `json:"duration"`
	Resolution  string   `json:"resolution"`
	Size        int64    `json:"size"`
	Tags        []string `json:"tags"`
	Subtitles   string   `json:"subtitles"`
}

type YClient struct {
	Client   *youtube.Service
	Category string
	UserName string
}

// UploadVideo загружает видео на YouTube, используя данные из video.
// Он открывает файл видео, формирует запрос к API и возвращает ошибку в случае неудачи.
func (y *YClient) UploadVideo(video *Video) error {
	// Открываем файл видео по указанному пути
	file, err := os.Open(video.FilePath)
	if err != nil {
		return fmt.Errorf("failed to open video file: %w", err)
	}
	defer file.Close()

	// Формируем объект запроса с информацией о видео
	call := y.Client.Videos.Insert([]string{"snippet", "status"}, &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       video.Title,
			Description: video.Description,
			Tags:        video.Tags,
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "public",
		},
	})

	// Выполняем запрос, передавая поток файла
	response, err := call.Media(file).Do()
	if err != nil {
		return fmt.Errorf("failed to upload video: %w", err)
	}

	log.Printf("Video uploaded successfully: %s", response.Id)
	return nil
}
