
package models

import (
	"time"
)

// infoCrp
type infoCrp struct {
	Consecutivo     int       `json:"consecutivo" bson:"consecutivo"`
	FechaExpedicion time.Time `json:"fechaExpedicion" bson:"fechaExpedicion"`
	Estado          int       `json:"estado" bson:"estado"`
}