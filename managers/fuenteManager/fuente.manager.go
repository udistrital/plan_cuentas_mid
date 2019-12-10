package fuentemanager

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// DeleteFuente ... Delete fuente data in plan_cuentas_mongo.
func DeleteFuenteFinanciamiento(objectID string, cg string, vigencia string) (response responseformat.Response, err error) {

	if err = request.SendJson(beego.AppConfig.String("financieraMongoCurdApiService")+"fuente_financiamiento/"+objectID+"/"+vigencia+"/"+cg, "DELETE", &response, ""); err == nil {
		if responseformat.CheckResponseError(response) {
			var errMessage = "Delete Mongo API Error"
			if messageStr, e := response.Body.(string); e {
				errMessage = errMessage + ": " + messageStr
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		}
	}

	return response, err
}
