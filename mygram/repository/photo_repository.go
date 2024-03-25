package repository

import (
	"mygram/domain"
	"mygram/interfaces"

	"gorm.io/gorm"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) interfaces.PhotoRepository {
	return &PhotoRepository{db}
}

func (r *PhotoRepository) Save(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Create(&photo).Error
	return photo, err
}

func (r *PhotoRepository) FindAll() ([]domain.Photo, error) {
	var photos []domain.Photo
	err := r.db.Preload("User").Find(&photos).Error
	return photos, err
}

func (r *PhotoRepository) FindByID(id uint) (domain.Photo, error) {
	var photo domain.Photo
	err := r.db.Preload("User").Where("id = ?", id).First(&photo).Error
	return photo, err
}

func (r *PhotoRepository) Update(photo domain.Photo) (domain.Photo, error) {
	err := r.db.Save(&photo).Error
	return photo, err
}

func (r *PhotoRepository) Delete(id uint) error {
	var photo domain.Photo
	err := r.db.Where("id = ?", id).Delete(&photo).Error
	return err
}
