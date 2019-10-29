package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	necesidadhelper "github.com/udistrital/plan_cuentas_mid/helpers/necesidadHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// NecesidadController operations for Necesidad
type NecesidadController struct {
	beego.Controller
}

// URLMapping ...
func (c *NecesidadController) URLMapping() {
	c.Mapping("GetFullNecesidad", c.GetFullNecesidad)
	c.Mapping("PostFullNecesidad", c.PostFullNecesidad)
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
// @router /post_full_necesidad [post]
func (c *NecesidadController) PostFullNecesidad() {
	var (
		v         models.TrNecesidad
		necesidad map[string]interface{}
	)
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		idNecesidad := strconv.FormatFloat((*v.Necesidad)["Id"].(float64), 'f', 0, 64)
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "/necesidad/" + idNecesidad

		if err := request.GetJson(urlcrud, &necesidad); err == nil {

			if necesidad["Id"] == nil { // La necesidad NO EXISTE

				if response, err := necesidadhelper.PostTrNecesidad(v); err != nil {
					responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
				} else {
					responseformat.SetResponseFormat(&c.Controller, response, "", 201)
				}

			} else { // La necesidad EXISTE

				if err := request.SendJson(urlcrud, "DELETE", nil, nil); err == nil {

					if response, err := necesidadhelper.PostTrNecesidad(v); err != nil {
						responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
					} else {
						responseformat.SetResponseFormat(&c.Controller, response, "", 201)
					}
				}
			}

		}
	} else {
		responseformat.SetResponseFormat(&c.Controller, err, "E_0458", 500)
	}
}
