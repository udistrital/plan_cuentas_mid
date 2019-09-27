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

// DocumentoPresupuestal ... estructura para guardar información de documentos presupuestales.
type DocumentoPresupuestal struct {
	ID                   string            `json:"Codigo" bson:"_id,omitempty"`
	Data                 interface{}       `json:"Data" bson:"Data" validate:"required"`
	Tipo                 string            `json:"Tipo" bson:"Tipo" validate:"required"`
	AfectacionIds        []string          `json:"AfectacionIds" bson:"AfectacionIds"`
	AfectacionMovimiento []Movimiento      `json:"AfectacionMovimiento" validate:"required"`
	Afectacion           []MovimientoMongo `json:"Afectacion"`
	FechaRegistro        string
}
