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

// MovimientoProcesoExterno ...
type MovimientoProcesoExterno struct {
	Id                       int
	TipoMovimientoId         *TipoGeneral `validate:"required"`
	ProcesoExterno           int64
	MovimientoProcesoExterno int
}

// MovimientoMongo ...
type MovimientoMongo struct {
	ID            string  `json:"_id" bson:"_id,omitempty"`
	IDPsql        int     `json:"IDPsql"`
	ValorInicial  float64 `json:"ValorInicial"`
	Tipo          string  `json:"Tipo"`
	Padre         string  `json:"Padre"`
	FechaRegistro string  `json:"FechaRegistro"`
	Descripcion   string  `json:"Descripcion"`
}
