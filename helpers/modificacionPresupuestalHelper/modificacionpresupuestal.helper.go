package modificacionpresupuestalhelper

import (
	"time"

	"github.com/udistrital/plan_cuentas_mid/models"
)

func ConvertModificacionToDocumentoPresupuestal(modData models.ModificacionPresupuestalReceiver) (dataFormated models.DocumentoPresupuestal) {
	var movimientos []models.Movimiento
	currDate := time.Now()
	dataFormated.Tipo = "modificacion_presupuestal"
	dataFormated.Vigencia = currDate.Year()
	dataFormated.CentroGestor = modData.Data.CentroGestor
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
