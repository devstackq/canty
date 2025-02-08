package downloader

import (
	"fmt"
	"os"
	"os/exec"

	"canty/internal/core/entities"
)

type Platform string

const (
	YouTube = "youtube"
	TikTok  = "tiktok"
)

type VideoDownloader struct {
	maxConcurrentDownloads int
}

func NewVideoDownloader(maxConcurrentDownloads int) *VideoDownloader {
	return &VideoDownloader{
		maxConcurrentDownloads: maxConcurrentDownloads,
	}
}

func (vd *VideoDownloader) DownloadVideo(platform, url string, outputPath string) (entities.Video, error) {
	// Реализуйте контроль максимального количества параллельных скачиваний
	switch platform {
	case YouTube:
		return vd.downloadYouTubeVideo(url, outputPath)
	case TikTok:
		return vd.downloadTikTokVideo(url, outputPath)
	default:
		return entities.Video{}, fmt.Errorf("unsupported platform")
	}
}

func (vd *VideoDownloader) downloadYouTubeVideo(videoURL string, outputPath string) (entities.Video, error) {
	// Создаем временный файл для скачивания видео.
	tmpFile, err := os.CreateTemp("", "video-*.mp4")
	if err != nil {
		return entities.Video{}, fmt.Errorf("не удалось создать временный файл: %w", err)
	}
	fileName := tmpFile.Name()
	// Закрываем файл, так как youtube-dl перезапишет его.
	tmpFile.Close()

	// Выполняем команду youtube-dl для скачивания видео.
	cmd := exec.Command("youtube-dl", "-o", fileName, videoURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return entities.Video{}, fmt.Errorf("не удалось скачать видео: %w, вывод: %s", err, output)
	}

	// Читаем содержимое файла в []byte.
	videoBytes, err := os.ReadFile(fileName)
	if err != nil {
		return entities.Video{}, fmt.Errorf("не удалось прочитать видеофайл: %w", err)
	}

	video := entities.Video{
		URL:      videoURL,
		Content:  videoBytes,
		FilePath: fileName,
	}

	fmt.Printf("downloadYouTubeVideo video: %+v", video)

	return video, nil
}

func (vd *VideoDownloader) downloadTikTokVideo(url string, outputPath string) (entities.Video, error) {
	// Реализуйте логику для скачивания видео из TikTok
	return entities.Video{}, fmt.Errorf("not implemented")
}
