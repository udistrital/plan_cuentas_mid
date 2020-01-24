package models

import "time"

// ModificacionPresupuestalReceiver ...
type ModificacionPresupuestalReceiver struct {
	Data       *ModificacionPresupuestalReceiverDetail `json:"detail"`
	Afectation []*ModificacionPresupuestalReceiverAfectation
}

// ModificacionPresupuestalReceiverDetail ...
type ModificacionPresupuestalReceiverDetail struct {
	DocumentNumber  string       `json:"NumeroDocumento" bson:"numero_documento"`
	DocumentDate    time.Time    `json:"FechaDocumento" bson:"fecha_documento"`
	DocumentType    *TipoGeneral `json:"TipoDocumento" bson:"tipo_documento"`
	Descripcion     string       `json:"Descripcion" bson:"descripcion_documento"`
	CentroGestor    string       `json:"CentroGestor" bson:"-"`
	OrganismoEmisor string       `json:"OrganismoEmisor" bson:"organismo_emisor"`
}

// ModificacionPresupuestalReceiverAfectation ...
type ModificacionPresupuestalReceiverAfectation struct {
	OriginAcc *Rubro       `json:"CuentaCredito"`
	TargetAcc *Rubro       `json:"CuentaContraCredito"`
	TypeMod   *TipoGeneral `json:"Tipo"`
	Amount    float64      `json:"Valor"`
}

// ModificacionPresupuestalResponseDetail ...
type ModificacionPresupuestalResponseDetail struct {
	ID               string  `json:"_id" bson:"_id"`
	DocumentNumber   string  `json:"NumeroDocumento" bson:"numero_documento"`
	DocumentDate     string  `json:"FechaDocumento" bson:"fecha_documento"`
	DocumentType     string  `json:"TipoDocumento" bson:"tipo_documento"`
	Descripcion      string  `json:"Descripcion" bson:"descripcion_documento"`
	CentroGestor     string  `json:"CentroGestor" bson:"-"`
	OrganismoEmisor  string  `json:"OrganismoEmisor" bson:"organismo_emisor"`
	RegistrationDate string  `json:"FechaRegistro" bson:"fecha_regitro"`
	Vigencia         int     `json:"Vigencia" bson:"vigencia"`
	ValorActual      float64 `json:"ValorActual" bson:"valor_actual"`
	ValorInicial     float64 `json:"ValorInicial" bson:"valor_inicial"`
}
