package models

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	Id                       int                       `orm:"column(id);pk"`
	Nombre                   string                    `orm:"column(nombre);null"`
	Descripcion              string                    `orm:"column(descripcion);null"`
	TipoFuenteFinanciamiento *TipoFuenteFinanciamiento `orm:"column(tipo_fuente_financiamiento);null"`
	Codigo                   string                    `orm:"column(codigo)"`
}

