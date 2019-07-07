package fuenteapropiacionhelper

import (
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// URLCRUD Path de plan_cuentas_crud
var URLCRUD = beego.AppConfig.String("planCuentasApiService")

// URLMOVIMIENTOSCRUD Path de movimientos_crud
var URLMOVIMIENTOSCRUD = beego.AppConfig.String("movimientosCrudService")

// RegistrarMultipleFuenteApropiacion utiliza la transacción de fuente_financiamiento_apropiacion/registrar_multiple
// para registrar multiples datos en fuente_financiamiento_apropiacion
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

// formatDataMovimientoExterno Establece el formato para utilizar como parámetros en la transacción movimiento_detalle/registrar_multiple
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
