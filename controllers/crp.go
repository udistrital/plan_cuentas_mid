package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	crphelper "github.com/udistrital/plan_cuentas_mid/helpers/crpHelper"
	"github.com/udistrital/utils_oas/responseformat"
)

// CrpController operations for Crp
type CrpController struct {
	beego.Controller
}

// ExpedirCrp ...
// @Title ExpedirCrp
// @Description expedir crp creando objeto infocrp en la solicitud
// @Param	id		path 	string	true		"The key for solicitudcrp"
// @Success 201 {object} models.SolicitudCRP
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
// @Param	body		body 	models.SolicitudCRP true "body for Solicitud content"
// @Success 201 {int} models.SolicitudCRP
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
