package models

// Rama ...
type Rama struct {
	Id         int    `orm:"column(id)"`
	RubroPadre *Rubro `orm:"column(rubro_padre);rel(fk)"`
	RubroHijo  *Rubro `orm:"column(rubro_hijo);rel(fk)"`
}
