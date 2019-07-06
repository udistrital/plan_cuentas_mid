package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	fuenteApropiacionHelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteApropiacionHelper"
	fuenteHelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteFinanciamientoHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/responseformat"
)

// FuenteFinanciamientoApropiacionController operations for FuenteFinanciamientoApropiacionController
type FuenteFinanciamientoApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoApropiacionController) URLMapping() {
	c.Mapping("RegistrarFuenteConApropiacion", c.RegistrarFuenteConApropiacion)
}

// RegistrarFuenteConApropiacion ...
// @Title RegistrarFuenteConApropiacion
// @Description Registra la fuente de financiamiento en postgres y mongo
// @Param	FuenteFinanciamiento		query 	models.Fuentefinanciamiento	true		"models.Fuentefinanciamiento"
// @Success 200 {string} resultado
// @Failure 403
// @router registrar_fuentes_con_apropiacion [post]
func (c *FuenteFinanciamientoApropiacionController) RegistrarFuenteConApropiacion() {
	var (
		v                                map[string]interface{}
		fuentesFinanciamientoApropiacion []interface{}
		fuenteFinanciamiento             *models.FuenteFinanciamiento
	)

	defer errorResponse(c.Controller)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {

		if err := formatdata.FillStruct(v["FuenteFinanciamiento"], &fuenteFinanciamiento); err != nil {
			log.Panicln(err.Error())
		}

		idFuente := fuenteHelper.RegistrarFuenteHelper(fuenteFinanciamiento)
		fuenteFinanciamiento.Id = idFuente

		if err := formatdata.FillStruct(v["FuentesFinanciamientoApropiacion"], &fuentesFinanciamientoApropiacion); err != nil {
			log.Panicln(err.Error())
		}

		fuentesContatenadas := fuenteApropiacionHelper.ConcatenarFuente(fuenteFinanciamiento, fuentesFinanciamientoApropiacion...)

		fuenteApropiacionHelper.RegistrarMultipleFuenteApropiacion(fuentesContatenadas)
		

		if fuentesContatenadas == nil {
			log.Panicln(err.Error())
		}

		response := make(map[string]interface{})
		response["Body"] = "success"
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)

	} else {
		log.Panicln(err.Error())
	}
}

func errorResponse(c beego.Controller) {
	if r := recover(); r != nil {
		beego.Error(r)
		responseformat.SetResponseFormat(&c, r, "E", 500)
	}
}
