package test

import (
	fuenteapropiacionhelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteApropiacionHelper"
	"github.com/udistrital/plan_cuentas_mid/models"
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
	t.Log(reflect.TypeOf(testOutput))
	//if reflect.TypeOf(testData) == ""
}
