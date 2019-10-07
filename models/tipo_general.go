package models

// TipoGeneral ... define the TipoGeneral struct for standar definition.
type TipoGeneral struct {
	Id          int `validate:"required"`
	Nombre      string
	Descripcion string
	Acronimo    string `validate:"required"`
}
