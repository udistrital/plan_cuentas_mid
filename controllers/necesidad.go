package controllers

import (
	"github.com/astaxie/beego"
	necesidadhelper "github.com/udistrital/plan_cuentas_mid/helpers/necesidadHelper"
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
