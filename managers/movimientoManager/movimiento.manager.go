package movimientomanager

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// AddMovimientoAPICrud ... Send movimiento data to crud for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPICrud(data ...models.Movimiento) (responseformat.Response, error) {
	var response responseformat.Response

	err := request.SendJson(beego.AppConfig.String("movimientosCrudService")+"movimiento_detalle/registrar_multiple", "POST", &response, data)
	if responseformat.CheckResponseError(response) {
		err = errors.New("Movimientos API Error")
		logs.Error(err.Error())
	}

	return response, err
}

// AddMovimientoAPIMongo ... Send movimiento data to mongo for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPIMongo(data ...models.MovimientoMongo) (responseformat.Response, error) {
	var response responseformat.Response

	err := request.SendJson(beego.AppConfig.String("financieraMongoCurdApiService")+"movimiento/RegistrarMovimientos", "POST", &response, data)
	if responseformat.CheckResponseError(response) {
		err = errors.New("Mongo API Error")
		logs.Error(err.Error())
	}

	return response, err
}

// DeleteMovimientoAPICrud ... Delete movimiento data in Movimiento API CRUD By Id.
func DeleteMovimientoAPICrud(id int) error {
	logs.Debug("Delete Movimiento From CRUD", id)
	return nil
}
