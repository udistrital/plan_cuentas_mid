package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	crphelper "github.com/udistrital/plan_cuentas_mid/helpers/crpHelper"
	"github.com/udistrital/utils_oas/responseformat"
)

// CrpController operations for Crp
type CrpController struct {
	beego.Controller
}

// ExpedirCrp ...
// @Title ExpedirCrp
// @Description expedir crp creando objeto infocrp en la solicitud - TODO: Este método semánticamente debería ser un POST!
// @Param	id		path 	string	true		"The key for solicitudcrp"
// @Success 201 {object} models.SolicitudCrp
// @router /expedirCRP/:id [get]
func (c *CrpController) ExpedirCrp() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	id := c.Ctx.Input.Param(":id")
	if response, err := crphelper.ExpedirCrp(id); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// SolicitarCrp ...
// @Title SolicitarCRP
// @Description create SolicitudCRP
// @Param	body		body 	models.SolicitudCrp true "body for Solicitud content"
// @Success 201 {object} models.SolicitudCrp
// @Failure 403 body is empty
// @router /solicitarCRP [post]
func (c *CrpController) SolicitarCrp() {
	var (
		v map[string]interface{}
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
	if response, err := crphelper.SolicitarCRP(v); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// GetInfoCrp ...
// @Title Get Info CRPs
// @Description get all the information about CRPs
// @Success 200 {object} models.RespuestaGetFullCrp
// @Failure 404 not found resource
// @router /getFullCrp [get]
func (c *CrpController) GetInfoCrp() {
	v, err := crphelper.GetFullCrp()
	if err != nil {
		logs.Error(err)
		c.Data["system"] = err
		c.Abort("404")
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
