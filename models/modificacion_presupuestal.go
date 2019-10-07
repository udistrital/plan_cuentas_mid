package models

import "time"

type ModificacionPresupuestalReceiver struct {
	Data       *ModificacionPresupuestalReceiverDetail `json:"detail"`
	Afectation []*ModificacionPresupuestalReceiverAfectation
}

type ModificacionPresupuestalReceiverDetail struct {
	DocumentNumber int          `json:"nDocumento"`
	DocumentDate   time.Time    `json:"fDocumento"`
	DocumentType   *TipoGeneral `json:"tipoDocumento"`
	Descripcion    string       `json:"descripcion"`
}

type ModificacionPresupuestalReceiverAfectation struct {
	OriginAcc *Rubro       `json:"CuentaCredito"`
	TargetAcc *Rubro       `json:"CuentaContraCredito"`
	TypeMod   *TipoGeneral `json:"Tipo"`
	Amount    float64      `json:"Valor"`
}
