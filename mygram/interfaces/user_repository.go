package interfaces

import "mygram/domain"

type UserRepository interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
}
