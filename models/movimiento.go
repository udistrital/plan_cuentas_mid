package models

import "time"

// Movimiento ... define the movimiento mase structure.
type Movimiento struct {
	Id                         int
	MovimientoProcesoExternoId *MovimientoProcesoExterno `validate:"required"`
	Valor                      float64                   `validate:"required"`
	FechaRegistro              time.Time
	Descripcion                string `validate:"required"`
}

// TipoMovimiento ... define the TipoMovimiento struct for movimiento_crud api.
type TipoMovimiento struct {
	Id          int `validate:"required"`
	Nombre      string
	Descripcion string
	Acronimo    string `validate:"required"`
}

type MovimientoProcesoExterno struct {
	Id                       int
	TipoMovimientoId         *TipoMovimiento `validate:"required"`
	ProcesoExterno           int64           `validate:"required"`
	MovimientoProcesoExterno int
}

type MovimientoMongo struct {
	ID             string  `json:"_id" bson:"_id,omitempty"`
	IDPsql         int     `json:"idpsql"`
	Valor          float64 `json:"valor"`
	Tipo           string  `json:"tipo"`
	DocumentoPadre int64   `json:"documento_padre"`
	FechaRegistro  string  `json:"fecha_registro"`
	Descripcion    string  `json:"unidad_ejecutora"`
}
