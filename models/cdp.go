package models

import (
	"time"
)

type estadoCdp struct {
	ID       string `json:"id"`
	Acronimo string `json:"acronimo"`
	Nombre   string `json:"nombre"`
}

// SolicitudCDP información de la solicitud de un CDP
type SolicitudCDP struct {
	ID                   string    `json:"_id" bson:"_id,omitempty"`
	Consecutivo          int       `json:"consecutivo" bson:"consecutivo"`
	Entidad              int       `json:"entidad" bson:"entidad"`
	CentroGestor         int       `json:"centroGestor" bson:"centroGestor"`
	Necesidad            int       `json:"necesidad" bson:"necesidad"`
	Vigencia             string    `json:"vigencia" bson:"vigencia"`
	FechaRegistro        time.Time `json:"fechaRegistro" bson:"fechaRegistro"`
	Estado               estadoCdp `json:"estado" bson:"estado"`
	JustificacionRechazo string    `json:"justificacionRechazo" bson:"justificacionRechazo"`
	Activo               bool      `json:"activo" bson:"activo"`
	FechaCreacion        time.Time `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaModificacion    time.Time `json:"fechaModificacion" bson:"fechaModificacion"`
}

// GetEstadoSolicitudCdp devuelve el estado solicitud de una solicitud de cdp
func GetEstadoSolicitudCdp() interface{} {
	estado := estadoCdp{
		ID:       "1",
		Acronimo: "sol",
		Nombre:   "solicitud",
	}
	return estado
}

// GetEstadoRechazadoCdp devuelve el estado rechazado de una solicitud de cdp
func GetEstadoRechazadoCdp() interface{} {
	estado := estadoCdp{
		ID:       "2",
		Acronimo: "rec",
		Nombre:   "rechazado",
	}
	return estado
}

// GetEstadoExpedidoCdp devuelve el estado aprobado de una solicitud de cdp
func GetEstadoExpedidoCdp() interface{} {
	estado := estadoCdp{
		ID:       "3",
		Acronimo: "exp",
		Nombre:   "expedido",
	}
	return estado
}
