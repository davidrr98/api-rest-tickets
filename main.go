package main

import (
	"log"

	"github.com/davidrr98/api-rest-tickets/bd"
	"github.com/davidrr98/api-rest-tickets/handlers"
)

func main() {
	
	//Hacemos conexion con BD y revisamos la conexion 
	if !bd.ProbarConexion() {
		log.Fatal("Sin conexion a la BD")
		return
	}
	err := bd.CreateTabla()
	if(err!=nil){
		log.Fatal(err)
	}
	handlers.Manejadores()
}
