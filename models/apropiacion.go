package models

//Apropiacion ...
type Apropiacion struct {
	Id              int                `orm:"column(id);pk"`
	Vigencia        float64            `orm:"column(vigencia);null"`
	RubroId           *Rubro             `orm:"column(rubro_id);rel(fk)"`
	UnidadEjecutora int                `orm:"column(unidad_ejecutora);null"`
	Valor           float64            `orm:"column(valor);null"`
	EstadoApropiacionId          *EstadoApropiacion `orm:"column(estado_apropiacion_id);rel(fk)"`
}
