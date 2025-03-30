package repositories

import "API3/usuarios/src/domain/entities"

type UserRepository interface {
	GetByUser(NombreUsuario string) (entities.User, error)
	CreateUser(user entities.User) (entities.User, error)
}
