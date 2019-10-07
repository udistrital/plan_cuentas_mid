package models

// Compromiso ...
type Compromiso struct {
	NumeroCompromiso int `orm:"column(numeroCompromiso);pk"`
	TipoCompromiso   int `orm:"column(tipoCompromiso);"`
}
