package movimientohelper

import (
	"github.com/udistrital/plan_cuentas_mid/models"
)

// FormatDataForMovimientosMongoAPI ... format Movimiento Data to MoimientosMOngoAPI data structure.
func FormatDataForMovimientosMongoAPI(data ...models.Movimiento) (dataFormated []models.MovimientoMongo) {
	for _, movimiento := range data {
		var element models.MovimientoMongo
		element.Tipo = movimiento.MovimientoProcesoExternoId.TipoMovimientoId.Acronimo
		element.DocumentoPadre = movimiento.MovimientoProcesoExternoId.ProcesoExterno
		element.Descripcion = movimiento.Descripcion
		element.Valor = movimiento.Valor
		element.IDPsql = movimiento.Id
		element.FechaRegistro = movimiento.FechaRegistro
		dataFormated = append(dataFormated, element)
	}

	return
}
