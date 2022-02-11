package consecutivohelper

import (
	"github.com/astaxie/beego"
	models_configuracion "github.com/udistrital/configuracion_api/models"
	models_consecutivos "github.com/udistrital/consecutivos_crud/models"
	"github.com/udistrital/utils_oas/request"
)

func ObtenerProcesoNecesidad() (configuracion models_configuracion.Proceso, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ObtenerProcesoNecesidad - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	urlproceso := beego.AppConfig.String("configuracionCrudService")
	if err := request.GetJson(urlproceso+"proceso?query=Sigla__contains%3Anc", &configuracion); err != nil {
		return configuracion, map[string]interface{}{
			"funcion": "ObtenerProcesoNecesidad - request.GetJson(urlproceso+\"proceso?query=Sigla__contains%3Anc\", &resp)",
			"err":     err,
			"status":  "500",
		}
	}
	return configuracion, nil
}

//CrearMovimiento, Crea un resgistro del movimiento realizado a partir de un movimiento proceso externo y los datos de la necesidad
func GenerarConsecutivo(modeloconsecutivo models_consecutivos.Consecutivo) (consecutivo models_consecutivos.Consecutivo, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "GenerarConsecutivo - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	urlgenerarconsecutivo := beego.AppConfig.String("movimientosCrudService") + "movimiento_detalle/crearMovimientosDetalle/"
	if err := request.SendJson(urlgenerarconsecutivo, "POST", &consecutivo, modeloconsecutivo); err != nil {
		return consecutivo, map[string]interface{}{
			"funcion": "GenerarConsecutivo - request.SendJson(urlgenerarconsecutivo, \"POST\", &consecutivo, modeloconsecutivo)",
			"err":     err,
			"status":  "500",
		}
	}

	return consecutivo, nil
}
