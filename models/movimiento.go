package models

// Movimiento ... define la estructura de un movimiento dentro del dominio de la aplicacion
type Movimiento struct {
	Id              int
	Tipo            string                 `validate:"required"`
	UnidadEjecutora int                    `validate:"required"`
	Afectacion      map[string]interface{} `validate:"required"`
	MovimientoPadre *Movimiento
}
