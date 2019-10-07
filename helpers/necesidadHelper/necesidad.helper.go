package necesidadhelper

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// InitNecesidad init necesidad
func InitNecesidad(id string) (necesidad map[string]interface{}, outputError map[string]interface{}) {
	var (
		necesidadPC  map[string][]interface{}
		necesidadADM map[string][]interface{}
	)
	necesidad = make(map[string]interface{})
	necesidad["id"] = id
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "necesidades/?query=idAdministrativa:" + id
	urladm := beego.AppConfig.String("adminCrudService") + "necesidad/?query=Id:" + id
	if err := request.GetJson(urladm, &necesidadADM); err != nil {
		panic("ADM API Service Error")
	} else {
		beego.Info(necesidadADM)
	}
	if err := request.GetJson(urlmongo, &necesidadPC); err != nil {
		panic("Mongo API Service Error")
	} else {
		necesidad["NecesidadPC"] = necesidadPC["Body"][0]
		necesidad["NecesidadADM"] = necesidadADM
		return necesidad, nil
	}
}
