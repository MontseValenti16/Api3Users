package repositories

import (
	core "API3/core/mysql"
	"API3/src/registro/domain/entities"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type PersonaRepositoryImpl struct {
	db *core.Conn_MySQL
}

func NewPersonaRepository() *PersonaRepositoryImpl {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &PersonaRepositoryImpl{db: conn}
}
func (repo *PersonaRepositoryImpl) Create(p entities.User) (entities.User, error) {
	query := `INSERT INTO persona (nombre, apellidoP, estatura, fecha_nac, lista_invitado, apellidoM) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	result, err := repo.db.ExecutePreparedQuery(query, p.Nombre, p.ApellidoP, p.Estatura, p.FechaNac, p.ListaInvitados, p.ApellidoM)
	if err != nil {
		return entities.User{}, err
	}
	id, _ := result.LastInsertId()
	p.IDPersona = int(id)
	return p, nil
}

// GetByID busca una persona por su ID.
func (repo *PersonaRepositoryImpl) GetByID(id int) (entities.User, error) {
	query := `SELECT id_persona, nombre, apellidoP, estatura, fecha_nac, lista_invitado, apellidoM
	          FROM persona
	          WHERE id_persona = ?`
	row := repo.db.DB.QueryRow(query, id)

	var persona entities.User
	err := row.Scan(
		&persona.IDPersona,
		&persona.Nombre,
		&persona.ApellidoP,
		&persona.Estatura,
		&persona.FechaNac,
		&persona.ListaInvitados,
		&persona.ApellidoM,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.User{}, errors.New("persona no encontrada")
		}
		return entities.User{}, err
	}
	return persona, nil
}

// GetAll retorna todas las personas registradas.
func (repo *PersonaRepositoryImpl) GetAll() ([]entities.User, error) {
	query := `SELECT id_persona, nombre, apellidoP, estatura, fecha_nac, lista_invitado, apellidoM
	          FROM persona`
	rows, err := repo.db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var personas []entities.User
	for rows.Next() {
		var p entities.User
		if err := rows.Scan(
			&p.IDPersona,
			&p.Nombre,
			&p.ApellidoP,
			&p.Estatura,
			&p.FechaNac,
			&p.ListaInvitados,
			&p.ApellidoM,
		); err != nil {
			return nil, err
		}
		personas = append(personas, p)
	}
	return personas, nil
}

// Update actualiza los campos de una persona existente.
func (repo *PersonaRepositoryImpl) Update(p entities.User) (entities.User, error) {
	query := `UPDATE persona
	          SET nombre = ?, apellidoP = ?, estatura = ?, fecha_nac = ?, lista_invitado = ? , apellidoM = ?
	          WHERE id_persona = ?`
	result, err := repo.db.DB.Exec(query, p.Nombre, p.ApellidoP, p.Estatura, p.FechaNac, p.ListaInvitados, p.ApellidoM, p.IDPersona)
	if err != nil {
		return entities.User{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return entities.User{}, fmt.Errorf("no se encontró la persona con id %d", p.IDPersona)
	}
	return p, nil
}

// Delete elimina una persona por su ID.
func (repo *PersonaRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM persona WHERE id_persona = ?`
	result, err := repo.db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró la persona con id %d", id)
	}
	return nil
}
