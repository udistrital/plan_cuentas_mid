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
// @Success 201 {object} models.Apropiacion
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
		v.EstadoApropiacionId.Id = 1
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
// @Param id       path  string             true  "vigencia a comprobar"
// @Param valor    path  string             true  "unidad ejecutora"
// @Param vigencia path  string             true  "vigencia a comprobar"
// @Param body     body  models.Apropiacion true  "body for Apropiacion content"
// @Success 201 {object} models.Apropiacion
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
}

// ArbolApropiaciones ...
// @Title ArbolApropiaciones
// @Description Get Arbol Rubros By UE
// @Param unidadEjecutora path  int64  true  "unidad ejecutora a consultar"
// @Param vigencia        path  string true  "vigencia a comprobar"
// @Param rama            query string false "rama a consultar"
// @Success 200 {object}  []models.Rubro
// @Failure 403
// @router /ArbolApropiaciones/:unidadEjecutora/:vigencia [get]
func (c *ApropiacionController) ArbolApropiaciones() {

	var response []map[string]interface{}
	var urlmongo string
	ueStr := c.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	rama := c.GetString("rama")
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

// ArbolRubroApropiaciones ...
// @Title ArbolRubroApropiaciones
// @Description Get Arbol Rubros apropiacion para usar en el cliente presupuesto
// @Param	unidadEjecutora		path 	int64	true		"unidad ejecutora a consultar"
// @Param	vigencia		path 	int64	true		"vigencia a consultar"
// @Param	raiz		path 	int64	true		"raiz a consultar"
// @Param	nivel		query 	string	false		"nivel a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /ArbolRubroApropiacion/:unidadEjecutora/:vigencia/:raiz [get]
func (c *ApropiacionController) ArbolRubroApropiaciones() {

	ueStr := c.Ctx.Input.Param(":unidadEjecutora")
	vigenciaStr := c.Ctx.Input.Param(":vigencia")
	raiz := c.Ctx.Input.Param(":raiz")
	nivel := c.GetString("nivel")
	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0458", 500)
		}
	}()

	arbol, err := apropiacionHelper.ConstruirArbolRubroApropiacion(ueStr, vigenciaStr, raiz, nivel)
	if err != nil {
		panic("Mongo API Service Error")
	}
	c.Data["json"] = arbol
	c.ServeJSON()
}

// SaldoApropiacion ...
// @Title SaldoApropiacion
// @Description Get Arbol Rubros By UE
// @Param rubro           path  int64  true  "rubro a consultar"
// @Param unidadEjecutora path  int64  true  "unidad ejecutora a consultar"
// @Param vigencia        path  int64  true  "vigencia a consultar"
// @Success 200 {object} models.Rubro
// @Failure 403
// @router /SaldoApropiacion/:rubro/:unidadEjecutora/:vigencia [get]
func (c *ApropiacionController) SaldoApropiacion() {
	var (
		rubroParam    string
		unidadEParam  int
		vigenciaParam int
		err           error
		res           map[string]float64
	)

	defer func() {

		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}

	}()
	res = make(map[string]float64)
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
