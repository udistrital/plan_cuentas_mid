package controllers 

import (
	// "github.com/udistrital/plan_cuentas_mid/helpers"
	"encoding/json"
	fuenteHelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteFinanciamientoHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/responseformat"
)


// AprobacionController operations for AprobacionController
type FuenteFinanciamientoController struct {
	beego.Controller
}

// URLMapping ...
func (c *FuenteFinanciamientoController) URLMapping() {
	c.Mapping("RegistrarFuenteFinanciamiento", c.RegistrarFuenteFinanciamiento)
}

// RegistrarFuenteFinanciamiento ...
// @Title RegistrarFuenteFinanciamiento
// @Description Registra la fuente de financiamiento en postgres y mongo
// @Param	FuenteFinanciamiento		query 	models.Fuentefinanciamiento	true		"models.Fuentefinanciamiento"
// @Success 200 {string} resultado
// @Failure 403
// @router RegistrarFuenteFinanciamiento/ [post]
func (c *FuenteFinanciamientoController) RegistrarFuenteFinanciamiento() {
	var v models.FuenteFinanciamiento

	response := make(map[string]interface{})
	//asignationInfo := map[string]float64{"2": 0.0, "3": 0.0}

	defer errorResponse(c)
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		fuenteHelper.RegistrarFuenteHelper(&v)
	} else {
		panic(err)
	}

	responseformat.SetResponseFormat(&c.Controller, response, "", 200)
}


func errorResponse(c *FuenteFinanciamientoController) {
	if r := recover(); r != nil {
		beego.Error(r)
		responseformat.SetResponseFormat(&c.Controller, r, "E", 500)
	}
}