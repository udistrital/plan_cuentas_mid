package controllers

import (
	"encoding/json"
	"fmt"

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
// @Description expedir cdp creando objeto infocdp en la solicitud - TODO: Esto semánticamente debería ser un POST!
// @Param	id		path 	string	true		"The key for solicitudcdp"
// @Success 201 {object} models.SolicitudCDP
// @router /expedirCDP/:id [get]
func (c *CdpController) ExpedirCdp() {
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	id := c.Ctx.Input.Param(":id")
	if response, err := cdphelper.ExpedirCdp(id); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// SolicitarCdp ...
// @Title SolicitarCDP
// @Description create SolicitudCDP
// @Param	body		body 	models.NecesidadesCrudNecesidadParcial1 true "body for Solicitud content"
// @Success 201 {object} models.SolicitudCDP
// @Failure 403 body is empty
// @router /solicitarCDP [post]
func (c *CdpController) SolicitarCdp() {
	var v map[string]interface{}

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}

	if response, err := cdphelper.SolicitarCDP(v); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}

// AprobarCdp ...
// @Title AprobarCdp
// @Description create SolicitudCDP
// @Param	body		body 	models.SolicitudAprobacionCdp true "body for Solicitud content"
// @Success 201 {object} models.PlanCuentasMongoCrudDocumentoPresupuestal
// @Failure 403 body is empty
// @router /aprobar_cdp [post]
func (c *CdpController) AprobarCdp() {
	var v map[string]interface{}

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()

	// logs.Debug("c.Ctx.Input.RequestBody: ", c.Ctx.Input.RequestBody)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err != nil {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}

	// logs.Debug("v: ", v)

	for i, value := range v {
		switch value.(type) {
		case float64:
			v[i] = fmt.Sprintf("%.0f", value)
			// logs.Debug("v[i]: ", v[i])
			// logs.Debug("reflect.TypeOf(v[i]): ", reflect.TypeOf(v[i]))
		default:
			// logs.Debug("reflect.TypeOf(value)", reflect.TypeOf(value))
		}
	}

	if response, err := cdphelper.AprobarCdp(v["_id"].(string), v["vigencia"].(string), v["area_funcional"].(string)); err == nil {
		responseformat.SetResponseFormat(&c.Controller, response, "", 200)
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}
