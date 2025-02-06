package processor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"canty/internal/core/entities"
)

//	type VideoProcessor interface {
//		AddTextToVideo(inputPath, outputPath, text string) error
//		AddAudioToVideo(videoPath, audioPath, outputPath string) error
//	}
type VideoProcessor struct{}

func (vp *VideoProcessor) ProcessVideo(video *entities.Video, outputPath, text string, audioContent []byte) (*entities.Video, error) {
	// Если видео содержится в []byte
	if len(video.Content) > 0 {
		// Временное сохранение видео в файл для обработки
		tempFile, err := ioutil.TempFile("", "temp-video-*.mp4")
		if err != nil {
			return nil, fmt.Errorf("failed to create temp file: %w", err)
		}
		defer os.Remove(tempFile.Name())

		_, err = tempFile.Write(video.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to write to temp file: %w", err)
		}
		err = tempFile.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to close temp file: %w", err)
		}

		video.FilePath = tempFile.Name()
	}

	// Временное сохранение аудио в файл для обработки
	audioTempFile, err := ioutil.TempFile("", "temp-audio-*.mp3")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp audio file: %w", err)
	}
	defer os.Remove(audioTempFile.Name())

	_, err = audioTempFile.Write(audioContent)
	if err != nil {
		return nil, fmt.Errorf("failed to write to temp audio file: %w", err)
	}
	err = audioTempFile.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close temp audio file: %w", err)
	}

	// Добавление текста к видео
	textTempFile, err := ioutil.TempFile("", "temp-text-video-*.mp4")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp text video file: %w", err)
	}
	defer os.Remove(textTempFile.Name())

	cmd := exec.Command("ffmpeg", "-i", video.FilePath, "-vf", fmt.Sprintf("drawtext=text='%s'", text), textTempFile.Name())
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to add text to video: %w", err)
	}

	// Добавление аудио к видео
	cmd = exec.Command("ffmpeg", "-i", textTempFile.Name(), "-i", audioTempFile.Name(), "-c:v", "copy", "-c:a", "aac", outputPath)
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to add audio to video: %w", err)
	}

	// Обновление пути к обработанному видео
	video.FilePath = outputPath
	return video, nil
}
