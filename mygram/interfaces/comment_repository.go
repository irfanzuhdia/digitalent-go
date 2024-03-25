package interfaces

import "mygram/domain"

type CommentRepository interface {
	Save(comment domain.Comment) (domain.Comment, error)
	FindAll() ([]domain.Comment, error)
	FindByID(id uint) (domain.Comment, error)
	Update(comment domain.Comment) (domain.Comment, error)
	Delete(id uint) error
}
