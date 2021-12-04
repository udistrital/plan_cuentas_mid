package necesidadhelper

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	models_movimientos "github.com/udistrital/movimientos_crud/models"
	necesidad_models "github.com/udistrital/necesidades_crud/models"
	movimientohelper "github.com/udistrital/plan_cuentas_mid/helpers/movimientoHelper"
	"github.com/udistrital/plan_cuentas_mid/helpers/utils"
)

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

	if necesidad.EstadoNecesidadId.Nombre != "Aprobada" {
		if v, err := PutNecesidadService(id, necesidadent); err != nil {
			outputError = err
		} else {
			necesidad = v
		}
	} else {
		if resp, err := EvaluarMovimiento(necesidad); err != nil {
			outputError = err
		} else if resp {
			if err := RealizarMovimiento(necesidad); err != nil {
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
	//TODO realizar movimiento
	var mov models_movimientos.CuentasMovimientoProcesoExterno
	var movext models_movimientos.MovimientoProcesoExterno
	if response, err := GetTrNecesidad(strconv.Itoa(necesidad.Id)); err != nil {
		outputError = err
	} else {
		movext.TipoMovimientoId.Id, _ = beego.AppConfig.Int("tipomovimiento")
		movext.Activo = true
		movext.MovimientoProcesoExterno = 0
		movext.ProcesoExterno = 0
		movext.Detalle, _ = utils.Serializar(map[string]interface{}{
			"Estado":      "Publicado",
			"NecesidadId": necesidad.Id,
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
							fuentes := fuentesi.([]map[string]interface{})
							for _, fuente := range fuentes {
								mov.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
									"RubroId":                valor.RubroId,
									"FuenteFinanciamientoId": fuente["Id"].(string),
									"ActividadId":            actividad["Id"].(string),
								})
								if movimiento, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
									outputError = err
								} else {
									mov.Mov_Proc_Ext = string(movimientoext.Id)
									mov.Saldo = fuente["MontoParcial"].(float64)
									mov.Valor = movimiento.Saldo + fuente["MontoParcial"].(float64)
									if movimientodet, err := movimientohelper.CrearMovimiento(mov); err != nil {
										outputError = err
									} else {
										logs.Info(movimientodet)
									}
								}
							}
						}
					}
				}
			} else if necesidad.TipoFinanciacionNecesidadId.Nombre == "Funcionamiento" {
				for _, valor := range response.Rubros {
					//Consultar el movimiento de cada rubro
					for _, fuentep := range valor.Fuentes {
						fuente := *fuentep
						mov.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
							"RubroId":                valor.RubroId,
							"FuenteFinanciamientoId": fuente["Id"].(string),
						})
						if movimiento, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
							outputError = err
						} else {
							mov.Mov_Proc_Ext = string(movimiento.MovimientoProcesoExternoId.Id)
							mov.Saldo = fuente["MontoParcial"].(float64)
							mov.Valor = movimiento.Saldo - fuente["MontoParcial"].(float64)
							if movimientodet, err := movimientohelper.CrearMovimiento(mov); err != nil {
								outputError = err
							} else {
								logs.Info(movimientodet)
							}
						}
					}
				}
			}
		}
	}
	return
}

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
	var mov models_movimientos.CuentasMovimientoProcesoExterno
	if response, err := GetTrNecesidad(strconv.Itoa(necesidad.Id)); err != nil {
		resultado = false
		outputError = err
	} else {
		if necesidad.TipoFinanciacionNecesidadId.Nombre == "Inversion" {
			for _, valor := range response.Rubros {
				//Consultar el movimiento de cada rubro
				for _, meta := range valor.Metas {
					for _, actividadp := range meta.Actividades {
						actividad := *actividadp
						fuentesi := actividad["FuentesActividad"]
						fuentes := fuentesi.([]map[string]interface{})
						for _, fuente := range fuentes {
							mov.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
								"RubroId":                valor.RubroId,
								"FuenteFinanciamientoId": fuente["Id"].(string),
								"ActividadId":            actividad["Id"].(string),
							})
							if movimiento, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
								resultado = false
								outputError = err
							} else {
								if movimiento.Saldo > fuente["MontoParcial"].(float64) {
									resultado = true
								} else {
									resultado = false
								}
							}
						}
					}
				}
			}
		} else if necesidad.TipoFinanciacionNecesidadId.Nombre == "Funcionamiento" {
			for _, valor := range response.Rubros {
				//Consultar el movimiento de cada rubro
				for _, fuentep := range valor.Fuentes {
					fuente := *fuentep
					mov.Cuen_Pre, _ = utils.Serializar(map[string]interface{}{
						"RubroId":                valor.RubroId,
						"FuenteFinanciamientoId": fuente["Id"].(string),
					})
					if movimiento, err := movimientohelper.ObtenerUltimoMovimiento(mov); err != nil {
						resultado = false
						outputError = err
					} else {
						if movimiento.Saldo > fuente["MontoParcial"].(float64) {
							resultado = true
						} else {
							resultado = false
						}
					}
				}
			}
		}
	}
	return
}
