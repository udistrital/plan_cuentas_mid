package models

import (
	"time"

	necesidades_crud "github.com/udistrital/necesidades_crud/models"
)

// MetaRubroNecesidad ...
type MetaRubroNecesidad struct {
	Id               int                    `json:"Id"`
	MetaId           string                 `json:"MetaId"`
	RubroNecesidadId map[string]interface{} `json:"RubroNecesidadId"`
	Actividades      []map[string]interface{}
}

// RubroNecesidad info del rubro
type RubroNecesidad struct {
	Id          int                      `json:"Id"`
	RubroId     string                   `json:"RubroId"`
	InfoRubro   *map[string]interface{}  `json:"InfoRubro"`
	NecesidadId map[string]interface{}   `json:"NecesidadId"`
	Fuentes     []map[string]interface{} `json:"Fuentes"`
	Productos   []map[string]interface{} `json:"Productos"`
	Metas       []*MetaRubroNecesidad    `json:"Metas"`
}

// TrNecesidad información completa de la necesidad
type TrNecesidad struct {
	Necesidad                          *map[string]interface{}  `json:"Necesidad" bson:"Necesidad"`
	DetalleServicioNecesidad           *map[string]interface{}  `json:"DetalleServicioNecesidad"`
	DetallePrestacionServicioNecesidad *map[string]interface{}  `json:"DetallePrestacionServicioNecesidad"`
	ProductosCatalogoNecesidad         []map[string]interface{} `json:"ProductosCatalogoNecesidad"`
	MarcoLegalNecesidad                []map[string]interface{} `json:"MarcoLegalNecesidad"`
	ActividadEspecificaNecesidad       []map[string]interface{} `json:"ActividadEspecificaNecesidad"`
	ActividadEconomicaNecesidad        []map[string]interface{} `json:"ActividadEconomicaNecesidad"`
	RequisitosMinimos                  []map[string]interface{} `json:"RequisitoMinimoNecesidad"`
	Rubros                             []*RubroNecesidad        `json:"Rubros"`
}

type NecesidadesCrudNecesidadCompleta struct {
	necesidades_crud.Necesidad
	// TODO: Si Beego lo llega a soportar al generar el Swagger,
	// redefinir este type simplemente como
	// type NecesidadesCrudNecesidadCompleta necesidades_crud.Necesidad
	// o bien, eliminar todos los campos siguientes

	Id                          int                                         `orm:"column(id);pk;auto"`
	Consecutivo                 int                                         `orm:"column(consecutivo);null"`
	Vigencia                    string                                      `orm:"column(vigencia)"`
	Objeto                      string                                      `orm:"column(objeto)"`
	Justificacion               string                                      `orm:"column(justificacion);null"`
	EstudioMercado              string                                      `orm:"column(estudio_mercado);null"`
	AnalisisRiesgo              string                                      `orm:"column(analisis_riesgo);null"`
	FechaSolicitud              time.Time                                   `orm:"column(fecha_solicitud);type(timestamp without time zone)"`
	Valor                       float64                                     `orm:"column(valor)"`
	AreaFuncional               int                                         `orm:"column(area_funcional)"`
	TipoDuracionNecesidadId     *necesidades_crud.TipoDuracionNecesidad     `orm:"column(tipo_duracion_necesidad_id);rel(fk)"`
	DiasDuracion                int                                         `orm:"column(dias_duracion);null"`
	ModalidadSeleccionId        *necesidades_crud.ModalidadSeleccion        `orm:"column(modalidad_seleccion_id);rel(fk)"`
	TipoContratoId              int                                         `orm:"column(tipo_contrato_id);null"`
	PlanAnualAdquisicionesId    int                                         `orm:"column(plan_anual_adquisiciones_id)"`
	TipoContratoNecesidadId     *necesidades_crud.TipoContratoNecesidad     `orm:"column(tipo_contrato_necesidad_id);rel(fk)"`
	TipoFinanciacionNecesidadId *necesidades_crud.TipoFinanciacionNecesidad `orm:"column(tipo_financiacion_necesidad_id);rel(fk)"`
	TipoNecesidadId             *necesidades_crud.TipoNecesidad             `orm:"column(tipo_necesidad_id);rel(fk)"`
	JustificacionRechazo        int                                         `orm:"column(justificacion_rechazo);null"`
	DependenciaNecesidadId      *necesidades_crud.DependenciaNecesidad      `orm:"column(dependencia_necesidad_id);rel(fk)"`
	EstadoNecesidadId           *necesidades_crud.EstadoNecesidad           `orm:"column(estado_necesidad_id);rel(fk)"`
	Activo                      bool                                        `orm:"column(activo)"`
	FechaCreacion               time.Time                                   `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	FechaModificacion           time.Time                                   `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
}

