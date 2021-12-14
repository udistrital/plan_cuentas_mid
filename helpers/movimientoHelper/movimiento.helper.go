package movimientohelper

import (
	"github.com/astaxie/beego"
	models_movimientos "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_cuentas_mid/helpers/utils"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// FormatDataForMovimientosMongoAPI ... format Movimiento Data to MoimientosMOngoAPI data structure.
func FormatDataForMovimientosMongoAPI(data ...models.Movimiento) (dataFormated []models.MovimientoMongo) {
	for _, movimiento := range data {
		var element models.MovimientoMongo
		element.Tipo = movimiento.MovimientoProcesoExternoId.TipoMovimientoId.Acronimo
		element.Padre = movimiento.DocumentoPadre
		element.Descripcion = movimiento.Descripcion
		element.ValorInicial = movimiento.Valor
		element.IDPsql = movimiento.Id
		element.FechaRegistro = movimiento.FechaRegistro
		dataFormated = append(dataFormated, element)
	}

	return
}

//ObtenerUltimoMovimiento, obtiene el ultimo movimiento realizado por los datos de rubro y fuentes (actividades si es de inversion)
func ObtenerUltimoMovimiento(ultMovimiento []models_movimientos.CuentasMovimientoProcesoExterno) (movimiento []models_movimientos.MovimientoDetalle, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "ObtenerUltimoMovimiento - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	urlultimomovimiento := beego.AppConfig.String("movimientosCrudService") + "movimiento_detalle/postUltimoMovDetalle/"
	if err := request.SendJson(urlultimomovimiento, "POST", &movimiento, ultMovimiento); err != nil {
		return movimiento, map[string]interface{}{
			"funcion": "ObtenerUltimoMovimiento - request.SendJson(urlultimomovimiento, \"POST\", &movimiento, ultMovimiento)",
			"err":     err,
			"status":  "500",
		}
	}
	return movimiento, nil
}

//CrearMovimiento, Crea un resgistro del movimiento realizado a partir de un movimiento proceso externo y los datos de la necesidad
func CrearMovimiento(movimientocreado []models_movimientos.CuentasMovimientoProcesoExterno) (movimiento []models_movimientos.MovimientoDetalle, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "CrearMovimiento - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	urlcrearmovimiento := beego.AppConfig.String("movimientosCrudService") + "movimiento_detalle/crearMovimientosDetalle/"
	if err := request.SendJson(urlcrearmovimiento, "POST", &movimiento, movimientocreado); err != nil {
		return movimiento, map[string]interface{}{
			"funcion": "CrearMovimiento - request.SendJson(urlcrearmovimiento, \"POST\", &respuestaPeticion, movimientocreado)",
			"err":     err,
			"status":  "500",
		}
	}

	return movimiento, nil
}

//CrearMovimientoProcesoExt, A partir de la informacion que se envie, crea un movimiento proceso externo con su Id respectivo
func CrearMovimientoProcesoExt(movimientocreado models_movimientos.MovimientoProcesoExterno) (movimiento models_movimientos.MovimientoProcesoExterno, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "CrearMovimientoProcesoExt - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	var respuestaPeticion utils.RespuestaEncapsulada1

	urlcrearmovimiento := beego.AppConfig.String("movimientosCrudService") + "movimiento_proceso_externo/"
	if err := request.SendJson(urlcrearmovimiento, "POST", &respuestaPeticion, movimientocreado); err != nil {
		return movimiento, map[string]interface{}{
			"funcion": "CrearMovimientoProcesoExt - request.SendJson(urlcrearmovimiento, \"POST\", &movimiento, movimientocreado)",
			"err":     err,
			"status":  "500",
		}
	}

	formatdata.FillStruct(respuestaPeticion.Body, &movimiento)

	return movimiento, nil
}
