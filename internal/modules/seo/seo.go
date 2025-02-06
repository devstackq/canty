package seo

import (
	"strings"
)

type SeoGenerator struct{}

func NewSeoGenerator() *SeoGenerator {
	return &SeoGenerator{}
}

// GenerateDescription генерирует описание для видео на основе заголовка.
func (sg *SeoGenerator) GenerateDescription(title string) string {
	description := "Watch " + title + " and enjoy the content. Don't forget to like, share, and subscribe!"
	return description
}

// GenerateHashtags генерирует хэштеги для видео на основе заголовка.
func (sg *SeoGenerator) GenerateHashtags(tags []string) []string {
	var hashtags []string
	for _, word := range tags {
		hashtags = append(hashtags, "#"+strings.ToLower(word))
	}
	return hashtags
}
