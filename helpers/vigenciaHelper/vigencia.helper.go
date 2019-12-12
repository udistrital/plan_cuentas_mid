package vigenciahelper

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// GetCierre obtiene las reservas y pasivos de una vigencia
func GetCierre(vigencia string, areaf string) (cierre map[string]interface{}, outErr map[string]interface{}) {
	var err map[string]interface{}
	cierre = make(map[string]interface{})
	if cierre["Reservas"], cierre["Pasivos"], err = GetReservas(vigencia, areaf); err != nil {
		return cierre, map[string]interface{}{"Function": "getreservas", "Error": err}
	}
	return cierre, nil
}

// GetReservas traer las reservas
func GetReservas(vigencia string, areaf string) (reservas []map[string]interface{}, pasivos []map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal/" + vigencia + "/" + areaf + "?query=tipo:rp"
	var res map[string]interface{}
	reservas = make([]map[string]interface{}, 0)
	pasivos = make([]map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getreservas", "Error": err.Error()}
		return nil, nil, outErr
	} else {
		if res["Body"] == nil {
			return reservas, pasivos, nil
		}
		for k, value := range res["Body"].([]interface{}) {
			estado := res["Body"].([]interface{})[k].(map[string]interface{})["Estado"]
			if value != nil && (estado == "expedido" || estado == "parcialmente_comprometido") {
				reservas = append(reservas, res["Body"].([]interface{})[k].(map[string]interface{}))
			}
			if value != nil && estado == "pasivo" {
				pasivos = append(pasivos, res["Body"].([]interface{})[k].(map[string]interface{}))
			}
		}
		return reservas, pasivos, nil
	}
}
