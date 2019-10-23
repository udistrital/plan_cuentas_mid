package compositor

import (
	"time"

	"github.com/udistrital/utils_oas/formatdata"

	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// AddMovimientoTransaction ... perform the transaction between mongo and postgres services for
// movimiento's data registration.
func AddMovimientoTransaction(detail interface{}, data models.DocumentoPresupuestal, afectation []models.Movimiento) (finalData interface{}, err error) {
	mapDetail, err := formatdata.ToMap(detail, "bson")

	if err != nil {
		panic(err.Error())
	}

	data.Data = mapDetail
	data.AfectacionMovimiento = afectation
	fechaRegistro := time.Now().Format(time.RFC3339)
	data.FechaRegistro = fechaRegistro
	var idsMovimientos []int
	// Send Data to CRUD
	response, err := movimientomanager.AddMovimientoAPICrud(data.AfectacionMovimiento...)

	if err != nil {
		return
	}

	crudIDs := response.Body.(map[string]interface{})
	intArr := crudIDs["Ids"].([]interface{})
	for i := 0; i < len(intArr); i++ {
		id := int(intArr[i].(float64))
		data.AfectacionMovimiento[i].Id = id
		idsMovimientos = append(idsMovimientos, id)
	}
	data.Afectacion = movimientohelper.FormatDataForMovimientosMongoAPI(data.AfectacionMovimiento...)
	for i := range data.Afectacion {
		data.Afectacion[i].FechaRegistro = fechaRegistro
	}
	// Send Data to Mongo
	response, err = movimientomanager.AddMovimientoAPIMongo(data)
	finalData = response.Body
	if err != nil {
		go movimientomanager.DeleteMovimientoAPICrud(idsMovimientos...)
	}
	return
}
