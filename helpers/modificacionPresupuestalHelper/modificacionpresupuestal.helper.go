package modificacionpresupuestalhelper

import (
	"encoding/json"
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
		if afectation.TypeMod.Parametros != "" && afectation.TargetAcc.Codigo != "" {
			parametersMap := make(map[string]interface{})
			if err := json.Unmarshal([]byte(afectation.TypeMod.Parametros), &parametersMap); err != nil {
				panic(err.Error())
			}
			if targetAccType, e := parametersMap["TipoMovimientoCuentaCredito"].(string); e {
				movimientoTargetAcc := models.Movimiento{}
				movimientoTargetAcc.Descripcion = modData.Data.Descripcion
				movimientoTargetAcc.DocumentoPadre = afectation.TargetAcc.Codigo
				movimientoTargetAcc.Valor = afectation.Amount
				movimientoTargetAcc.MovimientoProcesoExternoId = &models.MovimientoProcesoExterno{
					TipoMovimientoId: &models.TipoGeneral{
						Id:       afectation.TypeMod.Id,
						Acronimo: targetAccType,
					},
				}
				movimientos = append(movimientos, movimientoTargetAcc)
			}
		}
		movimientos = append(movimientos, movimiento)
	}
	dataFormated.AfectacionMovimiento = movimientos
	return
}
