package models

import "time"

// Movimiento ... define the movimiento mase structure.
type Movimiento struct {
	Id                         int
	MovimientoProcesoExternoId *MovimientoProcesoExterno
	Valor                      float64
	FechaRegistro              time.Time
	Descripcion                string
}

// TipoMovimiento ... define the TipoMovimiento struct for movimiento_crud api.
type TipoMovimiento struct {
	Id          int
	Nombre      string
	Descripcion string
	Acronimo    string
}

type MovimientoProcesoExterno struct {
	Id                       int
	TipoMovimientoId         *TipoMovimiento
	ProcesoExterno           int64
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
