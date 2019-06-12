package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/helpers/apropiacionHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// ApropiacionController operations for  Apropiacion
type ApropiacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApropiacionController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create Apropiacion
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 201 {int} models.Apropiacion
// @Failure 403 body is empty
// @router / [post]
func (c *ApropiacionController) Post() {
	var v models.Apropiacion
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.Estado.Id = 1
		response := apropiacionHelper.AddApropiacion(v)
		responseformat.SetResponseFormat(&c.Controller, response["Body"], response["Code"].(string), 200)
	} else {
		beego.Error(err.Error())
		responseformat.SetResponseFormat(&c.Controller, nil, "E_0458", 500)
	}
}

// Put ...
// @Title Put
// @Description Update Apropiacion
// @Param	body		body 	models.Apropiacion	true		"body for Apropiacion content"
// @Success 201 {int} models.Apropiacion
// @Failure 403 body is empty
// @router /:id/:valor/:vigencia [put]
func (c *ApropiacionController) Put() {
	var v map[string]interface{}
	idStr := c.Ctx.Input.Param(":id")
	valStr := c.Ctx.Input.Param(":valor")
	vigStr := c.Ctx.Input.Param(":vigencia")
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		response := apropiacionHelper.PutApropiacion(v, idStr, valStr, vigStr)
		responseformat.SetResponseFormat(&c.Controller, response["Body"], response["Code"].(string), 200)
	} else {
		beego.Error(err.Error())
		responseformat.SetResponseFormat(&c.Controller, nil, "E_0459", 500)
	}
	c.ServeJSON()
}

// ArbolApropiaciones ...
// @Title ArbolApropiaciones
// @Description Get Arbol Rubros By UE
// @Param	unidadEjecutora		path 	int64	true		"unidad ejecutora a consultar"
// @Param	rama		query 	string	false		"rama a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /ArbolApropiaciones/:unidadEjecutora/:vigencia [get]
func (c *ApropiacionController) ArbolApropiaciones() {

	var response []map[string]interface{}
	ueStr := c.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	rama := c.GetString("rama")
	urlmongo := ""
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()
	if rama == "" {
		urlmongo = beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiaciones/RaicesArbolApropiacion/" + ueStr + "/" + vigenciaStr
	} else {
		urlmongo = beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiaciones/ArbolApropiacion/" + rama + "/" + ueStr + "/" + vigenciaStr
	}
	if err := request.GetJson(urlmongo, &response); err != nil {
		panic("Mongo API Service Error")
	}
	c.Data["json"] = response
}

// SaldoApropiacion ...
// @Title SaldoApropiacion
// @Description Get Arbol Rubros By UE
// @Param	unidadEjecutora		path 	int64	true		"unidad ejecutora a consultar"
// @Param	rama		query 	string	false		"rama a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /SaldoApropiacion/:rubro/:unidadEjecutora/:vigencia [get]
func (c *ApropiacionController) SaldoApropiacion() {
	var (
		rubroParam    string
		unidadEParam  int
		vigenciaParam int
		err           error
	)

	defer func() {

		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}

	}()
	res := make(map[string]float64)
	rubroParam = c.GetString(":rubro")
	if unidadEParam, err = c.GetInt(":unidadEjecutora"); err != nil {
		panic(err.Error())
	}

	if vigenciaParam, err = c.GetInt(":vigencia"); err != nil {
		panic(err.Error())
	}
	res = apropiacionHelper.CalcularSaldoApropiacion(rubroParam, unidadEParam, vigenciaParam)
	responseformat.SetResponseFormat(&c.Controller, res, "", 200)

}
