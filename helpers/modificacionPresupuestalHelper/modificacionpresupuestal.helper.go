package modificacionpresupuestalhelper

import (
	"github.com/udistrital/plan_cuentas_mid/models"
)

func ConvertModificacionToDocumentoPresupuestal(modData models.ModificacionPresupuestalReceiver) (dataFormated models.DocumentoPresupuestal) {
	var movimientos []models.Movimiento
	dataFormated.Data = modData.Data
	dataFormated.Tipo = "modificacion_presupuestal"
	dataFormated.Vigencia = 2020
	dataFormated.CentroGestor = "1"
	for _, afectation := range modData.Afectation {
		movimiento := models.Movimiento{}
		movimiento.Descripcion = modData.Data.Descripcion
		movimiento.DocumentoPadre = afectation.OriginAcc.Codigo
		movimiento.Valor = afectation.Amount
		movimiento.MovimientoProcesoExternoId = &models.MovimientoProcesoExterno{
			TipoMovimientoId: afectation.TypeMod,
		}
		movimientos = append(movimientos, movimiento)
	}
	dataFormated.AfectacionMovimiento = movimientos
	return
}
