package models

type FuenteFinanciamientoApropiacion struct {
	Id                     int                   `orm:"column(id);pk"`
	ApropiacionId          *Apropiacion          `orm:"column(apropiacion_id);rel(fk)"`
	FuenteFinanciamientoId *FuenteFinanciamiento `orm:"column(fuente_financiamiento_id);rel(fk)"`
	Dependencia            int                   `orm:"column(dependencia)"`
}
type ModificacionFuenteReceiver struct {
	Data       *ModificacionPresupuestalReceiverDetail `json:"detail"`
	Afectation []*ModificacionFuenteReceiverAfectation
}

type ModificacionFuenteReceiverAfectation struct {
	OriginAcc *FuenteFinanciamiento `json:"MovimientoOrigen"`
	TargetAcc *FuenteFinanciamiento `json:"MovimientoDestino"`
	TypeMod   *TipoGeneral          `json:"Tipo"`
	Amount    float64               `json:"Valor"`
}
