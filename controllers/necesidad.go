package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	necesidad_models "github.com/udistrital/necesidades_crud/models"
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
	c.Mapping("Put", c.Put)
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
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, ok := (*v.Necesidad)["Id"].(float64); ok {
			idNecesidad := strconv.FormatFloat((*v.Necesidad)["Id"].(float64), 'f', 0, 64)
			urlcrud := beego.AppConfig.String("necesidadesCrudService") + "/necesidad/" + idNecesidad

			if err := request.GetJson(urlcrud, &necesidad); err == nil {

				if necesidad["Id"] == nil { // La necesidad NO EXISTE

					if response, err := necesidadhelper.PostTrNecesidad(v); err != nil {
						c.Abort("400")
					} else {
						c.Data["json"] = response
					}

				} else { // La necesidad EXISTE
					var resM map[string]interface{}
					if err := request.SendJson(urlcrud, "DELETE", &resM, nil); err == nil {
						if response, err := necesidadhelper.PostTrNecesidad(v); err != nil {
							c.Abort("400")
						} else {
							c.Data["json"] = response
						}
					} else {
						c.Abort("400")
					}
				}
			}

		} else {
			if response, err := necesidadhelper.PostTrNecesidad(v); err != nil {
				c.Abort("400")
			} else {
				c.Data["json"] = response
			}
		}
	} else {
		c.Abort("400")
	}

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Necesidad
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	necesidad_models.Necesidad	true		"body for Necesidad content"
// @Success 200 {object} necesidad_models.Necesidad
// @Failure 400 the request contains incorrect syntax
// @router /:id [put]
func (c *NecesidadController) Put() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error(err)
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "NecesidadController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("500") // Error no manejado!
			}
		}
	}()

	var (
		id        int
		necesidad necesidad_models.Necesidad
	)

	if v, err := c.GetInt(":id"); err != nil {
		panic(err)
	} else {
		id = v
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &necesidad); err != nil {
		panic(err)
	}

	if v, err := necesidadhelper.InterceptorMovimientoNecesidad(id, necesidad); err != nil {
		panic(err)
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}
