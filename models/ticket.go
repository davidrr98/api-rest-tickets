package models

type Ticket struct {
	ID                int    `json:"id"`
	Usuario           string `json:"usuario"`
	FechaCreacion     string `json:"fecha_creacion"`
	FechaModificacion string `json:"fecha_modificacion"`
	Estatus           bool   `json:"estatus"`
}

type RegistroApoyo struct {
	ID                   int    `json:"id"`
	TerceroId            int    `json:"terceroId"`
	SolicitudId          int    `json:"solicitudId"`
	PeriodoId            int    `json:"periodoId"`
	EspacioFisicoId      int    `json:"espacioFisicoId"`
	UsuarioAdministrador string `json:"usuarioAdministrador"`
	Activo               bool   `json:"activo"`
	FechaCreacion        string `json:"fecha_creacion"`
	FechaModificacion    string `json:"fecha_modificacion"`
}
