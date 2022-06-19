package fuenteapropiacionhelper

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// URLCRUD Path de plan_cuentas_crud
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
var URLCRUD = beego.AppConfig.String("planCuentasApiService")

// URLPLANADQUISICION Path de bodega jbpm para el plan de adquisiciones
var URLPLANADQUISICION = beego.AppConfig.String("bodegaService")

// URLMOVIMIENTOSCRUD Path de movimientos_crud
var URLMOVIMIENTOSCRUD = beego.AppConfig.String("movimientosCrudService")

// GetPlanAdquisicionbyFuente obtiene información de la fuente y totaliza los valores acordes a a cada una de ellas
func GetPlanAdquisicionbyFuente(vigencia string, id string) (fuente map[string]map[string]interface{}, outErr map[string]interface{}) {
	requestPath := URLPLANADQUISICION + "plan_adquisiciones_rubros_fuente/" + vigencia + "/" + id
	var res map[string]map[string]interface{}
	//request.SetHeader("application/json")
	if err := request.GetJsonWSO2(requestPath, &res); err != nil {
		outErr = map[string]interface{}{"Function": "GetPlanAdquisicionbyFuente", "Error": err.Error()}
		return nil, outErr
	}
	if res["Fault"] != nil {
		outErr = map[string]interface{}{"Function": "GetPlanAdquisicionbyFuente", "Error": res}
		return nil, outErr
	}
	fuente = AddTotal(res)
	return fuente, nil
}

// AddTotal suma los totales de las fuentes de financiamiento
func AddTotal(res map[string]map[string]interface{}) (newres map[string]map[string]interface{}) {

	var totalValue float64
	if err := formatdata.JsonPrint(len(res["fuente_financiamiento"]["rubros"].([]interface{}))); err != nil {
		log.Panicln(err.Error())
	}
	for _, item := range res["fuente_financiamiento"]["rubros"].([]interface{}) {
		if floatValue, err := strconv.ParseFloat(item.(map[string]interface{})["valor_fuente_financiamiento"].(string), 64); err == nil {
			item.(map[string]interface{})["valor_fuente_financiamiento"] = strconv.FormatFloat(floatValue, 'f', 0, 64)
			totalValue = totalValue + floatValue
		} else {
			panic(err.Error())
		}
	}
	res["fuente_financiamiento"]["total_saldo_fuente"] = strconv.FormatFloat(totalValue, 'f', 0, 64)
	newres = res
	return newres
}

// ConvertModificacionToDocumentoPresupuestal Convierte Modificación Fuente a Documento Presupuestal
func ConvertModificacionToDocumentoPresupuestal(modData models.ModificacionFuenteReceiver) (dataFormated models.DocumentoPresupuestal) {
	var movimientos []models.Movimiento
	currDate := time.Now()
	dataFormated.Tipo = "modificacion_fuente"
	dataFormated.Vigencia = currDate.Year()
	dataFormated.CentroGestor = modData.Data.CentroGestor
	for _, afectation := range modData.Afectation {
		if afectation.OriginAcc != nil {
			movimiento := models.Movimiento{}
			movimiento.Descripcion = modData.Data.Descripcion
			movimiento.DocumentoPadre = afectation.OriginAcc.Codigo
			movimiento.Valor = afectation.Amount
			movimiento.MovimientoProcesoExternoId = &models.MovimientoProcesoExterno{
				TipoMovimientoId: afectation.TypeMod,
			}
			movimientos = append(movimientos, movimiento)
		}
		if afectation.TypeMod.Parametros != "" && afectation.OriginRubro != nil {
			parametersMap := make(map[string]interface{})
			if err := json.Unmarshal([]byte(afectation.TypeMod.Parametros), &parametersMap); err != nil {
				panic(err.Error())
			}
			if originRubroType, e := parametersMap["MovOriginRubro"].(string); e {
				movimientoOriginRubro := models.Movimiento{}
				movimientoOriginRubro.Descripcion = modData.Data.Descripcion
				movimientoOriginRubro.DocumentoPadre = afectation.OriginRubro.Codigo
				movimientoOriginRubro.Valor = afectation.Amount
				movimientoOriginRubro.MovimientoProcesoExternoId = &models.MovimientoProcesoExterno{
					TipoMovimientoId: &models.TipoGeneral{
						Id:       afectation.TypeMod.Id,
						Acronimo: originRubroType,
					},
				}
				movimientos = append(movimientos, movimientoOriginRubro)
			}
		}
	}
	dataFormated.AfectacionMovimiento = movimientos
	return
}

