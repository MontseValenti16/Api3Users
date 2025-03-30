package entities

import "time"

type User struct {
	IDPersona      int       `json:"id_persona"`
	Nombre         string    `json:"nombre"`
	Apellido       string    `json:"apellido"`
	Edad           int       `json:"edad"`
	FechaNac       time.Time `json:"fecha_nac"`
	ListaInvitados *int       `json:"lista_invitado"`
}