type NecesidadesCrudNecesidadParcial1 struct {
	necesidades_crud.Necesidad
	// TODO: Si Beego lo llega a soportar al generar el Swagger,
	// redefinir este type simplemente como
	// type NecesidadesCrudNecesidadParcial1 necesidades_crud.Necesidad
	// o bien, eliminar todos los campos siguientes

	Id int `orm:"column(id);pk;auto"`
	// ConsecutivoSolicitud        int                        `orm:"column(consecutivo_solicitud)"`
	// ConsecutivoNecesidad        int                        `orm:"column(consecutivo_necesidad);null"`
	Vigencia string `orm:"column(vigencia)"`
	// Objeto                      string                     `orm:"column(objeto)"`
	// Justificacion               string                     `orm:"column(justificacion);null"`
	// EstudioMercado              string                     `orm:"column(estudio_mercado);null"`
	// AnalisisRiesgo              string                     `orm:"column(analisis_riesgo);null"`
	// FechaSolicitud              time.Time                  `orm:"column(fecha_solicitud);type(timestamp without time zone)"`
	// Valor                       float64                    `orm:"column(valor)"`
	AreaFuncional int `orm:"column(area_funcional)"`
	// TipoDuracionNecesidadId     *necesidades_crud.TipoDuracionNecesidad     `orm:"column(tipo_duracion_necesidad_id);rel(fk)"`
	// DiasDuracion                int                        `orm:"column(dias_duracion);null"`
	// ModalidadSeleccionId        *necesidades_crud.ModalidadSeleccion        `orm:"column(modalidad_seleccion_id);rel(fk)"`
	// TipoContratoId              int                        `orm:"column(tipo_contrato_id);null"`
	// PlanAnualAdquisicionesId    int                        `orm:"column(plan_anual_adquisiciones_id)"`
	// TipoContratoNecesidadId     *necesidades_crud.TipoContratoNecesidad     `orm:"column(tipo_contrato_necesidad_id);rel(fk)"`
	// TipoFinanciacionNecesidadId *necesidades_crud.TipoFinanciacionNecesidad `orm:"column(tipo_financiacion_necesidad_id);rel(fk)"`
	// TipoNecesidadId             *necesidades_crud.TipoNecesidad             `orm:"column(tipo_necesidad_id);rel(fk)"`
	// JustificacionRechazo        int                        `orm:"column(justificacion_rechazo);null"`
	// DependenciaNecesidadId      *necesidades_crud.DependenciaNecesidad      `orm:"column(dependencia_necesidad_id);rel(fk)"`
	// EstadoNecesidadId           *necesidades_crud.EstadoNecesidad           `orm:"column(estado_necesidad_id);rel(fk)"`
	// Activo                      bool                       `orm:"column(activo)"`
	// FechaCreacion               time.Time                  `orm:"auto_now_add;column(fecha_creacion);type(timestamp without time zone)"`
	// FechaModificacion           time.Time                  `orm:"auto_now;column(fecha_modificacion);type(timestamp without time zone)"`
}
