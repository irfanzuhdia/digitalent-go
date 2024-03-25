package repository

import (
	"mygram/domain"
	"mygram/interfaces"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Save(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *UserRepository) Update(user domain.User) (domain.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *UserRepository) Delete(id uint) error {
	var user domain.User
	err := r.db.Where("id = ?", id).Delete(&user).Error
	return err
}
