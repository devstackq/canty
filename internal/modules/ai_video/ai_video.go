package ai_video

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VideoGenerator struct {
	apiKey string
}

func NewVideoGenerator(apiKey string) *VideoGenerator {
	return &VideoGenerator{apiKey: apiKey}
}

func (vg *VideoGenerator) Generate(description, outputPath string) ([]byte, error) {
	// Параметры для запроса к API VEED.IO
	requestBody, err := json.Marshal(map[string]string{
		"description": description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create request body: %w", err)
	}

	// Создание запроса к API VEED.IO
	req, err := http.NewRequest("POST", "https://api.veed.io/generate-video", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+vg.apiKey)

	// Отправка запроса и получение ответа
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Обработка ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to generate video, status code: %d", resp.StatusCode)
	}

	// Сохранение видео в файл
	videoContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = ioutil.WriteFile(outputPath, videoContent, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to save video: %w", err)
	}

	return videoContent, nil
}
