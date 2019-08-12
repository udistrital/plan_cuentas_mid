package test

import (
	"testing"
	"time"

	"github.com/udistrital/utils_oas/formatdata"

	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
)

func TestFormatDataForMovimientosMongoAPI(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Error("A unexpected error ocurred!!")
			t.Fail()
		}
	}()

	movimientosTestData := models.Movimiento{}
	movimientosTestData.MovimientoProcesoExternoId = &models.MovimientoProcesoExterno{}
	movimientosTestData.MovimientoProcesoExternoId.TipoMovimientoId = &models.TipoMovimiento{}
	movimientosTestData.Descripcion = "test"
	movimientosTestData.FechaRegistro = time.Now().Format(time.RFC3339)
	movimientosTestData.MovimientoProcesoExternoId.ProcesoExterno = 3
	movimientosTestData.MovimientoProcesoExternoId.TipoMovimientoId.Acronimo = "rdc_apr"
	movimientosTestData.MovimientoProcesoExternoId.TipoMovimientoId.Id = 1
	movimientosTestData.Valor = 500
	testData := movimientohelper.FormatDataForMovimientosMongoAPI(movimientosTestData)
	err := formatdata.StructValidation(movimientosTestData)
	if len(testData) > 0 && len(err) == 0 {
		t.Log("Format MovimietosAPI data Success!! ")
		t.Log(testData)
	} else {
		t.Error(err)
		t.Error("Error at FormatDataForMovimientosMongoAPI function")
		t.Fail()
	}
}
