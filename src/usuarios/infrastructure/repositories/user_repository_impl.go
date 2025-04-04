package repositories

import (
	core "API3/core/mysql"
	"API3/src/usuarios/domain/entities"
	"database/sql"
	"errors"
	"log"
)

type UserRepositoryImpl struct {
	db *core.Conn_MySQL
}

func NewUserRepository() *UserRepositoryImpl {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &UserRepositoryImpl{db: conn}
}


func (repo *UserRepositoryImpl) GetByUser(NombreUsuario string) (entities.User, error) {
	query := `SELECT id_usuario, id_persona, nombre_usuario, password
	          FROM usuario
	          WHERE nombre_usuario = ?`
	row := repo.db.DB.QueryRow(query, NombreUsuario)

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
	result, err := repo.db.DB.Exec(query,  user.NombreUsuario, user.Password, user.IDPersona,)
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
