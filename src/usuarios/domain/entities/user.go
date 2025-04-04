package entities

type User struct {
	IDUsuario     int    `json:"id_usuario"`
	IDPersona     int    `json:"id_persona"`    
	NombreUsuario string `json:"nombre_usuario"`
	Password      string `json:"password"`
}
