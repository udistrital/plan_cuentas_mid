package crphelper

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/helpers"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// RegistrarSolicitudCrp ... Add apropiacion to mongo and postgres tr.
func RegistrarSolicitudCrp(data models.SolicitudCrp) map[string]interface{} {

	var (
		urlcrud   string
		res       map[string]interface{}
		mongoData map[string]interface{}
		resM      map[string]interface{}
		solCrpObj []models.SolicitudCrp
	)

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			go func() {
				resul := res["Body"].(map[string]interface{})
				urlcrud = urlcrud + "/" + strconv.Itoa(int(resul["_id"].(float64)))
				if err := request.SendJson(urlcrud, "DELETE", &resM, nil); err == nil {
					beego.Error(helpers.ExternalAPIErrorMessage())
				} else {
					beego.Error(err.Error())
				}
			}()
			panic(helpers.InternalErrorMessage())

		}
	}()

	urlcrud = beego.AppConfig.String("planCuentasApiService") + "solicitudesCRP"

}
