package compositor

import (
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	movimientomanager "github.com/udistrital/plan_cuentas_mid/managers/movimientoManager"
	"github.com/udistrital/plan_cuentas_mid/models"
)

// AddMovimientoTransaction ... perform the transaction between mongo and postgres services for
// movimiento's data registration.
func AddMovimientoTransaction(data models.Movimiento) (err error) {
	// Format Data for Movimientos API interaction.
	if movimeintosAPIData, err := movimientohelper.FormatDataForMovimientosAPI(data); err != nil {

		// Send Data to CRUD
		if err = movimientomanager.AddMovimientoAPICrud(movimeintosAPIData); err == nil {

			// Send Data to Mongo
			if err = movimientomanager.AddMovimientoAPIMongo(movimeintosAPIData); err != nil {

				// Rollback Data From CRUD if err isn't null.
				err = movimientomanager.DeleteMovimientoAPICrud(movimeintosAPIData["Id"].(int))

			}

		}

	}

	return err
}
