package necesidadhelper

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// InitNecesidad init necesidad
func InitNecesidad(id string) (necesidad map[string]interface{}, err map[string]interface{}) {
	var (
		necesidadPC map[string][]interface{}
	)
	necesidad = make(map[string]interface{})
	necesidad["id"] = id
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "necesidades/?query=idAdministrativa:" + id
	if necesidadADM, err := GetNecesidadADM(id); err != nil {
		panic(err)
	} else {
		necesidad["necesidadADM"] = necesidadADM
	}
	if err := request.GetJson(urlmongo, &necesidadPC); err != nil {
		panic(err)
	} else {
		necesidad["NecesidadPC"] = necesidadPC["Body"][0]
		return necesidad, nil
	}
}

// GetNecesidadADM obtiene necesidad de adm crud apí
func GetNecesidadADM(id string) (necesidad map[string]interface{}, outErr map[string]interface{}) {
	var necesidadADM []map[string]interface{}
	var necesidadAux models.NecesidadADM
	urladm := beego.AppConfig.String("adminCrudService") + "necesidad/?query=Id:" + id
	if err := request.GetJson(urladm, &necesidadADM); err != nil {
		outErr = map[string]interface{}{"Function": "GetNecesidadADM", "Error": err.Error()}
		return nil, outErr
	} else {
		if err := formatdata.FillStruct(necesidadADM[0], &necesidadAux); err == nil {
			necesidad, err = formatdata.ToMap(necesidadAux, "json")
			return necesidad, nil
		} else {
			outErr = map[string]interface{}{"Function": "GetNecesidadADM", "Error": err.Error()}
			return nil, outErr
		}
	}
}
