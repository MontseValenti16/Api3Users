package repositories

import "API3/src/registro/domain/entities"

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	GetByID(id int) (entities.User, error)
	GetAll() ([]entities.User, error)
	Update(user entities.User) (entities.User, error)
	Delete(id int) error
}
