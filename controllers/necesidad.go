package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	necesidadhelper "github.com/udistrital/plan_cuentas_mid/helpers/necesidadHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/responseformat"
)

// NecesidadController operations for Necesidad
type NecesidadController struct {
	beego.Controller
}

// URLMapping ...
func (c *NecesidadController) URLMapping() {
	c.Mapping("GetFullNecesidad", c.GetFullNecesidad)
}

// GetFullNecesidad ...
// @Title GetFullNecesidad
// @Description retorna full Necesidad
// @Param	id		path 	string	true		"The key for necesidad"
// @Success 201 {object} models.TrNecesidad
// @Failure 403 body is empty
// @router /getfullnecesidad/:id [get]
func (c *NecesidadController) GetFullNecesidad() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	id := c.Ctx.Input.Param(":id")
	if response, err := necesidadhelper.GetTrNecesidad(id); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// PostFullNecesidad ...
// @Title PostFullNecesidad
// @Description create TrNecesidad
// @Param	body		body 	models.TrNecesidad	true "body for TrNecesidad content"
// @Success 201 {object} models.TrNecesidad
// @Failure 403 body is empty
// @router /postfullnecesidad [post]
func (c *NecesidadController) PostFullNecesidad() {
	var (
		v models.TrNecesidad
	)
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		beego.Info("bien")
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
	if response, err := necesidadhelper.PostTrNecesidad(v); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 201)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}
