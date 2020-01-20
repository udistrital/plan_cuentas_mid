package test

import (
	fuenteapropiacionhelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"

	"reflect"
	"testing"
	"time"
)

func TestAddTotal(t *testing.T) {
	var testData map[string]map[string]interface{}
	var testDataMap []interface{}

	testDataMap = append(testDataMap,
		map[string]interface{}{"valor_fuente_financiamiento": "1"},
		map[string]interface{}{"valor_fuente_financiamiento": "2"},
		map[string]interface{}{"valor_fuente_financiamiento": "3"},
		map[string]interface{}{"valor_fuente_financiamiento": "4"},
		map[string]interface{}{"valor_fuente_financiamiento": "5"},
		map[string]interface{}{"valor_fuente_financiamiento": "80"},
		map[string]interface{}{"valor_fuente_financiamiento": "35"},
		map[string]interface{}{"valor_fuente_financiamiento": "30"})

	testData = make(map[string]map[string]interface{})
	testData["fuente_financiamiento"] = map[string]interface{}{"rubros": testDataMap}

	testOutput := fuenteapropiacionhelper.AddTotal(testData)
	result := testOutput["fuente_financiamiento"]["total_saldo_fuente"]
	if result == "160" {
		t.Logf("AddTotal Success, expected 160, got %v", result)
	} else {
		t.Errorf("AddTotal Fail, expected 160, got %v", result)
		t.Fail()
	}
}

func TestConvertModificacionToDocumentoPresupuestal(t *testing.T) {
	var testData models.ModificacionFuenteReceiver
	var dataObj models.ModificacionPresupuestalReceiverDetail
	var testTipoGeneral models.TipoGeneral
	var testRubro models.Rubro
	var testFuenteFinanciamiento models.FuenteFinanciamiento
	var testTipoFuenteFinanciamiento models.TipoFuenteFinanciamiento
	var testModificacionReciever models.ModificacionFuenteReceiverAfectation

	testTipoGeneral.Id = 1
	testTipoGeneral.Nombre = "test"
	testTipoGeneral.Descripcion = "test"
	testTipoGeneral.Acronimo = "test"
	testTipoGeneral.Parametros = "{\"test\" : true }"

	dataObj.DocumentNumber = "test"
	dataObj.DocumentDate = time.Now()
	dataObj.Descripcion = "test"
	dataObj.CentroGestor = "test"
	dataObj.OrganismoEmisor = "test"
	dataObj.DocumentType = &testTipoGeneral

	testData.Data = &dataObj

	testRubro.Id = 1
	testRubro.Organizacion = 1
	testRubro.Codigo = "test"
	testRubro.Descripcion = "test"
	testRubro.UnidadEjecutora = "test"
	testRubro.Nombre = "test"

	testTipoFuenteFinanciamiento.Id = 1
	testTipoFuenteFinanciamiento.Nombre = "test"
	testTipoFuenteFinanciamiento.Descripcion = "test"

	testFuenteFinanciamiento.Id = 1
	testFuenteFinanciamiento.Nombre = "test"
	testFuenteFinanciamiento.Descripcion = "test"
	testFuenteFinanciamiento.Codigo = "test"
	testFuenteFinanciamiento.TipoFuenteFinanciamiento = &testTipoFuenteFinanciamiento

	testModificacionReciever.OriginAcc = &testFuenteFinanciamiento
	testModificacionReciever.OriginRubro = &testRubro
	testModificacionReciever.TypeMod = &testTipoGeneral
	testModificacionReciever.Amount = 1
	testData.Afectation = append(testData.Afectation, &testModificacionReciever,
		&testModificacionReciever)

	testOutput := fuenteapropiacionhelper.ConvertModificacionToDocumentoPresupuestal(testData)
	response := reflect.TypeOf(testOutput)
	var testPrueba models.DocumentoPresupuestal
	err := formatdata.StructValidation(testOutput)
	if response == reflect.TypeOf(testPrueba) && len(err) == 1 {
		t.Logf("ConvertModificacionToDocumentoPresupuestal test Success, expected models.DocumentoPreupuestal, got %v", response)
	} else {
		t.Errorf("ConvertModificacionToDocumentoPresupuestal Fail, expected models.DocumentoPreupuestal, got %v", response)
		t.Error(err)
	}
}

func TestFormatDataMovimientoExterno(t *testing.T) {
	var testId []int

	testId = append(testId, 1, 2, 3, 4)

	testOutput := fuenteapropiacionhelper.FormatDataMovimientoExterno(testId, map[string]interface{}{"valor": 1},
		map[string]interface{}{"valor": 2},
		map[string]interface{}{"valor": 3},
		map[string]interface{}{"valor": 4})
	if len(testOutput) > 0 {
		t.Logf("TestFormatDataMovimientoExterno Success!!")
	} else {

		t.Errorf("TestFormatDataMovimientoExterno Fail!!")
	}

}

func TestConcatenarFuente(t *testing.T) {
	var testFuenteFinanciamiento models.FuenteFinanciamiento
	var testTipoFuente models.TipoFuenteFinanciamiento
	var testApropiacion models.Apropiacion
	var testRubro models.Rubro
	var testEstadoApropiacion models.EstadoApropiacion

	testRubro.Id = 1
	testRubro.Organizacion = 1
	testRubro.Codigo = "test"
	testRubro.Descripcion = "test"
	testRubro.UnidadEjecutora = "test"
	testRubro.Nombre = "test"

	testEstadoApropiacion.Id = 1
	testEstadoApropiacion.Nombre = "test"
	testEstadoApropiacion.Descripcion = "test"

	testApropiacion.Id = 1
	testApropiacion.Vigencia = 2019
	testApropiacion.UnidadEjecutora = 1
	testApropiacion.Valor = 1
	testApropiacion.RubroId = &testRubro
	testApropiacion.EstadoApropiacionId = &testEstadoApropiacion

	testTipoFuente.Id = 1
	testTipoFuente.Nombre = "test"
	testTipoFuente.Descripcion = "test"

	testFuenteFinanciamiento.Id = 1
	testFuenteFinanciamiento.Nombre = "test"
	testFuenteFinanciamiento.Descripcion = "test"
	testFuenteFinanciamiento.Codigo = "test"
	testFuenteFinanciamiento.TipoFuenteFinanciamiento = &testTipoFuente

	testOutput := fuenteapropiacionhelper.ConcatenarFuente(&testFuenteFinanciamiento,
		map[string]interface{}{"Apropiacion": testApropiacion})

	if len(testOutput) > 0 {
		t.Log("TestConcatenarFuente Success!!")
	} else {
		t.Error("TestConcatenarFuente Fail!!")
		t.Fail()
	}
}
