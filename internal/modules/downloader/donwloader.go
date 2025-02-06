package downloader

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"

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

func (vd *VideoDownloader) downloadYouTubeVideo(url string, outputPath string) (entities.Video, error) {
	fileName := filepath.Join(url, "video.mp4")
	cmd := exec.Command("youtube-dl", "-o", fileName, url)
	if err := cmd.Run(); err != nil {
		return entities.Video{}, fmt.Errorf("failed to download video: %w", err)
	}

	// Чтение содержимого файла как []byte
	videoBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return entities.Video{}, fmt.Errorf("failed to read video file: %w", err)
	}

	video := entities.Video{
		URL:      url,
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
