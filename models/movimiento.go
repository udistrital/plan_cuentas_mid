package models

// Movimiento ... define the movimiento mase structure.
type Movimiento struct {
	Id              int
	Tipo            *TipoMovimiento        `validate:"required"`
	UnidadEjecutora int                    `validate:"required"`
	Afectacion      map[string]interface{} `validate:"required"`
	MovimientoPadre *Movimiento
}

// TipoMovimiento ... define the TipoMovimiento struct for movimiento_crud api.
type TipoMovimiento struct {
	Id          int `validate:"required"`
	Nombre      string
	Descripcion string
	Acronimo    string `validate:"required"`
}
