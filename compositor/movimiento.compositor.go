package compositor

import (
	"reflect"

	"github.com/astaxie/beego/logs"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// AddMovimientoTransaction ... perform the transaction between mongo and postgres services for
// movimiento's data registration.
func AddMovimientoTransaction(data ...models.Movimiento) (err error) {

	// Send Data to CRUD
	if response, err := movimientomanager.AddMovimientoAPICrud(data...); err == nil {
		crudIDs := response.Body.(map[string]interface{})
		logs.Debug(reflect.TypeOf(crudIDs["Ids"]))
		intArr := crudIDs["Ids"].([]interface{})
		for i := 0; i < len(intArr); i++ {
			data[i].Id = int(intArr[i].(float64))
		}
		mongoData := movimientohelper.FormatDataForMovimientosMongoAPI(data...)
		// Send Data to Mongo
		_, err = movimientomanager.AddMovimientoAPIMongo(mongoData...)

	}
	return err
}
