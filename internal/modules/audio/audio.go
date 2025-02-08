package audio

import (
	"context"
	"io/ioutil"

	"cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

type AudioGenerator struct {
	ttsClient *texttospeech.Client
}

func NewAudioGenerator(ctx context.Context) (*AudioGenerator, error) {
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &AudioGenerator{ttsClient: client}, nil
}

func (ag *AudioGenerator) GenerateAudio(text, outputPath string) ([]byte, error) {
	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_NEUTRAL,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := ag.ttsClient.SynthesizeSpeech(context.Background(), req)
	if err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(outputPath, resp.AudioContent, 0644); err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}
