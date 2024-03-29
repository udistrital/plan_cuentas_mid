package apropiacionHelper

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/helpers"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

const success string = "success"

// AddApropiacion ... Add apropiacion to mongo and postgres tr.
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func AddApropiacion(data models.Apropiacion) map[string]interface{} {

	var (
		urlcrud   string
		res       map[string]interface{}
		mongoData map[string]interface{}
		resM      map[string]interface{}
		aprComp   []models.Apropiacion
	)

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			go func() {
				resul := res["Body"].(map[string]interface{})
				urlcrud = urlcrud + "/" + strconv.Itoa(int(resul["Id"].(float64)))
				if err := request.SendJson(urlcrud, "DELETE", &resM, nil); err == nil {
					beego.Error(helpers.ExternalAPIErrorMessage())
				} else {
					beego.Error(err.Error())
				}
			}()
			panic(helpers.InternalErrorMessage())

		}
	}()

	urlcrud = beego.AppConfig.String("planCuentasApiService") + "apropiacion"

	if err := request.GetJson(urlcrud+"apropiacion?query=RubroId.Codigo:"+data.RubroId.Codigo+",Vigencia:"+strconv.Itoa(int(data.Vigencia)), &aprComp); err != nil {
		if len(aprComp) > 0 {
			panic("Apr Exist")
		}
	} else {
		panic(err.Error())
	}

	if err := request.SendJson(urlcrud, "POST", &res, &data); err == nil {
		if res["Type"] != nil && res["Type"].(string) == success {
			if err = formatdata.FillStruct(data.RubroId, &mongoData); err != nil {
				panic(err.Error())
			}
			mongoData["Id"] = res["Body"].(map[string]interface{})["Id"]
			mongoData["ApropiacionInicial"] = data.Valor
			mongoData["UnidadEjecutora"] = data.RubroId.UnidadEjecutora
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiaciones/RegistrarApropiacionInicial/" + strconv.Itoa(int(res["Body"].(map[string]interface{})["Vigencia"].(float64)))

			if err = request.SendJson(urlmongo, "POST", &resM, &mongoData); err == nil {
				if resM["Type"].(string) == success {
					body := res["Body"].(map[string]interface{})
					body["Rubro"] = mongoData
					res["Body"] = body
				} else {
					beego.Error("res mongo", resM)
					panic(helpers.ExternalAPIErrorMessage())
				}
			} else {
				panic(err.Error())
			}

		} else {
			panic(helpers.ExternalAPIErrorMessage())
		}
	} else {
		panic(err.Error())
	}
	return res
}

// PutApropiacion ... Modify the Apr Value if these Apr isn't Aproved
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func PutApropiacion(data map[string]interface{}, idStr, valStr, vigStr string) map[string]interface{} {
	var (
		urlcrud   string
		res       map[string]interface{}
		mongoData map[string]interface{}
		resM      map[string]interface{}
	)

	valorAnt := data["ApropiacionInicial"].(float64)
	valorAntStr := strconv.Itoa(int(valorAnt))

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
			go func() {
				urlcrud = beego.AppConfig.String("planCuentasApiService") + "apropiacion/" + "/UpdateApropiacionValue/" + idStr + "/" + valorAntStr
				if err := request.SendJson(urlcrud, "PUT", &resM, nil); err == nil {
					beego.Info("Rollback TR")
				} else {
					beego.Error("Error In Rollback TR")
				}
			}()
			panic(helpers.InternalErrorMessage())

		}
	}()
	if vig, err := strconv.Atoi(vigStr); err == nil {
		aprobFlag := PresupuestoAprobado(vig, int(data["UnidadEjecutora"].(float64)))
		if aprobFlag {
			beego.Error("Apropiaciones aprobadas")
			panic(helpers.ExternalAPIErrorMessage())
		}
	} else {
		panic(helpers.InternalErrorMessage())
	}

	mongoData = data

	urlcrud = beego.AppConfig.String("planCuentasApiService") + "apropiacion/" + "UpdateApropiacionValue/" + idStr + "/" + valStr
	if err := request.SendJson(urlcrud, "PUT", &res, nil); err == nil {
		if res["Type"] != nil && res["Type"].(string) == success {
			var errStr error
			mongoData["ApropiacionInicial"], errStr = strconv.Atoi(valStr)
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiaciones/RegistrarApropiacionInicial/" + vigStr
			if err = request.SendJson(urlmongo, "POST", &resM, &mongoData); err == nil && errStr == nil {
				if resM["Type"].(string) != success {
					panic(helpers.ExternalAPIErrorMessage())
				}
			} else {
				panic(helpers.InternalErrorMessage())
			}
		} else {
			panic(helpers.ExternalAPIErrorMessage())
		}
	} else {
		panic(err.Error())
	}
	return res
}

