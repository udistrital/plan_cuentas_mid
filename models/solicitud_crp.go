package models

import (
	"time"
)

// SolicitudCrp ...
type SolicitudCrp struct {
	ID                string      `json:"_id" bson:"_id,omitempty"`
	Consecutivo       int         `json:"consecutivo" bson:"consecutivo"`
	ConsecutivoCDP    int         `json:"consecutivoCdp" bson:"consecutivoCdp"`
	Vigencia          string      `json:"vigencia" bson:"vigencia"`
	Beneficiario      string      `json:"beneficiario" bson:"beneficiario"`
	Valor             float64     `json:"valor" bson:"valor"`
	Compromiso        *Compromiso `json:"compromiso" bson:"compromiso"`
	InfoCRP           *infoCrp    `json:"infoCrp" bson:"infoCrp"`
	Activo            bool        `json:"activo" bson:"activo"`
	FechaCreacion     time.Time   `json:"fechaCreacion" bson:"fechaCreacion"`
	FechaModificacion time.Time   `json:"fechaModificacion" bson:"fechaModificacion"`
}
