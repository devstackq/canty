package processor

import (
	"fmt"
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

	// Если видео содержится в []byte, сохраняем его во временный файл.
	if len(video.Content) > 0 {
		tmpVideoFile, err := os.CreateTemp("", "temp-video-*.mp4")
		if err != nil {
			return video, fmt.Errorf("failed to create temp video file: %w", err)
		}
		// Удаляем временный файл после завершения работы.
		defer os.Remove(tmpVideoFile.Name())

		if err := os.WriteFile(tmpVideoFile.Name(), video.Content, 0644); err != nil {
			return video, fmt.Errorf("failed to write video content: %w", err)
		}
		video.FilePath = tmpVideoFile.Name()
	}

	// Сохраняем аудио во временный файл.
	tmpAudioFile, err := os.CreateTemp("", "temp-audio-*.mp3")
	if err != nil {
		return video, fmt.Errorf("failed to create temp audio file: %w", err)
	}
	defer os.Remove(tmpAudioFile.Name())

	if err := os.WriteFile(tmpAudioFile.Name(), audioContent, 0644); err != nil {
		return video, fmt.Errorf("failed to write audio content: %w", err)
	}

	// Создаем временный файл для видео с добавленным текстом.
	tmpTextVideoFile, err := os.CreateTemp("", "temp-text-video-*.mp4")
	if err != nil {
		return video, fmt.Errorf("failed to create temp text video file: %w", err)
	}
	defer os.Remove(tmpTextVideoFile.Name())

	// Добавляем текст к видео с помощью ffmpeg.
	ffmpegArgsText := []string{
		"-i", video.FilePath,
		"-vf", fmt.Sprintf("drawtext=text='%s':fontcolor=white:fontsize=24", text),
		"-y", // автоматическая перезапись выходного файла
		tmpTextVideoFile.Name(),
	}
	cmd := exec.Command("ffmpeg", ffmpegArgsText...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return video, fmt.Errorf("failed to add text to video: %w, output: %s", err, output)
	}

	// Добавляем аудио к видео.
	ffmpegArgsAudio := []string{
		"-i", tmpTextVideoFile.Name(),
		"-i", tmpAudioFile.Name(),
		"-c:v", "copy",
		"-c:a", "aac",
		"-y", // автоматическая перезапись выходного файла
		outputPath,
	}
	cmd = exec.Command("ffmpeg", ffmpegArgsAudio...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return video, fmt.Errorf("failed to add audio to video: %w, output: %s", err, output)
	}

	// Обновляем путь к итоговому видео.
	video.FilePath = outputPath

	// Читаем финальный обработанный видеофайл и обновляем video.Content.
	processedContent, err := os.ReadFile(outputPath)
	if err != nil {
		return video, fmt.Errorf("failed to read processed video file: %w", err)
	}
	video.Content = processedContent

	return video, nil

}
