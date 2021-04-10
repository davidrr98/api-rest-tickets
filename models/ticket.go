package models

type Ticket struct {
	ID                int    `json:"id"`
	Usuario           string `json:"usuario"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_modificacion"`
	Estatus           bool   `json:"estatus"`
}
