package managers

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// VigenciaManager ...
type VigenciaManager struct{}

// GetCurrentActiveVigencia ...
func (m *VigenciaManager) GetCurrentActiveVigencia(areaFuncional string) (vigencia int, err error) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("catch", r)
			errStr := fmt.Sprintf("%s", r)
			err = errors.New("Mongo API Error: " + errStr)
		}
	}()
	var response responseformat.Response
	// logs.Debug("url", beego.AppConfig.String("financieraMongoCurdApiService")+"vigencia/vigencia_actual_area/"+areaFuncional)
	if err = request.GetJson(beego.AppConfig.String("financieraMongoCurdApiService")+"vigencia/vigencia_actual_area/"+areaFuncional, &response); err == nil {
		if responseformat.CheckResponseError(response) {
			var errMessage = "Mongo API Error"
			if messageStr, e := response.Body.(string); e {
				errMessage = errMessage + ": " + messageStr
			} else {
				errMessage = errMessage + ": " + fmt.Sprintf("%s", response.Body)
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		} else {
			if err := formatdata.JsonPrint(response); err == nil {
				vigenciaMap := response.Body.([]interface{})[0].(map[string]interface{})
				vigencia = int(vigenciaMap["valor"].(float64))
			}
		}

	}

	return
}
