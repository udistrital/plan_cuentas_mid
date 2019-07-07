package models

import "time"

// FuenteFinanciamiento ...
type FuenteFinanciamiento struct {
	Id                       int                       `orm:"column(id);pk"`
	Nombre                   string                    `orm:"column(nombre)"`
	Descripcion              string                    `orm:"column(descripcion);null"`
	FechaCreacion            time.Time                 `orm:"column(fecha_creacion);type(date);null"`
	TipoFuenteFinanciamiento *TipoFuenteFinanciamiento `orm:"column(tipo_fuente_financiamiento);rel(fk)"`
	Codigo                   string                    `orm:"column(codigo)"`
}
