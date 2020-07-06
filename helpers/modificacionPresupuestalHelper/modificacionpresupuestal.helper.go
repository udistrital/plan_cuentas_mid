package modificacionpresupuestalhelper

import (
	"encoding/json"
	"fmt"

	"github.com/udistrital/plan_cuentas_mid/managers"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// ConvertModificacionToDocumentoPresupuestal ...
func ConvertModificacionToDocumentoPresupuestal(modData models.ModificacionPresupuestalReceiver) (dataFormated models.DocumentoPresupuestal) {
	var movimientos []models.Movimiento
	vigenciaManager := managers.VigenciaManager{}
	// currDate := time.Now()
	dataFormated.Tipo = "modificacion_presupuestal"
	currVigencia, err := vigenciaManager.GetCurrentActiveVigencia(modData.Data.CentroGestor)
	if err != nil {
		panic(err.Error())
	}
	dataFormated.Vigencia = currVigencia
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

// FormatDocumentoPresupuestalResponseToModificacionDetail ...
func FormatDocumentoPresupuestalResponseToModificacionDetail(rows []models.DocumentoPresupuestal) []models.ModificacionPresupuestalResponseDetail {
	var finalData []models.ModificacionPresupuestalResponseDetail
	for _, row := range rows {

		defer func() {
			if r := recover(); r != nil {
				// do nothing ...
				fmt.Println("error", r)
			}

		}()

		modificacion := FormatDocumentoPresupuestalToModificacion(row)

		finalData = append(finalData, modificacion)
	}
	return finalData
}

// FormatDocumentoPresupuestalToModificacion ...
func FormatDocumentoPresupuestalToModificacion(row models.DocumentoPresupuestal) models.ModificacionPresupuestalResponseDetail {
	var modificacion models.ModificacionPresupuestalResponseDetail
	var dataMap map[string]interface{}

	if err := formatdata.FillStruct(row.Data, &dataMap); err != nil {
		panic(err.Error())
	}

	modificacion.DocumentNumber, _ = dataMap["numero_documento"].(string)
	modificacion.DocumentDate, _ = dataMap["fecha_documento"].(string)
	modificacion.DocumentType, _ = dataMap["tipo_documento"].(map[string]interface{})["Nombre"].(string)
	modificacion.CentroGestor = row.CentroGestor
	modificacion.Descripcion, _ = dataMap["descripcion_documento"].(string)
	modificacion.OrganismoEmisor, _ = dataMap["organismo_emisor"].(string)
	modificacion.RegistrationDate = row.FechaRegistro
	modificacion.ID = row.ID
	modificacion.Vigencia = row.Vigencia
	modificacion.ValorActual = row.ValorActual
	modificacion.ValorInicial = row.ValorInicial

	return modificacion

}

// FormatDocumentoPresupuestalResponseToAnulationDetail ...
func FormatDocumentoPresupuestalResponseToAnulationDetail(rows []models.DocumentoPresupuestal) []models.AnulationDetail {
	var finalData []models.AnulationDetail
	for _, row := range rows {
		anulacion := models.AnulationDetail{}
		var dataMap map[string]interface{}

		defer func() {
			if r := recover(); r != nil {
				// do nothing ...
				fmt.Println("error", r)
			}
		}()
		if err := formatdata.FillStruct(row.Data, &dataMap); err != nil {
			panic(err.Error())
		}
		anulacion.Consecutivo = row.Consecutivo
		anulacion.Tipo, _ = dataMap["tipo_anulacion"].(string)
		anulacion.FechaRegistro = row.FechaRegistro
		anulacion.Descripcion, _ = dataMap["descripcion"].(string)
		anulacion.Valor = row.ValorActual

		finalData = append(finalData, anulacion)
	}
	return finalData
}
