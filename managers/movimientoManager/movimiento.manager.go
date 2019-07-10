package movimientomanager

import (
	"github.com/astaxie/beego/logs"
)

// AddMovimientoAPICrud ... Send movimiento data to crud for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPICrud(data map[string]interface{}) error {
	logs.Debug("Add Movimiento CRUD")
	data["Id"] = 100
	return nil
}

// AddMovimientoAPIMongo ... Send movimiento data to mongo for its registration.
// Returns in data["Id"] the result id for the operation.
func AddMovimientoAPIMongo(data map[string]interface{}) error {
	logs.Debug("Add movimiento Mongo")
	data["Id"] = 100
	return nil
}

// DeleteMovimientoAPICrud ... Delete movimiento data in Movimiento API CRUD By Id.
func DeleteMovimientoAPICrud(id int) error {
	logs.Debug("Delete Movimiento From CRUD", id)
	return nil
}
