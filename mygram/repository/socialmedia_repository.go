package repository

import (
	"mygram/domain"
	"mygram/interfaces"

	"gorm.io/gorm"
)

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) interfaces.SocialMediaRepository {
	return &SocialMediaRepository{db}
}

func (r *SocialMediaRepository) Save(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Create(&socialMedia).Error
	return socialMedia, err
}

func (r *SocialMediaRepository) FindAll(userID uint) ([]domain.SocialMedia, error) {
	var socialMedias []domain.SocialMedia
	err := r.db.Preload("User").Where("user_id = ?", userID).Find(&socialMedias).Error
	return socialMedias, err
}

func (r *SocialMediaRepository) FindByID(id uint) (domain.SocialMedia, error) {
	var socialMedia domain.SocialMedia
	err := r.db.Preload("User").Where("id = ?", id).First(&socialMedia).Error
	return socialMedia, err
}

func (r *SocialMediaRepository) Update(socialMedia domain.SocialMedia) (domain.SocialMedia, error) {
	err := r.db.Save(&socialMedia).Error
	return socialMedia, err
}

func (r *SocialMediaRepository) Delete(id uint) error {
	var socialMedia domain.SocialMedia
	err := r.db.Where("id = ?", id).Delete(&socialMedia).Error
	return err
}
