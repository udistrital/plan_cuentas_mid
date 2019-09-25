package models

// Movimiento ... define the movimiento mase structure.
type Movimiento struct {
	Id                         int
	MovimientoProcesoExternoId *MovimientoProcesoExterno `validate:"required"`
	Valor                      float64                   `validate:"required"`
	FechaRegistro              string
	Descripcion                string `validate:"required"`
	DocumentoPadre             string
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
	ProcesoExterno           int64
	MovimientoProcesoExterno int
}

type MovimientoMongo struct {
	ID            string  `json:"_id" bson:"_id,omitempty"`
	IDPsql        int     `json:"IDPsql"`
	Valor         float64 `json:"Valor"`
	Tipo          string  `json:"Tipo"`
	Padre         string  `json:"Padre"`
	FechaRegistro string  `json:"FechaRegistro"`
	Descripcion   string  `json:"Descripcion"`
}
