package repositories

import (
	"API3/core/mysql"
	"database/sql"
	"errors"
	"API3/usuarios/src/domain/entities"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: mysql.DB, 
	}
}


func (repo *UserRepositoryImpl) GetByUser(NombreUsuario string) (entities.User, error) {
	query := `SELECT id_usuario, id_persona, nombre_usuario, password
	          FROM usuario
	          WHERE nombre_usuario = ?`
	row := repo.db.QueryRow(query, NombreUsuario)

	var user entities.User
	err := row.Scan(
		&user.IDUsuario,
		&user.IDPersona,
		&user.NombreUsuario,
		&user.Password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, errors.New("usuario no encontrado")
		}
		return entities.User{}, err
	}
	return user, nil
}

// CreateUser inserta un nuevo registro en la tabla usuario y retorna el usuario creado.
func (repo *UserRepositoryImpl) CreateUser(user entities.User) (entities.User, error) {
	query := `INSERT INTO usuario (nombre_usuario,  password, id_persona )
	          VALUES (?, ?, ?)`
	result, err := repo.db.Exec(query,  user.NombreUsuario, user.Password, user.IDPersona,)
	if err != nil {
		return entities.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entities.User{}, err
	}
	user.IDUsuario = int(id)
	return user, nil
}
