package models

import "time"

// NecesidadADM objeto para tipar necesidad de administrativa
type NecesidadADM struct {
	Id                        int                      `json:"Id"`
	Numero                    int                      `json:"Numero"`
	Vigencia                  float64                  `json:"Vigencia"`
	Objeto                    string                   `json:"Objeto"`
	FechaSolicitud            time.Time                `json:"FechaSolicitud"`
	Valor                     float64                  `json:"Valor"`
	Justificacion             string                   `json:"Justificacion"`
	UnidadEjecutora           int                      `json:"UnidadEjecutora"`
	DiasDuracion              float64                  `json:"DiasDuracion"`
	UnicoPago                 bool                     `json:"UnicoPago"`
	AgotarPresupuesto         bool                     `json:"AgotarPresupuesto"`
	ModalidadSeleccion        map[string]interface{}   `json:"ModalidadSeleccion"`
	TipoContratoNecesidad     map[string]interface{}   `json:"TipoContratoNecesidad"`
	PlanAnualAdquisiciones    int                      `json:"PlanAnualAdquisiciones"`
	EstudioMercado            string                   `json:"EstudioMercado"`
	TipoFinanciacionNecesidad map[string]interface{}   `json:"TipoFinanciacionNecesidad"`
	Supervisor                int                      `json:"Supervisor"`
	AnalisisRiesgo            string                   `json:"AnalisisRiesgo"`
	NumeroElaboracion         int                      `json:"NumeroElaboracion"`
	FechaModificacion         time.Time                `json:"FechaModificacion"`
	EstadoNecesidad           map[string]interface{}   `json:"EstadoNecesidad"`
	JustificacionRechazo      string                   `json:"JustificacionRechazo"`
	JustificacionAnulacion    string                   `json:"JustificacionAnulacion"`
	TipoNecesidad             map[string]interface{}   `json:"TipoNecesidad"`
	FuenteReversa             []map[string]interface{} `json:"FuenteReversa"`
	DependenciaReversa        []map[string]interface{} `json:"DependenciaReversa"`
	ProductoReversa           []map[string]interface{} `json:"ProductoReversa"`
}

// Actividades asociadas a una meta
type actividad struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Metas de una necesidad
type meta struct {
	Codigo      string       `json:"codigo" bson:"codigo"`
	Actividades []*actividad `json:"actividades" bson:"actividades"`
}

// Apropiacion de la necesidad (es el que va a tener las metas)
type apropiacion struct {
	Codigo    string      `json:"codigo" bson:"codigo"`
	Metas     []*meta     `json:"metas" bson:"metas"`
	Fuentes   []*fuente   `json:"fuentes" bson:"fuentes"`
	Productos []*producto `json:"productos" bson:"productos"`
}

// Fuentes de la necesidad
type fuente struct {
	Codigo string  `json:"codigo" bson:"codigo"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Productos de la necesidad
type producto struct {
	Codigo string  `json:"_id" bson:"_id,omitempty"`
	Valor  float64 `json:"valor" bson:"valor"`
}

// Productos de la necesidad
type detalleServicio struct {
	Codigo      string  `json:"codigo" bson:"codigo"`
	Valor       float64 `json:"valor" bson:"valor"`
	Descripcion string  `json:"descripcion" bson:"descripcion"`
}

type MetaRubroNecesidad struct {
	Id               int                    `json:"Id"`
	MetaId           int                    `json:"MetaId"`
	RubroNecesidadId map[string]interface{} `json:"RubroNecesidadId"`
	Actividades      []*map[string]interface{}
}

// RubroNecesidad info del rubro
type RubroNecesidad struct {
	Id          int                       `json:"Id"`
	RubroId     string                    `json:"RubroId"`
	InfoRubro   *map[string]interface{}   `json:"InfoRubro"`
	NecesidadId map[string]interface{}    `json:"NecesidadId"`
	Fuentes     []*map[string]interface{} `json:"Fuentes"`
	Productos   []*map[string]interface{} `json:"Productos"`
	Metas       []*MetaRubroNecesidad     `json:"Metas"`
}

// TrNecesidad información completa de la necesidad
type TrNecesidad struct {
	Necesidad                          *map[string]interface{}   `json:"Necesidad" bson:"Necesidad"`
	DetalleServicioNecesidad           *map[string]interface{}   `json:"DetalleServicioNecesidad"`
	DetallePrestacionServicioNecesidad *map[string]interface{}   `json:"DetallePrestacionServicioNecesidad"`
	ProductosCatalogoNecesidad         []*map[string]interface{} `json:"ProductosCatalogoNecesidad"`
	MarcoLegalNecesidad                []*map[string]interface{} `json:"MarcoLegalNecesidad"`
	ActividadEspecificaNecesidad       []*map[string]interface{} `json:"ActividadEspecificaNecesidad"`
	ActividadEconomicaNecesidad        []*map[string]interface{} `json:"ActividadEconomicaNecesidad"`
	Rubros                             []*RubroNecesidad         `json:"Rubros"`
}
