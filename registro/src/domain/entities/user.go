package entities

import "time"

type User struct {
	IDPersona      int       `json:"id_persona"`
	Nombre         string    `json:"nombre"`
	ApellidoM      string    `json:"apellidom"`
	ApellidoP      string     `json:"apellidop"`
	Estatura       int       `json:"estatura"`
	FechaNac       time.Time `json:"fecha_nac"`
	ListaInvitados *int       `json:"lista_invitado"`
}
