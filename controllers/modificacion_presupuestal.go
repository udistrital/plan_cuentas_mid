package controllers

import (
	"encoding/json"
	"strconv"

	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mid/managers/documentoPresupuestalManager"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
	errorctrl "github.com/udistrital/utils_oas/errorctrl"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/compositor"
	commonhelper "github.com/udistrital/plan_cuentas_mid/helpers/commonHelper"
	modificacionpresupuestalhelper "github.com/udistrital/plan_cuentas_mid/helpers/modificacionPresupuestalHelper"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// ModificacionPresupuestalController operations for ModificacionPresupuestal
type ModificacionPresupuestalController struct {
	beego.Controller
}

// URLMapping ...
func (c *ModificacionPresupuestalController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("SimulacionAfectacion", c.SimulacionAfectacion)
	c.Mapping("GetAllModificacionPresupuestalByVigenciaAndCG", c.GetAllModificacionPresupuestalByVigenciaAndCG)
	c.Mapping("GetOneModificacionPresupuestalByVigenciaAndCG", c.GetOneModificacionPresupuestalByVigenciaAndCG)
}

// Post ...
// @Title Create
// @Description create Modificacion Presupuestal
// @Param	body		body 	models.ModificacionPresupuestalReceiver	true		"body for Movimiento content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router / [post]
func (c *ModificacionPresupuestalController) Post() {
	var (
		modificacionPresupuestalData models.ModificacionPresupuestalReceiver
		finalData                    map[string]interface{}
	)

	defer func() {
		if r := recover(); r != nil {
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
		cdpMessage := ""
		var cdpArr []int

		if finalData["Sequences"] != nil {
			for _, cdpNumber := range finalData["Sequences"].([]interface{}) {

				cdpArr = append(cdpArr, int(cdpNumber.(float64)))

			}
		}

		if len(cdpArr) > 0 {
			cdpMessage += "CDP Generados: "
			for _, n := range cdpArr {
				cdpMessage += strconv.Itoa(n)
				if n < len(cdpArr) {
					cdpMessage += ","
				}
			}
		}
		responseformat.SetResponseFormat(&c.Controller, "Modificación registrada correctamente. "+cdpMessage, "", 200)
		c.ServeJSON()
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &modificacionPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}

	documentoPresupuestalDataFormated := modificacionpresupuestalhelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)

	finalDataIntf, err := compositor.AddMovimientoTransaction(modificacionPresupuestalData.Data, documentoPresupuestalDataFormated, documentoPresupuestalDataFormated.AfectacionMovimiento)

	if err != nil {
		logs.Error("error", err)
		panic(err.Error())
	}

	finalData = finalDataIntf.(map[string]interface{})

}

// SimulacionAfectacion ...
// @Title Create
// @Description create Modificacion Presupuestal
// @Param centroGestor path  string                                  true  "centro gestor / unidad ejecutora"
// @Param vigencia     path  uint                                    true  "vigencia"
// @Param body         body  models.ModificacionPresupuestalReceiver true  "body for simulacion_afectacion_modificacion content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router /simulacion_afectacion_modificacion/:centroGestor/:vigencia [post]
func (c *ModificacionPresupuestalController) SimulacionAfectacion() {
	var (
		modificacionPresupuestalData models.ModificacionPresupuestalReceiver
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
		modificacionPresupuestalData.Data.CentroGestor = cgStr

		documentoPresupuestalDataFormated := modificacionpresupuestalhelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)
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

// GetAllModificacionPresupuestalByVigenciaAndCG función para obtener todos los objetos
// @Title GetAllModificacionPresupuestalByVigenciaAndCG
// @Description get all objects
// @Param vigencia path  uint   true  "vigencia"
// @Param CG       path  string true  "centro gestor / unidad ejecutora"
// @Param tipo     path  string true  "tipo de documento"
// @Success 200 DocumentoPresupuestal models.DocumentoPresupuestal
// @Failure 403 :objectId is empty
// @router /:vigencia/:CG/:tipo [get]
func (c *ModificacionPresupuestalController) GetAllModificacionPresupuestalByVigenciaAndCG() {
	defer errorctrl.ErrorControlController(c.Controller, "ModificacionPresupuestalController")
	vigencia := c.GetString(":vigencia")
	centroGestor := c.GetString(":CG")
	tipoModificacion := c.GetString(":tipo")
	var response []models.ModificacionPresupuestalResponseDetail
	rows, err := documentopresupuestalmanager.GetAllPresupuestalDocumentFromCRUDByType(vigencia, centroGestor, tipoModificacion)
	if err == nil {
		response = modificacionpresupuestalhelper.FormatDocumentoPresupuestalResponseToModificacionDetail(rows)
	}
	c.Data["json"] = commonhelper.DefaultResponse(200, err, response)

	c.ServeJSON()
}

// GetOneModificacionPresupuestalByVigenciaAndCG función para obtener todos los objetos
// @Title GetOneModificacionPresupuestalByVigenciaAndCG
// @Description get all objects
// @Param vigencia path  uint   true  "vigencia"
// @Param CG       path  string true  "centro gestor / unidad ejecutora"
// @Param UUID     path  string true  "Identificador"
// @Success 200 {object} models.ModificacionPresupuestalResponseDetail
// @Failure 403 :objectId is empty
// @router get_one/:vigencia/:CG/:UUID [get]
func (c *ModificacionPresupuestalController) GetOneModificacionPresupuestalByVigenciaAndCG() {
	defer errorctrl.ErrorControlController(c.Controller, "ModificacionPresupuestalController")
	vigencia := c.GetString(":vigencia")
	centroGestor := c.GetString(":CG")
	UUID := c.GetString(":UUID")
	var response models.ModificacionPresupuestalResponseDetail
	row, err := documentopresupuestalmanager.GetOnePresupuestalDocumentFromCRUDByID(vigencia, centroGestor, UUID)
	if err == nil {
		response = modificacionpresupuestalhelper.FormatDocumentoPresupuestalToModificacion(row)
	}
	c.Data["json"] = commonhelper.DefaultResponse(200, err, response)

	c.ServeJSON()
}
