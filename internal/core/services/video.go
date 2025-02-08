package services

import "canty/internal/core/entities"

type VideoService struct {
	repo entities.VideoRepository
}

func NewVideoService(repo entities.VideoRepository) *VideoService {
	return &VideoService{repo: repo}
}

func (s *VideoService) SaveVideo(video *entities.Video) error {
	return s.repo.Create(video)
}

func (s *VideoService) GetVideo(id string) (*entities.Video, error) {
	return s.repo.Read(id)
}

func (s *VideoService) UpdateVideo(video *entities.Video) error {
	return s.repo.Update(video)
}

func (s *VideoService) DeleteVideo(id string) error {
	return s.repo.Delete(id)
}
