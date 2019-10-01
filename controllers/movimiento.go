package controllers

import (
	"encoding/json"

	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/compositor"
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
// @Param	body		body 	models.Movimiento	true		"body for Movimiento content"
// @Success 201 {object} models.Movimiento
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoController) Post() {
	var (
		documentoPresupuestalData models.DocumentoPresupuestal
		finalData                 interface{}
	)

	defer func() {
		c.Data["json"] = finalData
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
	finalData, err := compositor.AddMovimientoTransaction(documentoPresupuestalData)
	if err != nil {
		logs.Debug("error", err)
		panic(err.Error())
	}

}
