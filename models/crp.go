package models

import (
	"time"
)

type estadoCrp struct {
	ID       string `json:"id"`
	Acronimo string `json:"acronimo"`
	Nombre   string `json:"nombre"`
}
// Compromiso ...
type Compromiso struct {
	NumeroCompromiso int `orm:"column(numeroCompromiso);pk"`
	TipoCompromiso   int `orm:"column(tipoCompromiso);"`
}

// SolicitudCrp ...
type SolicitudCrp struct {
	ID                string      `json:"_id" bson:"_id,omitempty"`
	Consecutivo       int         `json:"consecutivo" bson:"consecutivo"`
	ConsecutivoCDP    int         `json:"consecutivoCdp" bson:"consecutivoCdp"`
	Vigencia          string      `json:"vigencia" bson:"vigencia"`
	Beneficiario      string      `json:"beneficiario" bson:"beneficiario"`
	Valor             float64     `json:"valor" bson:"valor"`
	Compromiso        *Compromiso `json:"compromiso" bson:"compromiso"`
	Estado			  estadoCrp 	  `json:"estado"`
	Activo            bool        `json:"activo" bson:"activo"`
	FechaCreacion     time.Time   `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaInicioVigencia time.Time     `json:"fechaInicioVigencia"`
	FechaFinalVigencia time.Time     `json:"fechaFinalVigencia"`
	FechaModificacion time.Time   `json:"fechaModificacion" bson:"fechaModificacion"`
}


// GetEstadoSolicitudCrp devuelve el estado solicitud de una solicitud de crp
func GetEstadoSolicitudCrp() interface{} {
	estado := estadoCdp{
		ID:       "1",
		Acronimo: "sol",
		Nombre:   "solicitud",
	}
	return estado
}

// GetEstadoRechazadoCrp devuelve el estado rechazado de una solicitud de crp
func GetEstadoRechazadoCrp() interface{} {
	estado := estadoCdp{
		ID:       "2",
		Acronimo: "rec",
		Nombre:   "rechazado",
	}
	return estado
}

// GetEstadoExpedidoCrp devuelve el estado aprobado de una solicitud de crp
func GetEstadoExpedidoCrp() interface{} {
	estado := estadoCdp{
		ID:       "3",
		Acronimo: "exp",
		Nombre:   "expedido",
	}
	return estado
}