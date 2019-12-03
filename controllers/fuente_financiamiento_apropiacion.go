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
	c.Mapping("GetRubrosbyFuente", c.GetRubrosbyFuente)
}

// GetRubrosbyFuente ...
// @Title GetRubrosbyFuente
// @Description retorna rubros de la fuente desde el plan de adquisición
// @Success 201 {object} models.Fuentefinanciamiento
// @Failure 403 :vigencia is empty
// @Failure 403 :id is empty
// @router /plan_adquisiciones_rubros_fuente/:vigencia/:id [get]
func (c *FuenteFinanciamientoApropiacionController) GetRubrosbyFuente() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
	}()
	vigencia := c.GetString(":vigencia")
	objectID := c.GetString(":id")
	if response, err := fuenteApropiacionHelper.GetPlanAdquisicionbyFuente(vigencia, objectID); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
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

		// Primero registra la fuente
		idFuente := fuenteHelper.RegistrarFuenteHelper(fuenteFinanciamiento)
		// Le asigna el id registrado
		fuenteFinanciamiento.Id = idFuente

		if err := formatdata.FillStruct(v["FuentesFinanciamientoApropiacion"], &fuentesFinanciamientoApropiacion); err != nil {
			log.Panicln(err.Error())
		}
		/*
		 Apartir del atributo FuentesFinanciamientoApropiacion del json enviado como parámetro de esta petición, concatena todos
		 los valores en el arreglo y les asigna los id correspondientes de ApropiacionId y FuenteFinanciamientoId
		*/
		fuentesContatenadas := fuenteApropiacionHelper.ConcatenarFuente(fuenteFinanciamiento, fuentesFinanciamientoApropiacion...)
		// Registra todos los fuentes_financiamiento_apropiacion
		idsFuentesRegistrados := fuenteApropiacionHelper.RegistrarMultipleFuenteApropiacion(fuentesContatenadas)
		// Formatea los datos para que puedan ser enviados  para registrar movimientos
		dataFormateada := fuenteApropiacionHelper.FormatDataMovimientoExterno(idsFuentesRegistrados, fuentesFinanciamientoApropiacion...)
		// Registra los movimientos
		fuenteApropiacionHelper.RegistrarMultipleMovimientoExterno(dataFormateada)

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