// RegistrarMultipleFuenteApropiacion utiliza la transacción de fuente_financiamiento_apropiacion/registrar_multiple
// para registrar multiples datos en fuente_financiamiento_apropiacion
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func RegistrarMultipleFuenteApropiacion(fuentesApropiacion []*models.FuenteFinanciamientoApropiacion) (idRegistrados []int) {
	var (
		res          map[string]interface{} // Respuesta de una petición
		bodyResponse map[string][]int
	)

	requestPath := URLCRUD + "fuente_financiamiento_apropiacion/registrar_multiple"
	if err := request.SendJson(requestPath, "POST", &res, &fuentesApropiacion); err != nil {
		log.Panicln(err.Error())
		return
	}

	if err := formatdata.FillStruct(res["Body"], &bodyResponse); err != nil {
		log.Panicln(err.Error())
		return
	}
	idRegistrados = bodyResponse["Ids"]
	return
}

// RegistrarMultipleMovimientoExterno utiliza la transacción registrar_multiple de movimientos_crud
// para registrar multiples movimientos en una sola petición
func RegistrarMultipleMovimientoExterno(data interface{}) (idRegistrados []int) {
	var (
		res          map[string]interface{} // res: Respuesta de la petición
		bodyResponse map[string][]int       // Cuerpo de respuesta de la petición
	)

	requestPath := URLMOVIMIENTOSCRUD + "movimiento_detalle/registrar_multiple"
	if err := request.SendJson(requestPath, "POST", &res, &data); err != nil {
		log.Panicln(err.Error())
		return
	}

	if err := formatdata.FillStruct(res["Body"], &bodyResponse); err != nil {
		log.Panicln(err.Error())
		return
	}

	idRegistrados = bodyResponse["Ids"]

	return
}

// FormatDataMovimientoExterno Establece el formato para utilizar como parámetros en la transacción movimiento_detalle/registrar_multiple
// @param idsFuentes: id de los registros en fuente_financiamiento_apropiacion
// @param fuenteApropiaciones: la información del atributo FuentesFinanciamientoApropiacion enviado en la petición
func FormatDataMovimientoExterno(idsFuentes []int, fuenteApropiaciones ...interface{}) (dataArray []map[string]interface{}) {
	var valor float64 // valor enviado al movimiento

	for i, fuente := range fuenteApropiaciones {

		if err := formatdata.FillStruct(fuente.(map[string]interface{})["Valor"], &valor); err != nil {
			log.Panicln(err.Error())
			return nil
		}

		data := map[string]interface{}{
			"Valor":         valor,
			"FechaRegistro": time.Now(),
			"MovimientoProcesoExternoId": map[string]interface{}{
				"TipoMovimientoId": map[string]int{"Id": 3},
				"ProcesoExterno":   idsFuentes[i],
			},
		}

		dataArray = append(dataArray, data)
	}
	return
}

// ConcatenarFuente recibe un arreglo de interfaces y las contatena para formar un arreglo de models.FuenteFinanciamientoApropiacion
func ConcatenarFuente(fuenteFinanciamiento *models.FuenteFinanciamiento, fuenteApropiaciones ...interface{}) []*models.FuenteFinanciamientoApropiacion {
	var (
		apropiacion        *models.Apropiacion
		fuenteApropiacion  *models.FuenteFinanciamientoApropiacion   // una sola fuente
		fuentesApropiacion []*models.FuenteFinanciamientoApropiacion // Un conjunto de fuentes apropiacion
	)

	for _, value := range fuenteApropiaciones {
		// Obtiene la información de la apropición
		if err := formatdata.FillStruct(value.(map[string]interface{})["Apropiacion"], &apropiacion); err != nil {
			log.Panicln(err.Error())
			return nil
		}

		/* Convierte toda la información de value a tipo fuenteFinanciamiento
		(no se pueden convertir automáticamente structs definidos, como apropiacion y fuenteFinanciamiento) */
		if err := formatdata.FillStruct(value, &fuenteApropiacion); err != nil {
			log.Panicln(err.Error())
			return nil
		}
		fuenteApropiacion.ApropiacionId = apropiacion
		fuenteApropiacion.FuenteFinanciamientoId = fuenteFinanciamiento

		fuentesApropiacion = append(fuentesApropiacion, fuenteApropiacion)
	}

	return fuentesApropiacion
}
