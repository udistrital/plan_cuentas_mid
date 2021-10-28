package fuenteFinanciamientoHelper

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// URLCRUD Ruta de plan_cuentas_crud
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
var URLCRUD = beego.AppConfig.String("planCuentasApiService") + "fuente_financiamiento"

// URLMONGO Ruta de plan_cuentas_mongo_crud
var URLMONGO = beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento"

// RegistrarFuenteHelper ...
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func RegistrarFuenteHelper(fuenteFinancimiento *models.FuenteFinanciamiento) (idFuente int) {
	var (
		res              map[string]interface{}
		fuenteRegistrada models.FuenteFinanciamiento
	)
	data := fuenteFinancimiento

	if err := request.SendJson(URLCRUD, "POST", &res, &data); err != nil {
		log.Panicln(err.Error())
		return
	}

	if err := formatdata.FillStruct(res["Body"], &fuenteRegistrada); err != nil {
		log.Panicln(err.Error())
		return
	}

	idFuente = fuenteRegistrada.Id
	return
}
