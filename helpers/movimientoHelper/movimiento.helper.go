package movimientohelper

import (
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
)

// FormatDataForMovimientosAPI ... format Movimiento Data to MoimientosAPI data structure.
func FormatDataForMovimientosAPI(data models.Movimiento) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := formatdata.FillStruct(data, &res)

	return res, err
}
