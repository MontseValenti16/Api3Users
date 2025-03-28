package repositories

import "API3/usuarios/src/domain/entities"

// Solo se requiere el método para obtener por email (puedes extenderlo según tus necesidades)
type UserRepository interface {
	GetByEmail(email string) (entities.User, error)
}
