package movimientohelper

import (
	"github.com/astaxie/beego"
	models_movimientos "github.com/udistrital/movimientos_crud/models"
	"github.com/udistrital/plan_cuentas_mid/models"
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

func CrearMovimiento(movimientocreado []models_movimientos.CuentasMovimientoProcesoExterno) (movimiento []models_movimientos.MovimientoDetalle, outputError map[string]interface{}) {
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
	urlcrearmovimiento := beego.AppConfig.String("movimientosCrudService") + "movimiento_detalle/crearMovimientosDetalle/"
	if err := request.SendJson(urlcrearmovimiento, "POST", &movimiento, movimientocreado); err != nil {
		return movimiento, map[string]interface{}{
			"funcion": "CrearMovimiento - request.SendJson(urlcrearmovimiento, \"POST\", &movimiento, movimientocreado)",
			"err":     err,
			"status":  "500",
		}
	}
	return movimiento, nil
}

func CrearMovimientoProcesoExt(movimientocreado models_movimientos.MovimientoProcesoExterno) (movimiento models_movimientos.MovimientoProcesoExterno, outputError map[string]interface{}) {
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
	urlcrearmovimiento := beego.AppConfig.String("movimientosCrudService") + "movimiento_proceso_externo/"
	if err := request.SendJson(urlcrearmovimiento, "POST", &movimiento, movimientocreado); err != nil {
		return movimiento, map[string]interface{}{
			"funcion": "CrearMovimientoProcesoExt - request.SendJson(urlcrearmovimiento, \"POST\", &movimiento, movimientocreado)",
			"err":     err,
			"status":  "500",
		}
	}
	return movimiento, nil
}
