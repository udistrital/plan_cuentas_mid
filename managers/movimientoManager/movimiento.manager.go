package movimientomanager

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
	"github.com/udistrital/utils_oas/responseformat"
)

// AddMovimientoAPICrud ... Send movimiento data to crud for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPICrud(data ...models.Movimiento) (response responseformat.Response, err error) {

	if err = request.SendJson(beego.AppConfig.String("movimientosCrudService")+"movimiento_detalle/registrar_multiple", "POST", &response, data); err == nil {
		if responseformat.CheckResponseError(response) {
			var errMessage = "Movimientos API Error"
			if messageStr, e := response.Body.(string); e {
				errMessage = errMessage + ": " + messageStr
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		}
	}

	return response, err
}

// AddMovimientoAPIMongo ... Send movimiento data to mongo for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPIMongo(data models.DocumentoPresupuestal) (response responseformat.Response, err error) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("catch", r)
			errStr := fmt.Sprintf("%s", r)
			err = errors.New("Mongo API Error: " + errStr)
		}
	}()
	if err = request.SendJson(beego.AppConfig.String("financieraMongoCurdApiService")+"movimiento/RegistrarMovimientos", "POST", &response, data); err == nil {
		if responseformat.CheckResponseError(response) {
			var errMessage = "Mongo API Error"
			if messageStr, e := response.Body.(string); e {
				errMessage = errMessage + ": " + messageStr
			} else {
				errMessage = errMessage + ": " + fmt.Sprintf("%s", response.Body)
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		}
	}

	return
}

// DeleteMovimientoAPICrud ... Delete movimiento data in Movimiento API CRUD By Id.
func DeleteMovimientoAPICrud(id ...int) (response responseformat.Response, err error) {

	if err = request.SendJson(beego.AppConfig.String("movimientosCrudService")+"movimiento_detalle/eliminar_multiple", "POST", &response, id); err == nil {
		if responseformat.CheckResponseError(response) {
			var errMessage = "Movimientos API Error"
			if messageStr, e := response.Body.(string); e {
				errMessage = errMessage + ": " + messageStr
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		}
	}

	return response, err
}

func SimualteAfectationAPIMongo(cg, vigencia string, data ...models.MovimientoMongo) (response map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			logs.Error("catch", r)
			errStr := fmt.Sprintf("%s", r)
			err = errors.New("Mongo API Error: " + errStr)
		}
	}()
	if err = request.SendJson(beego.AppConfig.String("financieraMongoCurdApiService")+"arbol_rubro_apropiacion/comprobar_balance/"+cg+"/"+vigencia, "POST", &response, data); err == nil {
		//code, _ := strconv.Atoi(response.Code)
		if response["Code"] == 404 {
			var errMessage = "Mongo API Error"
			if messageStr, e := response["Body"].(string); e {
				errMessage = errMessage + ": " + messageStr
			} else {
				errMessage = errMessage + ": " + fmt.Sprintf("%s", response["Body"])
			}
			err = errors.New(errMessage)
			logs.Error(err.Error())
		}
	}

	return
}
