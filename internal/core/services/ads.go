package services

import (
	"canty/internal/core/entities"
)

type AdvertisementService struct {
	repo entities.AdvertisementRepository
}

func NewAdvertisementService(repo entities.AdvertisementRepository) *AdvertisementService {
	return &AdvertisementService{repo: repo}
}

func (s *AdvertisementService) CreateAd(ad *entities.Advertisement) error {
	return s.repo.Create(ad)
}

func (s *AdvertisementService) GetAd(id string) (*entities.Advertisement, error) {
	return s.repo.Read(id)
}

func (s *AdvertisementService) UpdateAd(ad *entities.Advertisement) error {
	return s.repo.Update(ad)
}

func (s *AdvertisementService) DeleteAd(id string) error {
	return s.repo.Delete(id)
}
