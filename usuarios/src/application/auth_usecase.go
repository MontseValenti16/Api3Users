package auth_usecase

import (
	"errors"
	"time"

	//"API3/usuarios/src/domain/entities"
	"API3/usuarios/src/domain/repositories"

	"github.com/dgrijalva/jwt-go"
)


var jwtKey = []byte("my_secret_key")

type AuthUseCase struct {
	Repo repositories.UserRepository
}

func NewAuthUseCase(repo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{Repo: repo}
}

// Login valida las credenciales y, de ser correctas, genera y retorna un token JWT.
func (uc *AuthUseCase) Login(email, password string) (string, error) {
	user, err := uc.Repo.GetByEmail(email)
	if err != nil {
		return "", err
	}
	// Simulación: en un entorno real se debe comparar con un hash de contraseña
	if password != user.Password {
		return "", errors.New("credenciales inválidas")
	}
	// Crear token JWT con claims
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
