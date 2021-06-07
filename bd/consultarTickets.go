package bd

import (
	"context"
	"time"

	"github.com/davidrr98/api-rest-tickets/models"
)

func ConsultarTickets(codiciones string) ([]models.Ticket, error) {
	q := `
	SELECT id, usuario, estatus, fecha_creacion, fecha_modificacion
	FROM public.ticket ` + codiciones

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := DBConect.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tickets []models.Ticket
	for rows.Next() {
		var t models.Ticket
		rows.Scan(&t.ID, &t.Usuario, &t.Estatus, &t.FechaCreacion,
			&t.FechaModificacion)
		/* p.FechaInicio = p.FechaInicio[:10]
		p.FechaFin = p.FechaFin[:10] */
		tickets = append(tickets, t)
	}

	return tickets, nil

}

func ConsultarRegistroApoyo(codiciones string) ([]models.RegistroApoyo, error) {
	q := `
	SELECT id, tercero_id, solicitud_id, espacio_fisico_id, periodo_id, activo, usuario_administrador, fecha_creacion, fecha_modificacion
	FROM public.registro_apoyo ` + codiciones

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := DBConect.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var registros []models.RegistroApoyo
	for rows.Next() {
		var t models.RegistroApoyo
		rows.Scan(&t.ID, &t.TerceroId, &t.SolicitudId,
			&t.EspacioFisicoId,
			&t.PeriodoId,
			&t.Activo,
			&t.UsuarioAdministrador,
			&t.FechaCreacion,
			&t.FechaModificacion)
		/* p.FechaInicio = p.FechaInicio[:10]
		p.FechaFin = p.FechaFin[:10] */
		registros = append(registros, t)
	}

	return registros, nil

}

func ConsultarTicket(id int) (models.Ticket, error) {
	q := `SELECT id, usuario, estatus, fecha_creacion, fecha_modificacion
	FROM public.ticket where ticket.id=$1 LIMIT 1;`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var t models.Ticket

	rows, err := DBConect.QueryContext(ctx, q, id)
	if err != nil {
		return t, err
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&t.ID, &t.Usuario, &t.Estatus, &t.FechaCreacion,
			&t.FechaModificacion)
		/* p.FechaInicio = p.FechaInicio[:10]
		p.FechaFin = p.FechaFin[:10] */
	}
	return t, nil
}

/*
func ActulizarTicket(u models.Ticket) (string, error) {
	sqlStatement := `
INSERT INTO ticket (nombre, estado, fecha_inicio, fecha_fin)
VALUES ($1, $2, $3, $4)
RETURNING id_ticket`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConect
	id := 0
	err := db.QueryRowContext(ctx, sqlStatement, u.Nombre, u.Estado, u.FechaInicio, u.FechaFin).Scan(&id)

	return strconv.Itoa(id), err
} */
