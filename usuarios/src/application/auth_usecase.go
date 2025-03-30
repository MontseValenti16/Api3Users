package auth_usecase

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"API3/usuarios/src/domain/entities"
	"API3/usuarios/src/domain/repositories"
)

// La llave secreta se carga desde la variable de entorno JWT_SECRET.
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type AuthUseCase struct {
	Repo repositories.UserRepository
}

func NewAuthUseCase(repo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{Repo: repo}
}

// Login valida las credenciales y, de ser correctas, genera y retorna un token JWT.
func (uc *AuthUseCase) Login(NombreUsuario, password string) (string, error) {
	user, err := uc.Repo.GetByUser(NombreUsuario)
	if err != nil {
		return "", err
	}
	// En un entorno real, comparar la contraseña con su hash (ej. bcrypt)
	if password != user.Password {
		return "", errors.New("credenciales inválidas")
	}

	claims := jwt.MapClaims{
		"user_id":        user.IDUsuario,
		"nombre_usuario": user.NombreUsuario,
		"exp":            time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// RegisterUser crea un nuevo usuario en la tabla usuario.
// Se espera que el front envíe también el id_persona obtenido tras registrar la persona.
func (uc *AuthUseCase) RegisterUser(user entities.User) (entities.User, error) {
	// Aquí podrías validar si el email ya existe, etc.
	return uc.Repo.CreateUser(user)
}
