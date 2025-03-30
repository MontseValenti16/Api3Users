package repositories

import (
	"API3/core/mysql"
	"database/sql"
	"errors"
	"fmt"
	"API3/registro/src/domain/entities"
)

type PersonaRepositoryImpl struct {
	db *sql.DB
}

func NewPersonaRepository() *PersonaRepositoryImpl {
	return &PersonaRepositoryImpl{
		db: mysql.DB, // Reutilizamos la conexión global inicializada en core/mysql
	}
}

// Create inserta un nuevo registro en la tabla persona y retorna la persona creada.
func (repo *PersonaRepositoryImpl) Create(p entities.User) (entities.User, error) {
	query := `INSERT INTO persona (nombre, apellido, edad, fecha_nac, lista_invitado) 
	          VALUES (?, ?, ?, ?, ?)`
	result, err := repo.db.Exec(query, p.Nombre, p.Apellido, p.Edad, p.FechaNac, p.ListaInvitados)
	if err != nil {
		return entities.User{}, err
	}
	id, _ := result.LastInsertId()
	p.IDPersona = int(id)
	return p, nil
}

// GetByID busca una persona por su ID.
func (repo *PersonaRepositoryImpl) GetByID(id int) (entities.User, error) {
	query := `SELECT id_persona, nombre, apellido, edad, fecha_nac, lista_invitado
	          FROM persona
	          WHERE id_persona = ?`
	row := repo.db.QueryRow(query, id)

	var persona entities.User
	err := row.Scan(
		&persona.IDPersona,
		&persona.Nombre,
		&persona.Apellido,
		&persona.Edad,
		&persona.FechaNac,
		&persona.ListaInvitados,
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
	query := `SELECT id_persona, nombre, apellido, edad, fecha_nac, lista_invitado
	          FROM persona`
	rows, err := repo.db.Query(query)
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
			&p.Apellido,
			&p.Edad,
			&p.FechaNac,
			&p.ListaInvitados,
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
	          SET nombre = ?, apellido = ?, edad = ?, fecha_nac = ?, lista_invitado = ?
	          WHERE id_persona = ?`
	result, err := repo.db.Exec(query, p.Nombre, p.Apellido, p.Edad, p.FechaNac, p.ListaInvitados, p.IDPersona)
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
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró la persona con id %d", id)
	}
	return nil
}
