package controllers

import (
	"encoding/json"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/compositor"
	fuenteApropiacionHelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteApropiacionHelper"
	fuenteHelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteFinanciamientoHelper"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	fuentemanager "github.com/udistrital/plan_cuentas_mid/managers/fuenteManager"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
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
	c.Mapping("RegistrarModificacion", c.RegistrarModificacion)
	c.Mapping("SimulacionAfectacion", c.SimulacionAfectacion)
	c.Mapping("Delete", c.Delete)
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

// RegistrarModificacion ...
// @Title RegistrarModificacion
// @Description create Modificacion Presupuestal Fuente
// @Param	body		body 	models.Movimiento	true		"body for Movimiento content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router /modificacion [post]
func (c *FuenteFinanciamientoApropiacionController) RegistrarModificacion() {
	var (
		modificacionPresupuestalData models.ModificacionFuenteReceiver
		// finalData                    map[string]interface{}
	)
	defer func() {
		if r := recover(); r != nil {
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
		responseformat.SetResponseFormat(&c.Controller, "Modificación en la fuente registrada correctamente", "", 200)
		c.ServeJSON()
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &modificacionPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}
	documentoPresupuestalDataFormated := fuenteApropiacionHelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)
	// formatdata.JsonPrint(documentoPresupuestalDataFormated)
	_, err := compositor.AddMovimientoTransaction(modificacionPresupuestalData.Data, documentoPresupuestalDataFormated, documentoPresupuestalDataFormated.AfectacionMovimiento)

	if err != nil {
		logs.Debug("error", err)
		panic(err.Error())
	}

	//finalData = documentoPresupuestalDataFormated
	// finalData = finalDataIntf.(map[string]interface{})
	//fmt.Println(finalData)

}

// Delete ...
// @Title Borrar FuenteFinanciamiento
// @Description Borrar FuenteFinanciamiento
// @Param	id		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:id/:vigencia/:unidadEjecutora [delete]
func (c *FuenteFinanciamientoApropiacionController) Delete() {
	objectID := c.Ctx.Input.Param(":id")
	vigencia := c.Ctx.Input.Param(":vigencia")
	unidadEjecutora := c.Ctx.Input.Param(":unidadEjecutora")
	defer func() {
		if r := recover(); r != nil {
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
		responseformat.SetResponseFormat(&c.Controller, "delete success!", "", 200)
		c.ServeJSON()
	}()

	response, _ := fuenteApropiacionHelper.GetPlanAdquisicionbyFuente(vigencia, objectID)
	if response != nil {
		responseformat.SetResponseFormat(&c.Controller, "La fuente esta distribuida", "", 403)
	} else {
		_, err := fuentemanager.DeleteFuenteFinanciamiento(objectID, unidadEjecutora, vigencia)
		if err == nil {
			responseformat.SetResponseFormat(&c.Controller, "delete success!", "", 200)
		}
	}

}

// SimulacionAfectacion ...
// @Title Create
// @Description create Modificacion Presupuestal
// @Param	body		body 	models.ModificacionFuenteReceiver	true		"body for simulacion_afectacion_modificacion content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router /simulacion_afectacion_modificacion/:centroGestor/:vigencia [post]
func (c *FuenteFinanciamientoApropiacionController) SimulacionAfectacion() {
	var (
		modificacionPresupuestalData models.ModificacionFuenteReceiver
		finalData                    interface{}
	)
	cgStr := c.Ctx.Input.Param(":centroGestor")
	vigenciaStr := c.GetString(":vigencia")
	var afectation []models.MovimientoMongo
	defer func() {
		if r := recover(); r != nil {
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
		c.ServeJSON()
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &modificacionPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}
	if modificacionPresupuestalData.Afectation != nil {
		modificacionPresupuestalData.Data = &models.ModificacionPresupuestalReceiverDetail{}
		documentoPresupuestalDataFormated := fuenteApropiacionHelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)
		afectation = movimientohelper.FormatDataForMovimientosMongoAPI(documentoPresupuestalDataFormated.AfectacionMovimiento...)
	}
	response, err := movimientomanager.SimualteAfectationAPIMongo(cgStr, vigenciaStr, afectation...)
	if err != nil {
		panic(err)
	}
	if responseType, e := response["Type"].(string); e {
		if responseType == "error" {
			panic(response["Body"])
		}
	}
	finalData = response["Body"]
	c.Data["json"] = finalData
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
