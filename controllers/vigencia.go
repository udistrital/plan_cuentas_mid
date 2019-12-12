package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	vigenciahelper "github.com/udistrital/plan_cuentas_mid/helpers/vigenciaHelper"
)

// VigenciaController operations for vigencia
type VigenciaController struct {
	beego.Controller
}

// ...
// @Title GetCierreVigencia
// @Description devuelve los objetos del cierre para una vigencia y un area funcional
// @Param	vigencia		path 	string	true		"vigencia del cierre"
// @Param	area			path 	string	true		"area funcional del cierre"
// @Success 201 {object} models.SolicitudCDP
// @router /get_cierre/:vigencia/:area [get]
func (c *VigenciaController) GetCierreVigencia() {
	vigencia := c.Ctx.Input.Param(":vigencia")
	areaf := c.Ctx.Input.Param(":area")
	v, err := vigenciahelper.GetCierre(vigencia, areaf)
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
