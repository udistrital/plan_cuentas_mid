package modificacionpresupuestalhelper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
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

func FormatDocumentoPresupuestalResponseToModificacionDetail(rows []models.DocumentoPresupuestal) []models.ModificacionPresupuestalResponseDetail {
	var finalData []models.ModificacionPresupuestalResponseDetail
	for _, row := range rows {
		var modificacion models.ModificacionPresupuestalResponseDetail
		var dataMap map[string]interface{}

		defer func() {
			if r := recover(); r != nil {
				// do nothing ...
				fmt.Println("error", r)
			}
		}()
		formatdata.FillStruct(row.Data, &dataMap)

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

		finalData = append(finalData, modificacion)
	}
	return finalData
}

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
		formatdata.FillStruct(row.Data, &dataMap)
		anulacion.Consecutivo = row.Consecutivo
		anulacion.Tipo, _ = dataMap["tipo_anulacion"].(string)
		anulacion.FechaRegistro = row.FechaRegistro
		anulacion.Descripcion, _ = dataMap["descripcion"].(string)
		anulacion.Valor = row.ValorActual

		finalData = append(finalData, anulacion)
	}
	return finalData
}
