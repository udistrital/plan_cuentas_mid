package compositor

import (
	"github.com/astaxie/beego/logs"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// AddMovimientoTransaction ... perform the transaction between mongo and postgres services for
// movimiento's data registration.
func AddMovimientoTransaction(data ...models.Movimiento) (err error) {

	// Send Data to CRUD
	response, err := movimientomanager.AddMovimientoAPICrud(data...)

	if err != nil {
		return
	}

	crudIDs := response.Body.(map[string]interface{})
	intArr := crudIDs["Ids"].([]interface{})
	for i := 0; i < len(intArr); i++ {
		data[i].Id = int(intArr[i].(float64))
	}
	mongoData := movimientohelper.FormatDataForMovimientosMongoAPI(data...)
	logs.Debug(mongoData)
	// Send Data to Mongo
	_, err = movimientomanager.AddMovimientoAPIMongo(mongoData...)

	return
}
