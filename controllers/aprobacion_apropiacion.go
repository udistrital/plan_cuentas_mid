package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/responseformat"

	"github.com/udistrital/plan_cuentas_mid/helpers"
	"github.com/udistrital/plan_cuentas_mid/helpers/apropiacionHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// AprobacionController operations for AprobacionController
type AprobacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *AprobacionController) URLMapping() {
	c.Mapping("InformacionAsignacionInicial", c.InformacionAsignacionInicial)
	c.Mapping("AprobacionAsignacionInicial", c.AprobacionAsignacionInicial)
	c.Mapping("Aprobado", c.Aprobado)
}

// InformacionAsignacionInicial ...
// @Title InformacionAsignacionInicial
// @Description Devuelve saldos iniciales antes de aprobar
// @Param	Vigencia		query 	string	true		"vigencia a comprobar"
// @Param	UnidadEjecutora		query 	string	true		"unidad ejecutora de los rubros a comprobar"
// @Success 200 {object} models.RespuestaInformacionAsignacionInicial
// @Failure 403
// @router /InformacionAsignacionInicial/ [get]
func (c *AprobacionController) InformacionAsignacionInicial() {

	asignationInfo := map[string]float64{"2": 0.0, "3": 0.0}

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}
	}()
	vigencia, err := c.GetInt("Vigencia")
	if err != nil {
		panic(helpers.InternalErrorMessage())
	}
	unidadejecutora, err := c.GetInt("UnidadEjecutora")
	if err != nil {
		panic(helpers.InternalErrorMessage())
	}

	compareFlag := apropiacionHelper.CompareApropiationNodes(&asignationInfo, unidadejecutora, vigencia)

	beego.Debug(compareFlag, asignationInfo)
	response := models.RespuestaInformacionAsignacionInicial{
		InfoApropiacion: asignationInfo,
		Aprobado:        compareFlag,
	}

	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}

// AprobacionAsignacionInicial ...
// @Title AprobacionAsignacionInicial
// @Description aprueba la asignacion inicial de presupuesto
// @Param	Vigencia		query 	string	true		"vigencia a comprobar"
// @Param	UnidadEjecutora		query 	string	true		"unidad ejecutora de los rubros a comprobar"
// @Success 200 {object} models.PorDefinir
// @Failure 403
// @router /AprobacionAsignacionInicial/ [post]
func (c *AprobacionController) AprobacionAsignacionInicial() {

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}
	}()

	vigencia, err := c.GetInt("Vigencia")

	if err != nil {
		panic(helpers.InternalErrorMessage())
	}

	unidadejecutora, err := c.GetInt("UnidadEjecutora")

	if err != nil {
		panic(helpers.InternalErrorMessage())
	}

	response := apropiacionHelper.AprobarPresupuesto(vigencia, unidadejecutora)
	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}

// Aprobado ...
// @Title Aprobado
// @Description aprueba la asignacion inicial de presupuesto
// @Param	Vigencia		query 	string	true		"vigencia a comprobar"
// @Param	UnidadEjecutora		query 	string	true		"unidad ejecutora de los rubros a comprobar"
// @Success 200 {object} bool
// @Failure 403
// @router /Aprobado [get]
func (c *AprobacionController) Aprobado() {

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			responseformat.SetResponseFormat(&c.Controller, r, "E_0459", 500)
		}
	}()

	vigencia, err := c.GetInt("Vigencia")

	if err != nil {
		panic(helpers.InternalErrorMessage())
	}

	unidadejecutora, err := c.GetInt("UnidadEjecutora")

	if err != nil {
		panic(helpers.InternalErrorMessage())
	}

	response := apropiacionHelper.PresupuestoAprobado(vigencia, unidadejecutora)
	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}
