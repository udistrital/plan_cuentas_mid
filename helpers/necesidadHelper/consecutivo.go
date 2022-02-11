package necesidadhelper

import (
	models_consecutivos "github.com/udistrital/consecutivos_crud/models"
	consecutivohelper "github.com/udistrital/plan_cuentas_mid/helpers/consecutivoHelper"
)

func CrearConsecutivo(vigencia int) (id int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "CrearConsecutivo - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	if idProceso, err := GetIdProcesoNecesidad(); err != nil {
		outputError = err
	} else {
		if consecutivo, err := SolicitudConsecutivo(vigencia, idProceso); err != nil {
			outputError = err
		} else {
			id = consecutivo.Consecutivo
		}
	}

	return id, nil
}

func GetIdProcesoNecesidad() (id int, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "GetIdProcesoNecesidad - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	if proceso, err := consecutivohelper.ObtenerProcesoNecesidad(); err != nil {
		outputError = err
	} else {
		id = int(proceso.Id)
	}
	return
}

func SolicitudConsecutivo(vigencia int, proceso int) (respconsecutivo models_consecutivos.Consecutivo, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "SolicitudConsecutivo - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()

	var modelconsecutivo models_consecutivos.Consecutivo
	modelconsecutivo.Activo = true
	modelconsecutivo.ContextoId = proceso
	modelconsecutivo.Year = float64(vigencia)

	if consecutivo, err := consecutivohelper.GenerarConsecutivo(modelconsecutivo); err != nil {
		outputError = err
	} else {
		respconsecutivo = consecutivo
	}
	return respconsecutivo, nil
}