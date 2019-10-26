package controllers

import (
	"encoding/json"

	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/compositor"
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
}

// Post ...
// @Title Create
// @Description create Modificacion Presupuestal
// @Param	body		body 	models.Movimiento	true		"body for Movimiento content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router / [post]
func (c *ModificacionPresupuestalController) Post() {
	var (
		modificacionPresupuestalData models.ModificacionPresupuestalReceiver
		finalData                    interface{}
	)

	defer func() {
		c.Data["json"] = finalData
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &modificacionPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}

	documentoPresupuestalDataFormated := modificacionpresupuestalhelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)

	finalData, err := compositor.AddMovimientoTransaction(modificacionPresupuestalData.Data, documentoPresupuestalDataFormated, documentoPresupuestalDataFormated.AfectacionMovimiento)
	if err != nil {
		logs.Debug("error", err)
		panic(err.Error())
	}

}

// SimulacionAfectacion ...
// @Title Create
// @Description create Modificacion Presupuestal
// @Param	body		body 	models.ModificacionPresupuestalReceiverAfectation	true		"body for simulacion_afectacion_modificacion content"
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
		c.Data["json"] = finalData
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &modificacionPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}
	if modificacionPresupuestalData.Afectation != nil {
		modificacionPresupuestalData.Data = &models.ModificacionPresupuestalReceiverDetail{}
		documentoPresupuestalDataFormated := modificacionpresupuestalhelper.ConvertModificacionToDocumentoPresupuestal(modificacionPresupuestalData)
		afectation = movimientohelper.FormatDataForMovimientosMongoAPI(documentoPresupuestalDataFormated.AfectacionMovimiento...)
	}
	response, err := movimientomanager.SimualteAfectationAPIMongo(cgStr, vigenciaStr, afectation...)
	if err != nil {
		panic(err.Error())
	}
	finalData = response.Body

}