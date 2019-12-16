package models

// DocumentoPresupuestal ... estructura para guardar información de documentos presupuestales.
type DocumentoPresupuestal struct {
	ID                   string            `json:"_id" bson:"_id,omitempty"`
	Data                 interface{}       `json:"Data" bson:"Data" validate:"required"`
	Tipo                 string            `json:"Tipo" bson:"Tipo" validate:"required"`
	AfectacionIds        []string          `json:"AfectacionIds" bson:"AfectacionIds"`
	Vigencia             int               `json:"Vigencia" bson:"Vigencia" validate:"required"`
	CentroGestor         string            `json:"CentroGestor" bson:"CentroGestor" validate:"required"`
	AfectacionMovimiento []Movimiento      `json:"AfectacionMovimiento" validate:"required"`
	Afectacion           []MovimientoMongo `json:"Afectacion"`
	FechaRegistro        string
	Consecutivo          int     `json:"Consecutivo" bson:"consecutivo"`
	ValorActual          float64 `json:"ValorActual" bson:"valor_actual"`
	ValorInicial         float64 `json:"ValorInicial" bson:"valor_inicial"`
}
