package bd

import (
	"database/sql"
	"fmt"
	"log"

	/*  */
	_ "github.com/lib/pq"
)

var DBConect = ConectarBD()

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "prueba_tecnica"
)

/* Variables de entorno */

/* ConectarBD es la funcion que me permite conectar la BD*/
func ConectarBD() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	err = db.Ping()
	// Create an empty user and make the sql query (using $1 for the parameter)
	log.Println("Conexion exitosa Base de Datos")
	return db
}

/* Prueba la coneccion a la base de datos */
func ProbarConexion() bool {
	err := DBConect.Ping()
	if err != nil {
		return false
	}
	return true

}
