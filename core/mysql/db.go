package mysql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Aviso: no se pudo cargar el archivo .env o no existe. Se usarán variables de entorno existentes.")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", user, pass, host, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error al abrir la conexión a la base de datos: %v", err))
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		panic(fmt.Sprintf("No se pudo establecer conexión con la base de datos: %v", err))
	}

	DB = db
	fmt.Println("Conexión exitosa a la base de datos!")
}
