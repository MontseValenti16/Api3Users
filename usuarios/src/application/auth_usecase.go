package auth_usecase

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"API3/usuarios/src/domain/entities"
	"API3/usuarios/src/domain/repositories"
)


var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type AuthUseCase struct {
	Repo repositories.UserRepository
}

func NewAuthUseCase(repo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{Repo: repo}
}

func (uc *AuthUseCase) Login(NombreUsuario, password string) (string, int, string,error) {
	user, err := uc.Repo.GetByUser(NombreUsuario)
	if err != nil {
		return "",0,"", err
	}
	// Comparar la contraseña usando bcrypt.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "",0,"", errors.New("credenciales inválidas")
	}

	// Crear el token JWT con claims.
	claims := jwt.MapClaims{
		"user_id":        user.IDUsuario,
		"nombre_usuario": user.NombreUsuario,
		"exp":            time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "",0,"", err
	}
	return tokenString, user.IDUsuario, user.NombreUsuario , nil
}


func (uc *AuthUseCase) RegisterUser(user entities.User) (entities.User, error) {
	// Verificar si ya existe el usuario (esto lo puedes agregar según tus necesidades).
	// Cifrar la contraseña.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}
	user.Password = string(hashedPassword)

	// Crear el usuario en el repositorio.
	return uc.Repo.CreateUser(user)
}
