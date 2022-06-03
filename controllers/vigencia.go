package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	vigenciahelper "github.com/udistrital/plan_cuentas_mid/helpers/vigenciaHelper"
	errorctrl "github.com/udistrital/utils_oas/errorctrl"
)

// VigenciaController operations for vigencia
type VigenciaController struct {
	beego.Controller
}

// URLMapping ...
func (c *VigenciaController) URLMapping() {
	c.Mapping("GetCierreVigencia", c.GetCierreVigencia)
	c.Mapping("CerrarVigencia", c.CerrarVigencia)
}

// GetCierreVigencia s...
// @Title GetCierreVigencia
// @Description devuelve los objetos del cierre para una vigencia y un area funcional
// @Param	vigencia		path 	string	true		"vigencia del cierre"
// @Param	area			path 	string	true		"area funcional del cierre"
// @Param	cerrada  path  uint  true "'1' para NO cerrada"
// @Success 201 {object} models.SolicitudCDP
// @router /get_cierre/:vigencia/:area/:cerrada [get]
func (c *VigenciaController) GetCierreVigencia() {
	defer errorctrl.ErrorControlController(c.Controller, "VigenciaController")
	vigencia := c.Ctx.Input.Param(":vigencia")
	areaf := c.Ctx.Input.Param(":area")
	cerrada := c.Ctx.Input.Param(":cerrada") != "1"
	v, err := vigenciahelper.GetCierre(vigencia, areaf, cerrada)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// CerrarVigencia ...
// @Title CerrarVigencia
// @Description realiza los procesos del cierre para una vigencia y un area funcional
// @Param	vigencia		path 	string	true		"vigencia del cierre"
// @Param	area			path 	string	true		"area funcional del cierre"
// @Success 201 {object} models.SolicitudCDP
// @router /cerrar_vigencia/:vigencia/:area [get]
func (c *VigenciaController) CerrarVigencia() {
	defer errorctrl.ErrorControlController(c.Controller, "VigenciaController")
	vigencia := c.Ctx.Input.Param(":vigencia")
	areaf := c.Ctx.Input.Param(":area")
	v, err := vigenciahelper.CerrarVigencia(vigencia, areaf)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
