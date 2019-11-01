package models

type MetaRubroNecesidad struct {
	Id               int                    `json:"Id"`
	MetaId           string                 `json:"MetaId"`
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
