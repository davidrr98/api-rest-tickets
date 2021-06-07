package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/davidrr98/api-rest-tickets/middlew"
	"github.com/davidrr98/api-rest-tickets/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores es una dependencia*/
func Manejadores() {
	router := mux.NewRouter()

	//Rutas
	router.HandleFunc("/health_check", middlew.ChequeoBD(routers.HealthCheck)).Methods("GET")
	router.HandleFunc("/ticket", middlew.ChequeoBD(routers.ConsultarTickets)).Methods("GET")

	router.HandleFunc("/ticket/{id}", middlew.ChequeoBD(routers.ConsultarTicket)).Methods("GET")
	router.HandleFunc("/ticket", middlew.ChequeoBD(routers.RegistroTicket)).Methods("POST")
	router.HandleFunc("/ticket", middlew.ChequeoBD(routers.ActualizarTicket)).Methods("PUT")
	router.HandleFunc("/ticket/{id}", middlew.ChequeoBD(routers.EliminarTicket)).Methods("DELETE")

	router.HandleFunc("/registro_apoyo", middlew.ChequeoBD(routers.ConsultarRegistroDiario)).Methods("GET")
	router.HandleFunc("/registro_apoyo", middlew.ChequeoBD(routers.AgregarRegistroDiario)).Methods("POST")

	/* Swagger documentacion*/

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
