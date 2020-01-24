package models

// FuenteFinanciamientoApropiacion ...
type FuenteFinanciamientoApropiacion struct {
	Id                     int                   `orm:"column(id);pk"`
	ApropiacionId          *Apropiacion          `orm:"column(apropiacion_id);rel(fk)"`
	FuenteFinanciamientoId *FuenteFinanciamiento `orm:"column(fuente_financiamiento_id);rel(fk)"`
	Dependencia            int                   `orm:"column(dependencia)"`
}

// ModificacionFuenteReceiver ...
type ModificacionFuenteReceiver struct {
	Data       *ModificacionPresupuestalReceiverDetail `json:"detail"`
	Afectation []*ModificacionFuenteReceiverAfectation
}

// ModificacionFuenteReceiverAfectation ...
type ModificacionFuenteReceiverAfectation struct {
	OriginAcc   *FuenteFinanciamiento `json:"MovimientoOrigen"`
	OriginRubro *Rubro                `json:"CuentaCredito"`
	TypeMod     *TipoGeneral          `json:"Tipo"`
	Amount      float64               `json:"Valor"`
}
