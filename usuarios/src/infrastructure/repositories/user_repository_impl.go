package repositories

import (
	"errors"
	"API3/usuarios/src/domain/entities"
)

// Usuarios simulados (en un entorno real se consultaría una base de datos)
var users = []entities.User{
	{ID: 1, Email: "test@example.com", Password: "password123"}, // Contraseña en claro para fines de ejemplo
}

type UserRepositoryImpl struct{}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetByEmail(email string) (entities.User, error) {
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}
	return entities.User{}, errors.New("usuario no encontrado")
}
