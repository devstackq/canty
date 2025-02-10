package audio

import (
	"context"
	"os"

	"cloud.google.com/go/texttospeech/apiv1"
	"google.golang.org/api/option"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

// AudioGenerator инкапсулирует клиента Text-to-Speech.
type AudioGenerator struct {
	ttsClient *texttospeech.Client
}

// NewAudioGenerator создаёт экземпляр AudioGenerator.
func NewAudioGenerator(ctx context.Context) (*AudioGenerator, error) {
	client, err := texttospeech.NewClient(ctx, option.WithCredentialsFile("config/g_config.json"))
	if err != nil {
		return nil, err
	}
	return &AudioGenerator{ttsClient: client}, nil
}

// GenerateAudio синтезирует речь по заданному тексту, сохраняет аудиофайл по outputPath
// и возвращает аудиоданные в виде []byte.
// Контекст передается в метод, чтобы можно было управлять временем выполнения и отменой операции.
func (ag *AudioGenerator) GenerateAudio(ctx context.Context, text, outputPath string) ([]byte, error) {
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			// Если требуется, можно вынести язык в конфигурацию
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US", // язык можно вынести в конфигурацию
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := ag.ttsClient.SynthesizeSpeech(ctx, req)
	if err != nil {
		return nil, err
	}

	// Записываем полученный аудиоконтент в файл с помощью os.WriteFile
	if err = os.WriteFile(outputPath, resp.AudioContent, 0644); err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}
