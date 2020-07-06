package movimientohelper

import (
	"github.com/udistrital/plan_cuentas_mid/models"
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
