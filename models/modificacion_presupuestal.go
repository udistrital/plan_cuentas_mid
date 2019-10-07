package models

import "time"

type ModificacionPresupuestalReceiver struct {
	Data       *ModificacionPresupuestalReceiverDetail `json:"detail"`
	Afectation []*ModificacionPresupuestalReceiverAfectation
}

type ModificacionPresupuestalReceiverDetail struct {
	DocumentNumber int          `json:"nDocumento" bson:"numero_documento"`
	DocumentDate   time.Time    `json:"fDocumento" bson:"fecha_documento"`
	DocumentType   *TipoGeneral `json:"tipoDocumento" bson:"tipo_documento"`
	Descripcion    string       `json:"descripcion" bson:"descripcion_documento"`
	CentroGestor   string       `json:"CentroGestor" bson:"-"`
}

type ModificacionPresupuestalReceiverAfectation struct {
	OriginAcc *Rubro       `json:"CuentaCredito"`
	TargetAcc *Rubro       `json:"CuentaContraCredito"`
	TypeMod   *TipoGeneral `json:"Tipo"`
	Amount    float64      `json:"Valor"`
}
