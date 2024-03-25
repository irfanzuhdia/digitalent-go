package interfaces

import "mygram/domain"

type SocialMediaRepository interface {
	Save(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	FindAll(userID uint) ([]domain.SocialMedia, error)
	FindByID(id uint) (domain.SocialMedia, error)
	Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error)
	Delete(id uint) error
}
