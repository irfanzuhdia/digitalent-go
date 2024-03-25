package usecase

import (
	"mygram/domain"
	"mygram/interfaces"
)

type CommentUseCase struct {
	commentRepo interfaces.CommentRepository
}

func NewCommentUseCase(commentRepo interfaces.CommentRepository) *CommentUseCase {
	return &CommentUseCase{commentRepo}
}

func (c *CommentUseCase) Create(comment domain.Comment) (domain.Comment, error) {
	return c.commentRepo.Save(comment)
}

func (c *CommentUseCase) GetAll() ([]domain.Comment, error) {
	return c.commentRepo.FindAll()
}

func (c *CommentUseCase) GetByID(id uint) (domain.Comment, error) {
	return c.commentRepo.FindByID(id)
}

func (c *CommentUseCase) Update(comment domain.Comment) (domain.Comment, error) {
	return c.commentRepo.Update(comment)
}

func (c *CommentUseCase) Delete(id uint) error {
	return c.commentRepo.Delete(id)
}
