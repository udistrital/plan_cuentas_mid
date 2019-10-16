package models

import (
	"time"
)

type infoCdp struct {
	Consecutivo     int       `json:"consecutivo" bson:"consecutivo"`
	FechaExpedicion time.Time `json:"fechaExpedicion" bson:"fechaExpedicion"`
	Estado          int       `json:"estado" bson:"estado"`
}
// SolicitudCrp ...
type SolicitudCdp struct {
	ID                   string		   `json:"_id" bson:"_id,omitempty"`
	Consecutivo          int           `json:"consecutivo" bson:"consecutivo"`
	Entidad              int           `json:"entidad" bson:"entidad"`
	CentroGestor         int           `json:"centroGestor" bson:"centroGestor"`
	Necesidad            int           `json:"necesidad" bson:"necesidad"`
	Vigencia             string        `json:"vigencia" bson:"vigencia"`
	FechaRegistro        time.Time     `json:"fechaRegistro" bson:"fechaRegistro"`
	Estado               int           `json:"estado" bson:"estado"`
	JustificacionRechazo string        `json:"justificacionRechazo" bson:"justificacionRechazo"`
	InfoCDP              *infoCdp      `json:"infoCdp" bson:"infoCdp"`
	Activo               bool          `json:"activo" bson:"activo"`
	FechaCreacion        time.Time     `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaModificacion    time.Time     `json:"fechaModificacion" bson:"fechaModificacion"`
}
