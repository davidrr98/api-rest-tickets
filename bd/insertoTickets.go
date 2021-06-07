package bd

import (
	"context"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/davidrr98/api-rest-tickets/models"
)

func CreateTabla() error {
	//path := filepath.Join("modelo.sql")

	c, ioErr := ioutil.ReadFile("modelo.sql")
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)

	db := DBConect
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func InsertoTicket(u models.Ticket) (string, error) {
	sqlStatement := `INSERT INTO public.ticket(
		usuario, estatus, fecha_creacion, fecha_modificacion)
		VALUES  ($1, $2, $3, $4) RETURNING id;`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConect
	id := 0
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)
	err := db.QueryRowContext(ctx, sqlStatement, u.Usuario, u.Estatus, dt, dt).Scan(&id)

	return strconv.Itoa(id), err
}

func InsertoRegistro(u models.RegistroApoyo) (int, error) {

	sqlStatement := `INSERT INTO public.registro_apoyo(
		tercero_id, solicitud_id, espacio_fisico_id, periodo_id, activo, usuario_administrador, fecha_creacion, fecha_modificacion)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`

	log.Print(u)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := DBConect
	id := 0
	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)
	err := db.QueryRowContext(ctx, sqlStatement, u.TerceroId, u.SolicitudId, u.EspacioFisicoId, u.PeriodoId, u.Activo, u.UsuarioAdministrador, dt, dt).Scan(&id)
	return id, err
}

func ActualizarTicket(u models.Ticket) error {
	sqlStatement := `UPDATE ticket
	SET usuario=$1, estatus=$2, fecha_modificacion=$3
	WHERE id=$4;`

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	db := DBConect
	_, err := db.ExecContext(ctx, sqlStatement, u.Usuario, u.Estatus, dt, u.ID)

	return err
}

func EliminarTicket(id int) error {
	sqlStatement := `DELETE FROM ticket
	WHERE id=$1;`
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := DBConect
	_, err := db.ExecContext(ctx, sqlStatement, id)
	return err
}
