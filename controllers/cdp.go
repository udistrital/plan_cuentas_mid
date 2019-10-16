package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	cdphelper "github.com/udistrital/plan_cuentas_mid/helpers/cdpHelper"
	"github.com/udistrital/utils_oas/responseformat"
)

// CdpController operations for Cdp
type CdpController struct {
	beego.Controller
}

// ExpedirCdp ...
// @Title ExpedirCdp
// @Description expedir cdp creando objeto infocdp en la solicitud
// @Param	id		path 	string	true		"The key for solicitudcdp"
// @Success 201 {object} models.SolicitudCDP
// @router /expedirCdp/:id [get]
func (c *CdpController) ExpedirCdp() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	id := c.Ctx.Input.Param(":id")
	response := make(map[string]interface{})
	response, _ = cdphelper.ExpedirCdp(id)
	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}

// SolicitarCdp ...
// @Title SolicitarCDP
// @Description create SolicitudCDP
// @Param	body		body 	models.SolicitudCDP true "body for Solicitud content"
// @Success 201 {int} models.SolicitudCDP
// @Failure 403 body is empty
// @router /solicitarCDP [post]
func (c *CdpController) SolicitarCdp() {
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
	if response, err := cdphelper.SolicitarCDP(v); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}
