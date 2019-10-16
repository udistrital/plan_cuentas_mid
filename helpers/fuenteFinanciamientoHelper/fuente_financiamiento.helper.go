package fuenteFinanciamientoHelper

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// URLCRUD Ruta de plan_cuentas_crud
var URLCRUD = beego.AppConfig.String("planCuentasApiService") + "fuente_financiamiento"

// URLMONGO Ruta de plan_cuentas_mongo_crud
var URLMONGO = beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento"

// RegistrarFuenteHelper ...
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

// RegistrarFuenteMongo registra una fuente de financiamiento en mongo
func RegistrarFuenteMongo(fuenteFinancimiento *models.FuenteFinanciamiento) {
	// if err := request.SendJson(URLMONGO, "POST", &res, &mongoData); err != nil {
	// 	log.Println("mongoData: ", mongoData["Id"])
	// 	fuenteFinancimiento.Id = mongoData["Id"].(int)
	// 	eliminarFuente(*fuenteFinancimiento)
	// }
}

func eliminarFuente(fuenteFinanciamiento models.FuenteFinanciamiento) {
	if err := request.SendJson(URLCRUD, "DELETE", nil, &fuenteFinanciamiento); err == nil {
		log.Panicln(err.Error())
	}
}
