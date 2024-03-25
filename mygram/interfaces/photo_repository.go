package interfaces

import "mygram/domain"

type PhotoRepository interface {
	Save(photo domain.Photo) (domain.Photo, error)
	FindAll() ([]domain.Photo, error)
	FindByID(id uint) (domain.Photo, error)
	Update(photo domain.Photo) (domain.Photo, error)
	Delete(id uint) error
}
