package repository

import (
	"mygram/domain"
	"mygram/interfaces"

	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) interfaces.CommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Save(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Create(&comment).Error
	return comment, err
}

func (r *CommentRepository) FindAll() ([]domain.Comment, error) {
	var comments []domain.Comment
	err := r.db.Preload("User").Preload("Photo").Find(&comments).Error
	return comments, err
}

func (r *CommentRepository) FindByID(id uint) (domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", id).First(&comment).Error
	return comment, err
}

func (r *CommentRepository) Update(comment domain.Comment) (domain.Comment, error) {
	err := r.db.Save(&comment).Error
	return comment, err
}

func (r *CommentRepository) Delete(id uint) error {
	var comment domain.Comment
	err := r.db.Where("id = ?", id).Delete(&comment).Error
	return err
}