//CalcularSaldoApropiacion .... Return Apropiation's balance by Rubro's code , ue and year.
func CalcularSaldoApropiacion(rubroParam string, unidadEParam, vigenciaParam int) (res map[string]float64) {
	var saldo float64
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiaciones/SaldoApropiacion/" + rubroParam + "/" + strconv.Itoa(unidadEParam) + "/" + strconv.Itoa(vigenciaParam)

	if err := request.GetJson(urlmongo, &res); err != nil {
		panic(err.Error())
	} else {
		for key, value := range res {

			if !strings.Contains(key, "mes") && !strings.Contains(key, "rp") {
				switch tipoMovimiento := key; tipoMovimiento {
				case "Suspensión", "mes_cdp", "Reducción":
				//rp
				case "Adición", "Traslado_cuenta_contra_credito", "total_anulado_cdp", "valor_inicial":
					saldo += value
				default:
					saldo -= value
				}
			}
		}

	}
	res["Adicion"] = res["Adición"]
	res["Reduccion"] = res["Reducción"]
	res["saldo"] = saldo
	return
}

// CompareApropiationNodes return tru if nodes to compare have the same balance.
func CompareApropiationNodes(nodesToCompare *map[string]float64, ue, vigencia int) (compareFlag bool) {

	var lastValue interface{}
	balance := make(map[string]float64)

	defer func() {
		if r := recover(); r != nil {
			beego.Error(r)
		}

	}()

	for code := range *nodesToCompare {
		defer func() {
			balance = map[string]float64{code: 0.0}
		}()
		balance = CalcularSaldoApropiacion(code, ue, vigencia)
		(*nodesToCompare)[code] = balance["saldo"]

	}

	auxMap := *nodesToCompare
	for key, value := range auxMap {

		if lastValue == nil {
			lastValue = value
			beego.Debug(key, value)
		} else {
			if lastValue.(float64) != value {

				return false
			}
		}
	}

	return true

}

// AprobarPresupuesto retorna res siempre y cuando el tipo sea success
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func AprobarPresupuesto(vigencia, unidadejecutora int) (res map[string]interface{}) {
	asignationInfo := map[string]float64{"2": 0.0, "3": 0.0}

	compareFlag := CompareApropiationNodes(&asignationInfo, unidadejecutora, vigencia)
	aprobFlag := PresupuestoAprobado(vigencia, unidadejecutora)
	if compareFlag && !aprobFlag {
		if err := request.GetJson(beego.AppConfig.String("planCuentasApiService")+"apropiacion/AprobacionAsignacionInicial"+"?Vigencia="+strconv.Itoa(vigencia)+"&UnidadEjecutora="+strconv.Itoa(unidadejecutora), &res); err == nil {
			if res["Type"] != nil && res["Type"].(string) == success {
				return
			}
			panic(helpers.ExternalAPIErrorMessage())
		}
	}

	panic(helpers.InternalErrorMessage())
}

// PresupuestoAprobado retorna true si la apropiacion esta aprobada, retorna falso en caso contrario
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func PresupuestoAprobado(vigencia, unidadejecutora int) bool {
	var res map[string]interface{}
	if err := request.GetJson(beego.AppConfig.String("planCuentasApiService")+"apropiacion?"+"query=Vigencia:"+strconv.Itoa(vigencia)+",RubroId.UnidadEjecutora:"+strconv.Itoa(unidadejecutora)+",EstadoApropiacionId:2", &res); err == nil {
		if len(res["Body"].([]interface{})) > 0 {
			return true
		}
		return false
	} else {
		beego.Error(err.Error())
		panic(helpers.InternalErrorMessage())
	}

}

// ConstruirArbolRubroApropiacion construir abrol rubro apropiacion
func ConstruirArbolRubroApropiacion(ueStr string, vigenciaStr string, raiz string, nivel string) (resultado map[string]interface{}, err interface{}) {
	var urlmongo string
	registros := make([]map[string]interface{}, 0)
	var res map[string]interface{}
	var respuesta map[string]interface{}

	urlmongo = beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiacion/arbol_apropiacion_valores/" + ueStr + "/" + vigenciaStr + "/" + raiz + "?nivel=" + nivel
	if err := request.GetJson(urlmongo, &res); err != nil {
		return nil, err
	}
	hijosNodo := res["Body"].([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})["Hijos"]
	hijos := interfaceToStringArray(hijosNodo)
	for _, element := range hijos {
		urlmongo = beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiacion/arbol_apropiacion_valores/" + ueStr + "/" + vigenciaStr + "/" + element + "?nivel=-1"
		if err := request.GetJson(urlmongo, &respuesta); err != nil {
			return nil, err
		} else {
			registros = append(registros, respuesta["Body"].([]interface{})[0].(map[string]interface{}))
		}
	}

	res["Body"].([]interface{})[0].(map[string]interface{})["children"] = registros
	return res, err

}

// interfaceToStringArray convierte array de interface{} a string
func interfaceToStringArray(hijos interface{}) (Hijos []string) {
	aInterface := hijos.([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString

}

// // interfaceToStringArray convierte array de interface{} a string
// func UnirArbol(registros []map[string]interface{}, ) (resultado map[string]interface{}) {
// 	aInterface := hijos.([]interface{})
// 	aString := make([]string, len(aInterface))
// 	for i, v := range aInterface {
// 		aString[i] = v.(string)
// 	}
// 	return aString

// }
