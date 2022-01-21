// Modelos para completar Swagger

package models

import (
	plan_cuentas_mongo_crud "github.com/udistrital/plan_cuentas_mongo_crud/models"
)

// --------------------------------------------------------------------------------------------
// AprobacionController

type RespuestaInformacionAsignacionInicial struct {
	InfoApropiacion map[string]float64
	Aprobado        bool
}

type RespuestaAprobacionAsignacionInicial map[string]interface{}

type PorDefinir struct {
}

// --------------------------------------------------------------------------------------------
// CdpController

type PlanCuentasMongoCrudDocumentoPresupuestal struct {
	plan_cuentas_mongo_crud.DocumentoPresupuestal
	ID            string       `json:"_id" bson:"_id,omitempty"`
	Data          interface{}  `json:"Data" bson:"data" validate:"required"`
	Tipo          string       `json:"Tipo" bson:"tipo" validate:"required"`
	AfectacionIds []string     `json:"AfectacionIds" bson:"afectacion_ids"`
	Afectacion    []Movimiento `bson:"-" validate:"required"`
	FechaRegistro string       `json:"FechaRegistro" bson:"fecha_registro" validate:"required"`
	Estado        string       `json:"Estado" bson:"estado"`
	ValorActual   float64      `json:"ValorActual" bson:"valor_actual"`
	ValorInicial  float64      `json:"ValorInicial" bson:"valor_inicial"`
	Vigencia      int          `json:"Vigencia" bson:"vigencia" validate:"required"`
	CentroGestor  string       `json:"CentroGestor" bson:"centro_gestor" validate:"required"`
	Consecutivo   int          `json:"Consecutivo" bson:"consecutivo"`
}

// TODO: Cuando Beego soporte lo siguiente al generar el Swagger, cambiar lo anterior por
// type PlanCuentasMongoCrudDocumentoPresupuestal struct {
// 	plan_cuentas_mongo_crud.DocumentoPresupuestal
// }
// o simplemente por
// type PlanCuentasMongoCrudDocumentoPresupuestal plan_cuentas_mongo_crud.DocumentoPresupuestal

type SolicitudAprobacionCdp struct {
	Id            string `json:"_id"`
	Vigencia      string `json:"vigencia"`
	AreaFuncional string `json:"area_funcional"`
}

// --------------------------------------------------------------------------------------------
// CrpController

type RespuestaGetFullCrp struct {
	SolicitudCrp          interface{} `json:"solicitudCrp"`
	ConsecutivoCdp        interface{} `json:"consecutivoCdp"`
	Vigencia              interface{} `json:"vigencia"`
	MovimientoCdp         interface{} `json:"movimiento_cdp"`
	CentroGestor          interface{} `json:"centroGestor"`
	Estado                interface{} `json:"estado"`
	NecesidadFinanciacion interface{} `json:"necesidadFinanciacion"`
}
