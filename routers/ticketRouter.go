package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davidrr98/api-rest-tickets/bd"
	"github.com/davidrr98/api-rest-tickets/models"
	"github.com/gorilla/mux"
)

// HealthCheck comprueba que la api este funcionando
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// RegistroTicket Agrega un ticket a la base de datos
func RegistroTicket(w http.ResponseWriter, r *http.Request) {

	var t models.Ticket
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Usuario) == 0 {
		http.Error(w, "El nombre del ticket es requerido", 400)
		return
	}

	_, err = bd.InsertoTicket(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realziar el registro del ticket "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ConsultarTickets consulta la tabla tickets y la filtra
func ConsultarTickets(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	queries := []string{"id", "usuario", "estatus"}


	condiciones := ""

	for _ , s := range queries {
		valor := v.Get(s)
		if(valor!=""){
			if(condiciones!=""){
				condiciones = condiciones + " AND "
			}else{
				condiciones= condiciones + "WHERE "
			}
			condiciones = condiciones +   "ticket."+s+"="+valor
		}
	}
	
	limite := v.Get("limit") 
	if(limite!=""){
		condiciones += " LIMIT "+limite
	}

	tickets, err := bd.ConsultarTickets(condiciones)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realziar el registro del periodo "+err.Error(), 400)
		return
	}
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

/*  Edita un ticket */
func ActualizarTicket(w http.ResponseWriter, r *http.Request) {

	var t models.Ticket
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	err = bd.ActualizarTicket(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realzar la actualizacion del ticket "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// retorna un ticket dado por la id
func ConsultarTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	ID_num, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "El id ingresado no es valido", 400)
		return
	}
	ticket, err := bd.ConsultarTicket(ID_num)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro del ticket "+err.Error(), 400)
		return
	}
	if ticket.Usuario == "" {
		http.Error(w, "Ticket no encontrado", 400)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(ticket)
}

// Elimina un ticket por la id
func EliminarTicket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
	}
	ID_num, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "El id ingresado no es valido", 400)
		return
	}
	err = bd.EliminarTicket(ID_num)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el registro del ticket "+err.Error(), 400)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
