package usecase

import (
	"mygram/domain"
	"mygram/interfaces"
)

type PhotoUseCase struct {
	photoRepo interfaces.PhotoRepository
}

func NewPhotoUseCase(photoRepo interfaces.PhotoRepository) *PhotoUseCase {
	return &PhotoUseCase{photoRepo}
}

func (p *PhotoUseCase) Create(photo domain.Photo) (domain.Photo, error) {
	return p.photoRepo.Save(photo)
}

func (p *PhotoUseCase) GetAll() ([]domain.Photo, error) {
	return p.photoRepo.FindAll()
}

func (p *PhotoUseCase) GetByID(id uint) (domain.Photo, error) {
	return p.photoRepo.FindByID(id)
}

func (p *PhotoUseCase) Update(photo domain.Photo) (domain.Photo, error) {
	return p.photoRepo.Update(photo)
}

func (p *PhotoUseCase) Delete(id uint) error {
	return p.photoRepo.Delete(id)
}
