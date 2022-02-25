package configuracion

import (
	"net/url"

	"github.com/astaxie/beego"
	models_configuracion "github.com/udistrital/configuracion_api/models"
	"github.com/udistrital/utils_oas/request"
)

func ObtenerProceso(sigla string) (configuracion []models_configuracion.Proceso, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ObtenerProceso - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	parametros := url.Values{}
	parametros.Add("query", "Sigla__contains:"+sigla)
	urlproceso := beego.AppConfig.String("configuracionCrudService") + "proceso?" + parametros.Encode()
	if err := request.GetJson(urlproceso, &configuracion); err != nil {
		outputError = map[string]interface{}{
			"funcion": "ObtenerProceso - request.GetJson(urlproceso, &configuracion)",
			"err":     err,
			"status":  "502",
		}
	}
	return
}

func CrearProceso(modeloproceso models_configuracion.Proceso) (configuracion models_configuracion.Proceso, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "CrearProceso - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	urlproceso := beego.AppConfig.String("configuracionCrudService") + "proceso"
	if err := request.SendJson(urlproceso, "POST", &configuracion, modeloproceso); err != nil {
		outputError = map[string]interface{}{
			"funcion": "CrearProceso - request.SendJson(urlproceso, \"POST\", &configuracion, modeloproceso)",
			"err":     err,
			"status":  "502",
		}
	}
	return
}
