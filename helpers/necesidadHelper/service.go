package necesidadhelper

import (
	"strconv"

	"github.com/astaxie/beego"
	necesidad_models "github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/utils_oas/request"
)

func PutNecesidadService(id int, necesidadent necesidad_models.Necesidad) (necesidad necesidad_models.Necesidad, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "PutNecesidadService - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	urlputnecesidadcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad/" + strconv.Itoa(id)
	if err := request.SendJson(urlputnecesidadcrud, "PUT", &necesidad, necesidadent); err != nil {
		return necesidad, map[string]interface{}{
			"funcion": "PutNecesidadService - request.SendJson(urlputnecesidadcrud, \"PUT\", &necesidad, necesidadent)",
			"err":     err,
			"status":  "500",
		}
	}
	return necesidad, nil
}
