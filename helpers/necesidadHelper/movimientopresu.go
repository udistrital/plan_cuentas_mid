package necesidadhelper

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientos "github.com/udistrital/movimientos_crud/models"
	necesidad_models "github.com/udistrital/necesidades_crud/models"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mid/helpers/utils"
)

// InterceptorMovimientoNecesidad, toma la necesidad y su id para determinar si se debe solo actualizar o si es para aprobar y en tal caso
// hacer un movimiento con esta informacion de la necesidad
func InterceptorMovimientoNecesidad(id int, necesidadent necesidad_models.Necesidad) (necesidad necesidad_models.Necesidad, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "InterceptorMovimientoNecesidad - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	if necesidadent.EstadoNecesidadId.Nombre != "Aprobada" {
		if v, err := PutNecesidadService(id, necesidadent); err != nil {
			outputError = err
		} else {
			necesidad = v
		}
	} else {
		if resp, err := EvaluarMovimiento(necesidadent); err != nil {
			outputError = err
		} else if resp {
			if err := RealizarMovimiento(necesidadent); err != nil {
				outputError = err
			} else {
				if v, err := PutNecesidadService(id, necesidadent); err != nil {
					outputError = err
				} else {
					necesidad = v
				}
			}
		}
	}
	return
}

//RealizarMovimiento, toma la informacion de la necesidad para poder generar y estructurar el movimiento
func RealizarMovimiento(necesidad necesidad_models.Necesidad) (outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "RealizarMovimiento - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	var mov []models_movimientos.CuentasMovimientoProcesoExterno
	var mov1 models_movimientos.CuentasMovimientoProcesoExterno
	var movext models_movimientos.MovimientoProcesoExterno
	var tipomov models_movimientos.TipoMovimiento
	if response, err := GetTrNecesidad(strconv.Itoa(necesidad.Id)); err != nil {
		outputError = err
	} else {
		tipomov.Id, _ = strconv.Atoi(beego.AppConfig.String("tipomovimiento"))
		movext.TipoMovimientoId = &tipomov
		movext.Activo = true
		movext.MovimientoProcesoExterno = 0
		movext.ProcesoExterno = 0
		movext.Detalle, _ = utils.Serializar(map[string]interface{}{
			"Estado":      "Publicado",
			"NecesidadId": fmt.Sprintf("%v", necesidad.Id),
		})
		if movimientoext, err := movimientohelper.CrearMovimientoProcesoExt(movext); err != nil {
			outputError = err
		} else {
			if necesidad.TipoFinanciacionNecesidadId.Nombre == "Inversion" {
				for _, valor := range response.Rubros {
					for _, meta := range valor.Metas {
						for _, actividadp := range meta.Actividades {
							actividad := *actividadp
							fuentesi := actividad["FuentesActividad"]
							fuentesp := fuentesi.([]*map[string]interface{})
							for _, fuentep := range fuentesp {
								fuente := *fuentep
								mov1.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
									"RubroId":                valor.RubroId,
									"ActividadId":            fmt.Sprintf("%v", int(actividad["Id"].(float64))),
									"FuenteFinanciamientoId": fmt.Sprintf("%v", int(fuente["Id"].(float64))),
								})
								mov1.Mov_Proc_Ext = strconv.Itoa(movimientoext.Id)
								mov1.Valor = -fuente["MontoParcial"].(float64)
								mov = append(mov, mov1)
								if _, err := movimientohelper.CrearMovimiento(mov); err != nil {
									outputError = err
								}
								mov = nil
							}
						}
					}
				}
			} else if necesidad.TipoFinanciacionNecesidadId.Nombre == "Funcionamiento" {
				for _, valor := range response.Rubros {
					for _, fuentep := range valor.Fuentes {
						fuente := *fuentep
						mov1.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
							"RubroId":                valor.RubroId,
							"FuenteFinanciamientoId": fuente["Id"].(string),
						})
						mov1.Mov_Proc_Ext = strconv.Itoa(movimientoext.Id)
						mov1.Valor = -fuente["MontoParcial"].(float64)
						mov = append(mov, mov1)
						if movimientodet, err := movimientohelper.CrearMovimiento(mov); err != nil {
							outputError = err
						} else {
							logs.Info(movimientodet)
						}
						mov = nil
					}
				}
			}
		}
	}
	return
}

// EvaluarMovimiento, A partir de la necesidad se determina si hay fondos para crubir la necesidad evaluando si es por inversion o funcionamiento
func EvaluarMovimiento(necesidad necesidad_models.Necesidad) (resultado bool, outputError map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			outputError = map[string]interface{}{
				"funcion": "EvaluarMovimiento - Unhandled Error!",
				"err":     err,
				"status":  "500",
			}
			panic(outputError)
		}
	}()
	resultado = false
	var mov []models_movimientos.CuentasMovimientoProcesoExterno
	var mov1 models_movimientos.CuentasMovimientoProcesoExterno
	if response, err := GetTrNecesidad(strconv.Itoa(necesidad.Id)); err != nil {
		outputError = err
	} else {
		if necesidad.TipoFinanciacionNecesidadId.Nombre == "Inversion" {
			for _, valor := range response.Rubros {
				for _, meta := range valor.Metas {
					for _, actividadp := range meta.Actividades {
						actividad := *actividadp
						fuentesi := actividad["FuentesActividad"]
						fuentesp := fuentesi.([]*map[string]interface{})
						for _, fuentep := range fuentesp {
							fuente := *fuentep
							mov1.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
								"RubroId":                valor.RubroId,
								"ActividadId":            fmt.Sprintf("%v", int(actividad["Id"].(float64))),
								"FuenteFinanciamientoId": fmt.Sprintf("%v", int(fuente["Id"].(float64))),
							})
							mov = append(mov, mov1)
							if movimientos, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
								outputError = err
							} else {
								for _, movimiento := range movimientos {
									if movimiento.Saldo > fuente["MontoParcial"].(float64) {
										resultado = true
									}
								}
							}
							mov = nil
						}
					}
				}
			}
		} else if necesidad.TipoFinanciacionNecesidadId.Nombre == "Funcionamiento" {
			for _, valor := range response.Rubros {
				for _, fuentep := range valor.Fuentes {
					fuente := *fuentep
					mov1.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
						"RubroId":                valor.RubroId,
						"FuenteFinanciamientoId": fuente["Id"].(string),
					})
					mov = append(mov, mov1)
					if movimientos, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
						outputError = err
					} else {
						for _, movimiento := range movimientos {
							if movimiento.Saldo > fuente["MontoParcial"].(float64) {
								resultado = true
							}
						}
					}
					mov = nil
				}
			}
		}
	}
	return
}
