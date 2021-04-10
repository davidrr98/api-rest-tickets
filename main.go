package main

import (
	"log"

	"github.com/davidrr98/api-rest-tickets/bd"
)

func main() {
	if !bd.ProbarConexion() {
		log.Fatal("Sin conexion a la BD")
		return
	}
	log.Output("Listo para trabajar")
	//	handlers.Manejadores()
}
