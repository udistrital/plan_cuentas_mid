package test

import (
	"testing"

	modificacionpresupuestalhelper "github.com/udistrital/plan_cuentas_mid/helpers/modificacionPresupuestalHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
)

func TestFormatDocumentoPresupuestalResponseToModificacionDetail(t *testing.T) {
	var testData []models.DocumentoPresupuestal
	var testDocumento models.DocumentoPresupuestal
	var testDataInterface interface{}

	testDataInterface = map[string]interface{}{
		"numero_documento":      "test",
		"fecha_documento":       "test",
		"tipo_documento":        map[string]interface{}{"Nombre": "test"},
		"descripcion_documento": "test",
		"organismo_emisor":      "test",
	}

	testDocumento.ID = "test"
	testDocumento.Tipo = "test"
	testDocumento.AfectacionIds = append(testDocumento.AfectacionIds, "test")
	testDocumento.Vigencia = 2020
	testDocumento.CentroGestor = "test"
	testDocumento.FechaRegistro = "test"
	testDocumento.Consecutivo = 1
	testDocumento.ValorActual = 1
	testDocumento.ValorInicial = 1
	testDocumento.Data = testDataInterface

	testData = append(testData, testDocumento)
	testOutput := modificacionpresupuestalhelper.FormatDocumentoPresupuestalResponseToModificacionDetail(testData)

	if len(testOutput) > 0 {
		t.Logf("TestFormatDocumentoPresupuestalResponseToModificacionDetail Success!!")
	} else {
		t.Error("TestFormatDocumentoPresupuestalResponseToModificacionDetail Fail!!")
		t.Fail()
	}

}

func TestFormatDocumentoPresupuestalResponseToAnulationDetail(t *testing.T) {
	var testData []models.DocumentoPresupuestal
	var testDocumento models.DocumentoPresupuestal
	var testDataInterface interface{}

	testDataInterface = map[string]interface{}{
		"numero_documento":      "test",
		"fecha_documento":       "test",
		"tipo_documento":        map[string]interface{}{"Nombre": "test"},
		"descripcion_documento": "test",
		"organismo_emisor":      "test",
	}

	testDocumento.ID = "test"
	testDocumento.Tipo = "test"
	testDocumento.AfectacionIds = append(testDocumento.AfectacionIds, "test")
	testDocumento.Vigencia = 2020
	testDocumento.CentroGestor = "test"
	testDocumento.FechaRegistro = "test"
	testDocumento.Consecutivo = 1
	testDocumento.ValorActual = 1
	testDocumento.ValorInicial = 1
	testDocumento.Data = testDataInterface

	testData = append(testData, testDocumento)

	testOutput := modificacionpresupuestalhelper.FormatDocumentoPresupuestalResponseToAnulationDetail(testData)

	if len(testOutput) > 0 {
		t.Logf("TestFormatDocumentoPresupuestalResponseToAnulationDetail Success!!")
	} else {
		t.Error("TestFormatDocumentoPresupuestalResponseToAnulationDetail Fail!!")
		t.Fail()
	}
}
