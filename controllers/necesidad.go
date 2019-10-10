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
	c.Mapping("InitNecesidad", c.InitNecesidad)
}

// InitNecesidad ...
// @Title initnecesidad
// @Description retorna full Necesidad
// @Param	id		path 	string	true		"The key for necesidad"
// @Success 201 {object} models.Necesidad
// @Failure 403 body is empty
// @router /initnecesidad/:id [get]
func (c *NecesidadController) InitNecesidad() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	id := c.Ctx.Input.Param(":id")
	response := make(map[string]interface{})
	response, _ = necesidadhelper.InitNecesidad(id)
	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}
