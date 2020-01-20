package test

import (
	fuenteapropiacionhelper "github.com/udistrital/plan_cuentas_mid/helpers/fuenteApropiacionHelper"
	"testing"
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
