package usecase

import (
	"mygram/domain"
	"mygram/interfaces"
	"mygram/utils"
)

type UserUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(userRepo interfaces.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo}
}

func (u *UserUseCase) Register(user domain.User) (domain.User, error) {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = hashedPassword
	return u.userRepo.Save(user)
}

func (u *UserUseCase) Login(email, password string) (domain.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return domain.User{}, err
	}

	return user, nil
}

func (u *UserUseCase) Update(user domain.User) (domain.User, error) {
	return u.userRepo.Update(user)
}

func (u *UserUseCase) Delete(id uint) error {
	return u.userRepo.Delete(id)
}
