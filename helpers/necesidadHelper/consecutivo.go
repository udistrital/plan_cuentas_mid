package necesidadhelper

import (
	models_configuracion "github.com/udistrital/configuracion_api/models"
	models_consecutivos "github.com/udistrital/consecutivos_crud/models"
	"github.com/udistrital/plan_cuentas_mid/helpers/crud/configuracion"
	"github.com/udistrital/plan_cuentas_mid/helpers/crud/consecutivo"
)

func CrearConsecutivoNecesidad(vigencia int) (id int, outputError map[string]interface{}) {
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

	return
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
	if proceso, err := configuracion.ObtenerProceso("nc"); err != nil {
		outputError = err
	} else {
		if proceso[0].Id != 0 {
			id = int(proceso[0].Id)
		} else {
			if Aplicacion, err := configuracion.ObtenerIdAplicacionNecesidades(); err != nil {
				outputError = err
			} else {
				var procesomodel models_configuracion.Proceso
				procesomodel.Activo = true
				procesomodel.Nombre = "Necesidades"
				procesomodel.Sigla = "nc"
				procesomodel.AplicacionId.Id = Aplicacion.Id
				procesomodel.Descripcion = "solicitudes de necesidades Kronos"
				if proaux, err := configuracion.CrearProceso(procesomodel); err != nil {
					outputError = err
				} else {
					id = int(proaux.Id)
				}
			}
		}
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
	modelconsecutivo.Descripcion = "Necesidad"
	if consecutivo, err := consecutivo.GenerarConsecutivo(modelconsecutivo); err != nil {
		outputError = err
	} else {
		respconsecutivo = consecutivo
	}
	return
}
