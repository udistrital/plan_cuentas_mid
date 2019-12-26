package vigenciahelper

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// GetCierre obtiene las reservas y pasivos de una vigencia
func GetCierre(vigencia string, areaf string, cerrada bool) (cierre map[string]interface{}, outErr map[string]interface{}) {
	var err map[string]interface{}
	cierre = make(map[string]interface{})
	cierre["Reservas"], cierre["Pasivos"], err = GetReservas(vigencia, areaf, cerrada)
	cierre["Fuentes"], err = GetFuentesCierre(vigencia, areaf)
	if err != nil {
		return cierre, map[string]interface{}{"Function": "getreservas", "Error": err}
	}
	return cierre, nil
}

// GetReservas traer las reservas
func GetReservas(vigencia string, areaf string, cerrada bool) (reservas []map[string]interface{}, pasivos []map[string]interface{}, outErr map[string]interface{}) {
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
			var condicionR bool
			var condicionP bool
			estado := res["Body"].([]interface{})[k].(map[string]interface{})["Estado"]
			if cerrada {
				condicionR = (estado == "reserva")
				condicionP = (estado == "pasivo")
			} else {
				condicionR = (estado == "expedido" || estado == "parcialmente_comprometido")
				condicionP = (estado == "reserva")
			}
			if value != nil && condicionR {
				reservas = append(reservas, res["Body"].([]interface{})[k].(map[string]interface{}))
			}
			if value != nil && condicionP {
				pasivos = append(pasivos, res["Body"].([]interface{})[k].(map[string]interface{}))
			}
		}
		return reservas, pasivos, nil
	}
}

// GetFuentesCierre trae las fuentes que aun tienen saldo disponible
func GetFuentesCierre(vigencia string, areaf string) (fuentes []map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento/" + vigencia + "/" + areaf
	var res map[string]interface{}
	fuentes = make([]map[string]interface{}, 0)
	fmt.Println(urlcrud)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "GetFuentesCierre", "Error": err.Error()}
		return nil, outErr
	} else {
		if res["Body"] == nil {
			return fuentes, nil
		}
		for k, value := range res["Body"].([]interface{}) {
			valor := res["Body"].([]interface{})[k].(map[string]interface{})["ValorActual"]
			if value != nil && (valor.(float64) > 0) {
				fuentes = append(fuentes, res["Body"].([]interface{})[k].(map[string]interface{}))
			}
		}
		return fuentes, nil
	}
}

// CerrarVigencia realiza los procesos de cierre de vigencia
func CerrarVigencia(vigencia string, areaf string) (cierre map[string]interface{}, outErr map[string]interface{}) {
	cierre = make(map[string]interface{})
	reservas, pasivos, err := GetReservas(vigencia, areaf, false)
	if err == nil {
		for _, reserva := range reservas {
			cambiarEstadoRP(vigencia, areaf, reserva["_id"].(string), "reserva")
		}
		for _, pasivo := range pasivos {
			cambiarEstadoRP(vigencia, areaf, pasivo["_id"].(string), "pasivo")
		}
		var response map[string]interface{}
		if errorCrud := request.GetJson(beego.AppConfig.String("financieraMongoCurdApiService")+"/vigencia/cerrar_vigencia_actual/"+areaf, &response); errorCrud == nil {
			return cierre, nil
		} else {
			return cierre, map[string]interface{}{"Function": "CerrarVigencia", "Error": err}
		}
	} else {
		return cierre, map[string]interface{}{"Function": "CerrarVigencia", "Error": err}
	}

}

// cambiar estado rp
func cambiarEstadoRP(vigencia string, areaf string, id string, estado string) (res map[string]interface{}, outErr map[string]interface{}) {
	var response map[string]interface{}
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal/"
	if err := request.GetJson(urlmongo+"documento/"+vigencia+"/"+areaf+"/"+id, &response); err != nil {
		outErr = map[string]interface{}{"Function": "CerrarVigencia", "Error": "No se pudo obtener el documento"}
		return nil, outErr
	} else {
		docPresupuestal := response["Body"].(map[string]interface{})
		docPresupuestal["Estado"] = estado

		if err = request.SendJson(urlmongo+vigencia+"/"+areaf+"/"+id, "PUT", &response, &docPresupuestal); err != nil {
			return nil, map[string]interface{}{"Function": "cambiarEstadoRP", "Error": "No se actualizar el documento"}
		} else {
			res = response
			return res, nil
		}
	}

}
