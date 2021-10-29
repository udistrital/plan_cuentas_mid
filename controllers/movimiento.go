package controllers

import (
	"encoding/json"

	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/compositor"
	documentopresupuestalmanager "github.com/udistrital/plan_cuentas_mid/managers/documentoPresupuestalManager"

	commonhelper "github.com/udistrital/plan_cuentas_mid/helpers/commonHelper"
	modificacionpresupuestalhelper "github.com/udistrital/plan_cuentas_mid/helpers/modificacionPresupuestalHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// MovimientoController operations for Movimiento
type MovimientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovimientoController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Movimiento
// @Param	body		body 	models.DocumentoPresupuestal	true		"body for Movimiento content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoController) Post() {
	var (
		documentoPresupuestalData models.DocumentoPresupuestal
		finalData                 interface{}
	)

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "", 500)
		}
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &documentoPresupuestalData); err != nil {
		logs.Error(err.Error())
		panic(err.Error())
	}

	if errStrc := formatdata.StructValidation(documentoPresupuestalData); len(errStrc) > 0 {
		responseformat.SetResponseFormat(&c.Controller, errStrc, "", 422)
		return
	}

	for _, elmnt := range documentoPresupuestalData.Afectacion {
		if errStrc := formatdata.StructValidation(elmnt); len(errStrc) > 0 {
			responseformat.SetResponseFormat(&c.Controller, errStrc, "", 422)
			return
		}
	}

	// Send Data to Movimientos API to Add the current movimiento data to postgres.
	finalData, err := compositor.AddMovimientoTransaction(documentoPresupuestalData.Data, documentoPresupuestalData, documentoPresupuestalData.AfectacionMovimiento)
	if err != nil {
		logs.Debug("error", err)
		panic(err.Error())
	}

	responseformat.SetResponseFormat(&c.Controller, finalData, "", 200)

}

// GetAllAnulacionesByVigenciaCGAndUUID función para obtener todos los objetos
// @Title GetAllAnulacionesByVigenciaCGAndUUID
// @Description get all objects
// @Param vigencia path  uint   true  "vigencia"
// @Param CG       path  string true  "centro gestor / unidad ejecutora"
// @Param UUID     path  string true  "Identificador"
// @Success 200 {object} []models.AnulationDetail
// @Failure 403 :objectId is empty
// @router /get_doc_by_mov_parentUUID/:vigencia/:CG/:UUID [get]
func (c *MovimientoController) GetAllAnulacionesByVigenciaCGAndUUID() {
	vigencia := c.GetString(":vigencia")
	centroGestor := c.GetString(":CG")
	documentUUID := c.GetString(":UUID")
	var response []models.AnulationDetail
	rows, err := documentopresupuestalmanager.GetAllPresupuestalDocumentFromCRUDByMovParentUUID(vigencia, centroGestor, documentUUID)
	if err == nil {
		response = modificacionpresupuestalhelper.FormatDocumentoPresupuestalResponseToAnulationDetail(rows)
	}
	c.Data["json"] = commonhelper.DefaultResponse(200, err, response)

	c.ServeJSON()
}
