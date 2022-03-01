package consecutivo

import (
	"github.com/astaxie/beego"
	models_consecutivos "github.com/udistrital/consecutivos_crud/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

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
	var temporal map[string]interface{}
	urlgenerarconsecutivo := beego.AppConfig.String("consecutivoApiService") + "consecutivo"
	if err := request.SendJson(urlgenerarconsecutivo, "POST", &temporal, modeloconsecutivo); err != nil {
		return consecutivo, map[string]interface{}{
			"funcion": "GenerarConsecutivo - request.SendJson(urlgenerarconsecutivo, \"POST\", &consecutivo, modeloconsecutivo)",
			"err":     err,
			"status":  "502",
		}
	}
	formatdata.FillStruct(temporal["Data"], &consecutivo)
	return consecutivo, nil
}
