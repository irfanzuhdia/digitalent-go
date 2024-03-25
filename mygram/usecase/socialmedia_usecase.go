package usecase

import (
	"mygram/domain"
	"mygram/interfaces"
)

type SocialMediaUseCase struct {
	socialMediaRepo interfaces.SocialMediaRepository
}

func NewSocialMediaUseCase(socialMediaRepo interfaces.SocialMediaRepository) *SocialMediaUseCase {
	return &SocialMediaUseCase{socialMediaRepo}
}

func (s *SocialMediaUseCase) Create(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	return s.socialMediaRepo.Save(socialMedia)
}

func (s *SocialMediaUseCase) GetAll(userID uint) ([]domain.SocialMedia, error) {
	return s.socialMediaRepo.FindAll(userID)
}

func (s *SocialMediaUseCase) GetByID(id uint) (domain.SocialMedia, error) {
	return s.socialMediaRepo.FindByID(id)
}

func (s *SocialMediaUseCase) Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	return s.socialMediaRepo.Update(socialMedia)
}

func (s *SocialMediaUseCase) Delete(id uint) error {
	return s.socialMediaRepo.Delete(id)
}
