package necesidadhelper

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// Variables para obtener la información de una fuente desde mongo
var (
	vigencia        string
	unidadEjecutora string
)

//funciones relacionadas a get de necesidad

// GetTrNecesidad obtiene necesidad de crud apí con sus objetos relacionados
func GetTrNecesidad(id string) (trnecesidad models.TrNecesidad, outErr map[string]interface{}) {
	var err map[string]interface{}

	if trnecesidad.Necesidad, err = getNecesidadCrud(id); err != nil {
		return trnecesidad, err
	}
	nec := trnecesidad.Necesidad
	vig := (*nec)["Vigencia"].(string)
	af := fmt.Sprintf("%.0f", (*nec)["AreaFuncional"].(float64))
	if trnecesidad.DetalleServicioNecesidad, err = getDetalleServicio(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.DetallePrestacionServicioNecesidad, err = getDetallePrestacionServicio(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.ProductosCatalogoNecesidad, err = getProductosCatalogo(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.MarcoLegalNecesidad, err = getMarcoLegal(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.ActividadEconomicaNecesidad, err = getActividadEconomica(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.ActividadEspecificaNecesidad, err = getActividadEspecifica(id); err != nil {
		return trnecesidad, err
	}
	if trnecesidad.RequisitosMinimos, err = getRequisitosNecesidad(id); err != nil {
		return trnecesidad, err
	}
	vigencia = vig
	unidadEjecutora = af
	if trnecesidad.Rubros, err = getRubrosNecesidad(id, vig, af); err != nil {
		return trnecesidad, err
	}
	return trnecesidad, nil

}

// traer la necesidad
func getNecesidadCrud(necesidadId string) (necesidad *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad/?query=Id:" + necesidadId
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getNecesidadCrud", "Error": err.Error()}
		return nil, outErr
	}
	if len(res) == 1 {
		necesidad = &res[0]
	} else {
		err := fmt.Errorf("no existe necesidad con id:%s", necesidadId)
		logs.Error(err)
		outErr = map[string]interface{}{"Function": "getNecesidadCrud", "Error": err.Error()}
	}
	return
}

// TODO se pueden generalizar las funciones que traen valores y arreglos con un closure para reducir lineas de codigo
// traer detalle servicio asociado a la necesidad
func getDetalleServicio(necesidadId string) (ds *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_servicio_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetalleServicio", "Error": err.Error()}
		return nil, outErr
	}
	if len(res) == 1 {
		res[0]["NecesidadId"] = nil
		ds = &res[0]
	}
	return
}

// traer detalle prestacion servicio asociado a la necesidad
func getDetallePrestacionServicio(necesidadId string) (dps *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_prestacion_servicio/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetallePrestacionServicio", "Error": err.Error()}
		return nil, outErr
	}
	if len(res) == 1 {
		res[0]["NecesidadId"] = nil
		dps = &res[0]
	}
	return
}

// traer productos catalogo asociados a la necesidad
func getProductosCatalogo(necesidadId string) (productos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_catalogo_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	productos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getProductosCatalogo", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["NecesidadId"] = nil
		if value["RequisitosMinimos"], outErr = getRequisitosProducto(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		} else {
			productos = append(productos, &value)
		}
	}
	return productos, nil
}

// traer req minimos de un producto catalogo
func getRequisitosProducto(productoCatalogoNecesidadId string) (requisitos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo/?query=ProductoCatalogoNecesidadId:"
	urlcrud += productoCatalogoNecesidadId
	var res []map[string]interface{}
	requisitos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosProducto", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["ProductoCatalogoNecesidadId"] = nil
		requisitos = append(requisitos, &value)
	}
	return requisitos, nil
}

// get marco legal de la necesidad
func getMarcoLegal(necesidadId string) (ml []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "marco_legal_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	ml = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getMarcoLegal", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["NecesidadId"] = nil
		ml = append(ml, &value)
	}
	return ml, nil
}

// get actividad especifica de la necesidad
func getActividadEspecifica(necesidadId string) (ae []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_especifica_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	ae = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEspecifica", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["NecesidadId"] = nil
		ae = append(ae, &value)
	}
	return ae, nil
}

// get requisitos minimos de la necesidad
func getRequisitosNecesidad(necesidadId string) (rm []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	rm = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosNecesidad", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["NecesidadId"] = nil
		rm = append(rm, &value)
	}
	return rm, nil
}

// get actividad economica de la necesidad
func getActividadEconomica(necesidadId string) (aec []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_economica_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	aec = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEconomica", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["NecesidadId"] = nil
		aec = append(aec, &value)
	}
	return aec, nil
}

// get rubros de la necesidad
func getRubrosNecesidad(necesidadId string, vigencia string, areafuncional string) (rubros []*models.RubroNecesidad, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "rubro_necesidad/?query=NecesidadId:" + necesidadId
	var res []map[string]interface{}
	rubros = make([]*models.RubroNecesidad, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {

		value["NecesidadId"] = nil
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiacion/" + value["RubroId"].(string) + "/" + vigencia + "/" + areafuncional
		var resMongo map[string]interface{}
		if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
			value["InfoRubro"] = resMongo["Body"]
		}
		var rubro models.RubroNecesidad
		if err := formatdata.FillStruct(value, &rubro); err != nil {
			outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
			return nil, outErr
		}
		if rubro.Fuentes, outErr = getFuentesRubro(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		}
		if rubro.Productos, outErr = getProductosRubro(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		}
		if rubro.Metas, outErr = getMetasRubro(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		} else {
			rubros = append(rubros, &rubro)
		}
	}
	return rubros, nil
}

// get fuentes rubro
func getFuentesRubro(rubroNecesidadId string) (fuentes []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_rubro_necesidad/?query=RubroNecesidadId:" + rubroNecesidadId
	var res []map[string]interface{}
	fuentes = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getFuentesRubro", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		var resMongo map[string]interface{}
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento/" + value["FuenteId"].(string) + "/" + vigencia + "/" + unidadEjecutora
		if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
			value["InfoFuente"] = resMongo["Body"]
		}
		value["RubroNecesidadId"] = nil
		fuentes = append(fuentes, &value)
	}
	return fuentes, nil
}

// get productos rubro
func getProductosRubro(RubroNecesidadId string) (productos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_rubro_necesidad/?query=RubroNecesidadId:" + RubroNecesidadId
	var res []map[string]interface{}
	productos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getProductosRubro", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		var resMongo map[string]interface{}
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "producto/" + value["ProductoId"].(string)
		if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
			value["InfoProducto"] = resMongo["Body"]
		}
		value["RubroNecesidadId"] = nil
		productos = append(productos, &value)
	}
	return productos, nil
}

// get metas rubro
func getMetasRubro(RubroNecesidadId string) (metas []*models.MetaRubroNecesidad, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "meta_rubro_necesidad/?query=RubroNecesidadId:" + RubroNecesidadId
	var res []map[string]interface{}
	metas = make([]*models.MetaRubroNecesidad, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getMetasRubro", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["RubroNecesidadId"] = nil
		var meta models.MetaRubroNecesidad
		if err := formatdata.FillStruct(value, &meta); err != nil {
			outErr = map[string]interface{}{"Function": "getMetasRubro", "Error": err.Error()}
			return nil, outErr
		}
		if meta.Actividades, outErr = getActividadesMeta(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		} else {
			metas = append(metas, &meta)
		}
	}
	return metas, nil

}

// get actividades meta
func getActividadesMeta(MetaRubroNecesidadId string) (actividades []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_meta/?query=MetaRubroNecesidadId:" + MetaRubroNecesidadId
	var res []map[string]interface{}
	actividades = make([]*map[string]interface{}, 0)
	// logs.Debug(urlcrud)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadesMeta", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		value["MetaRubroNecesidadId"] = nil
		if value["FuentesActividad"], outErr = getFuentesActividad(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
			return nil, outErr
		} else {
			actividades = append(actividades, &value)

		}
	}

	return actividades, nil
}

// getFuentesActividad
func getFuentesActividad(ActividadMetaNecesidadId string) (fuentes []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_actividad/?query=ActividadMetaNecesidadId:" + ActividadMetaNecesidadId
	var res []map[string]interface{}
	fuentes = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosProducto", "Error": err.Error()}
		return nil, outErr
	}
	for _, value := range res {
		var resMongo map[string]interface{}
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento/" + value["FuenteId"].(string) + "/" + vigencia + "/" + unidadEjecutora
		if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
			value["InfoFuente"] = resMongo["Body"]
		}
		value["ActividadMetaNecesidadId"] = nil
		fuentes = append(fuentes, &value)
	}
	return fuentes, nil
}

// fin funciones get necesidad

//funciones post necesidad

// PostTrNecesidad insertar la necesidad completa
func PostTrNecesidad(trnecesidad models.TrNecesidad) (out models.TrNecesidad, outErr map[string]interface{}) {
	var (
		resDependencia map[string]interface{}
	)
	errOut := make(map[string]interface{})
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "dependencia_necesidad/"
	if err := request.SendJson(urlcrud, "POST", &resDependencia, (*trnecesidad.Necesidad)["DependenciaNecesidadId"]); err == nil {
		(*trnecesidad.Necesidad)["DependenciaNecesidadId"].(map[string]interface{})["Id"] = resDependencia["Id"]
		urlcrud = beego.AppConfig.String("necesidadesCrudService") + "necesidad/"

		switch (*trnecesidad.Necesidad)["EstadoNecesidadId"].(map[string]interface{})["CodigoAbreviacionn"] {
		case "G": // Necesidad guardada
			_, existe := (*trnecesidad.Necesidad)["ConsecutivoSolicitud"]
			if !existe {
				(*trnecesidad.Necesidad)["ConsecutivoSolicitud"] = agregarConsecutivoSolicitiud()
			}

		case "A": // Necesidad aprobada
			(*trnecesidad.Necesidad)["ConsecutivoNecesidad"] = agregarConsecutivoNecesidad()
		}

		if err = request.SendJson(urlcrud, "POST", &out.Necesidad, trnecesidad.Necesidad); err == nil {
			// TODO: Hacer que esto sea transaccional, es decir, que
			//       si alguno de los POST falla, reintentar hasta
			//       que sea exitoso.
			//       Posiblemente se pueda hacer con WSO2
			var err map[string]interface{}
			if out.DetalleServicioNecesidad, err = postDetalleServicio(trnecesidad.DetalleServicioNecesidad, out.Necesidad); err != nil {
				errOut["postDetalleServicio"] = err
			}
			if out.DetallePrestacionServicioNecesidad, err = postDetallePrestacionServicio(trnecesidad.DetallePrestacionServicioNecesidad, out.Necesidad); err != nil {
				errOut["postDetallePrestacionServicio"] = err
			}
			if out.ProductosCatalogoNecesidad, err = postProductosCatalogo(trnecesidad.ProductosCatalogoNecesidad, out.Necesidad); err != nil {
				errOut["postProductosCatalogo"] = err
			}
			if out.MarcoLegalNecesidad, err = postMarcoLegal(trnecesidad.MarcoLegalNecesidad, out.Necesidad); err != nil {
				errOut["postMarcoLegal"] = err
			}
			if out.ActividadEconomicaNecesidad, err = postActividadesEconomicas(trnecesidad.ActividadEconomicaNecesidad, out.Necesidad); err != nil {
				errOut["postActividadesEconomicas"] = err
			}
			if out.ActividadEspecificaNecesidad, err = postActividadesEspecificas(trnecesidad.ActividadEspecificaNecesidad, out.Necesidad); err != nil {
				errOut["postActividadesEspecificas"] = err
			}
			if out.RequisitosMinimos, err = postRequisitosNecesidad(trnecesidad.RequisitosMinimos, out.Necesidad); err != nil {
				errOut["postRequisitosNecesidad"] = err
			}
			if out.Rubros, err = postRubros(trnecesidad.Rubros, out.Necesidad); err != nil {
				errOut["postRubros"] = err
			}

			if len(errOut) > 0 {
				return out, map[string]interface{}{"Function": "PostTrNecesidad", "Error": errOut}
			}

		} else {
			return out, map[string]interface{}{"Function": "PostTrNecesidad", "Error": err.Error()}
		}
	} else {
		outErr = map[string]interface{}{"Function": "PostTrNecesidad", "Error": err.Error()}
		return out, outErr
	}
	return out, nil

}

// agregarConsecutivoSolicitiud calcula el consecutivo sumando todas las necesitades existentes hasta el momento
//
// Deprecated: Refactorizar para usar consecutivos_crud !!
func agregarConsecutivoSolicitiud() int {
	var necesidades []map[string]interface{}
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad?limit=-1"
	urlcrud += "&fields=Id"
	if err := request.GetJson(urlcrud, &necesidades); err != nil {
		return 0
	}
	return len(necesidades) + 1
}

// agregarConsecutivoNecesidad calcula el consecutivo sumando todas las necesidades existenes hasta el momento que estén
// en estado: aprobada, rechazada, anulada, modificada, enviada y cdp solicitado
func agregarConsecutivoNecesidad() int {
	var necesidades []interface{}
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad?limit=-1&query=" +
		"EstadoNecesidadId.CodigoAbreviacionn:A," + // Aprobada
		"EstadoNecesidadId.CodigoAbreviacionn:R," + // Rechazada
		"EstadoNecesidadId.CodigoAbreviacionn:AN," + // Anulada
		"EstadoNecesidadId.CodigoAbreviacionn:M," + // Modificada
		"EstadoNecesidadId.CodigoAbreviacionn:E," + // Enviada
		"EstadoNecesidadId.CodigoAbreviacionn:CS" // CDP Solicitado
	urlcrud += "&fields=Id"
	if err := request.GetJson(urlcrud, &necesidades); err != nil {
		return 0
	}

	return len(necesidades) + 1
}

// post detalle servicio necesidad
func postDetalleServicio(detalle *map[string]interface{}, necesidad *map[string]interface{}) (out *map[string]interface{}, outErr map[string]interface{}) {
	if detalle == nil || len(*detalle) == 0 {
		return nil, nil
	}
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_servicio_necesidad/"
	(*detalle)["NecesidadId"] = necesidad
	if err := request.SendJson(urlcrud, "POST", &out, detalle); err != nil {
		return nil, map[string]interface{}{"Function": "postDetalleServicio", "Error": err.Error()}
	}
	return out, nil
}

// post detalle servicio necesidad
func postDetallePrestacionServicio(detalle *map[string]interface{}, necesidad *map[string]interface{}) (out *map[string]interface{}, outErr map[string]interface{}) {
	if detalle == nil || len(*detalle) == 0 {
		return nil, nil
	}
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_prestacion_servicio/"
	(*detalle)["NecesidadId"] = necesidad
	if err := request.SendJson(urlcrud, "POST", &out, detalle); err != nil {
		return nil, map[string]interface{}{"Function": "postDetallePrestacionServicio", "Error": err.Error()}
	}
	return out, nil

}

// post productos catalogo necesidad
func postProductosCatalogo(productos []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for k, value := range productos {
		(*productos[k])["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_catalogo_necesidad/"
		var prOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &prOut, productos[k]); err == nil {
			reqmin := (*productos[k])["RequisitosMinimos"].([]interface{})
			urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo/"
			for i := range reqmin {
				reqmin[i].(map[string]interface{})["ProductoCatalogoNecesidadId"] = prOut
				var reqOut map[string]interface{}
				if err = request.SendJson(urlcrud, "POST", &reqOut, reqmin[i]); err == nil {
					(*productos[k])["RequisitosMinimos"] = append((*productos[k])["RequisitosMinimos"].([]interface{}), reqOut)
				} else {
					return nil, map[string]interface{}{"Function": "postProductosCatalogo", "Error": err.Error()}
				}

			}
		} else {
			return nil, map[string]interface{}{"Function": "postProductosCatalogo", "Error": err.Error()}
		}
		(*value)["NecesidadId"] = nil
		out = append(out, &prOut)
	}
	return
}

// post marco legal
func postMarcoLegal(marcolegal []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range marcolegal {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "marco_legal_necesidad/"
		var mlOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &mlOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postMarcoLegal", "Error": err.Error()}
		}
		mlOut["NecesidadId"] = nil
		out = append(out, &mlOut)
	}
	return
}

// post actividad especifica
func postActividadesEspecificas(ae []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range ae {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_especifica_necesidad/"
		var aeOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &aeOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postActividadesEspecificas", "Error": err.Error()}
		}
		aeOut["NecesidadId"] = nil
		out = append(out, &aeOut)
	}
	return
}

// post requisitos minimos necesidad
func postRequisitosNecesidad(rm []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range rm {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo_necesidad/"
		var rmOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &rmOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postpostRequisitosNecesidad", "Error": err.Error()}
		}
		rmOut["NecesidadId"] = nil
		out = append(out, &rmOut)
	}
	return
}

// post actividad economica

func postActividadesEconomicas(ae []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range ae {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_economica_necesidad/"
		var aeOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &aeOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postActividadesEspecificas", "Error": err.Error()}
		}
		aeOut["NecesidadId"] = nil
		out = append(out, &aeOut)
	}
	return
}

// post rubro
func postRubros(rubros []*models.RubroNecesidad, necesidad *map[string]interface{}) (out []*models.RubroNecesidad, outErr map[string]interface{}) {
	var errOut map[string]interface{}
	for _, rubro := range rubros {
		rubro.NecesidadId = *necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "rubro_necesidad/"
		var rOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &rOut, rubro); err != nil {
			return nil, map[string]interface{}{"Function": "postRubros", "Error": err.Error()}
		}
		rOut["NecesidadId"] = nil
		if rOut["Fuentes"], errOut = postFuentesRubro(rubro.Fuentes, &rOut); errOut == nil {
			logs.Warn(errOut)
			// TODO: Lazo vacío, no hace nada...
			// Revisar si se debe retornar o concatenar el error
		}
		if rOut["Productos"], errOut = postProductosRubro(rubro.Productos, &rOut); errOut == nil {
			logs.Warn(errOut)
			// TODO: Lazo vacío, no hace nada...
			// Revisar si se debe retornar o concatenar el error
		}
		if rOut["Metas"], errOut = postMetasRubro(rubro.Metas, &rOut); errOut != nil {
			return nil, map[string]interface{}{"Function": "postRubros", "Error": errOut}
		}
		var rubroOut models.RubroNecesidad
		if errConvert := formatdata.FillStruct(rOut, &rubroOut); errConvert == nil {
			out = append(out, &rubroOut)
		}

	}
	return
}

// post fuentes

func postFuentesRubro(f []*map[string]interface{}, rubro *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range f {
		(*value)["RubroNecesidadId"] = rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_rubro_necesidad/"
		var fOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &fOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postFuentesRubro", "Error": err.Error()}
		}
		fOut["RubroNecesidadId"] = nil
		out = append(out, &fOut)
	}
	return
}

// post productos

func postProductosRubro(p []*map[string]interface{}, rubro *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range p {
		(*value)["RubroNecesidadId"] = rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_rubro_necesidad/"
		var pOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &pOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postProductosRubro", "Error": err.Error()}
		}
		pOut["RubroNecesidadId"] = nil
		out = append(out, &pOut)
	}
	return
}

// post metas

func postMetasRubro(m []*models.MetaRubroNecesidad, rubro *map[string]interface{}) (out []*models.MetaRubroNecesidad, outErr map[string]interface{}) {
	for _, meta := range m {
		meta.RubroNecesidadId = *rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "meta_rubro_necesidad/"
		var mOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &mOut, meta); err != nil {
			return nil, map[string]interface{}{"Function": "postMetasRubro", "Error": err.Error()}
		}
		mOut["RubroNecesidadId"] = nil
		if mOut["Actividades"], outErr = postActividadesMeta(meta.Actividades, &mOut); outErr != nil {
			logs.Warn(outErr)
			// TODO: Evaluar si sería buena idea retornar acá
			//       En caso afirmativo, eliminar el logs.Warn
			//       anterior y este comentario, y
			//       descomentar el siguiente return:
			// return
		}
		var metaOut models.MetaRubroNecesidad
		if errConvert := formatdata.FillStruct(mOut, &metaOut); errConvert == nil {
			out = append(out, &metaOut)
		}

	}
	return
}

// post actividades
func postActividadesMeta(act []*map[string]interface{}, meta *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range act {
		(*value)["MetaRubroNecesidadId"] = meta
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_meta/"
		var actOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &actOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postActividadesMeta", "Error": err.Error()}
		}
		actOut["MetaRubroNecesidadId"] = nil
		postFuentesActividad((*value)["FuentesActividad"].([]interface{}), &actOut)
		out = append(out, &actOut)
	}
	return
}

// post fuentes actividades

func postFuentesActividad(fuentes []interface{}, act *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	for _, value := range fuentes {
		value.(map[string]interface{})["ActividadMetaNecesidadId"] = act
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_actividad/"
		var fOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &fOut, value); err != nil {
			return nil, map[string]interface{}{"Function": "postFuentesActividad", "Error": err.Error()}
		}
		fOut["ActividadMetaNecesidadId"] = nil
		out = append(out, &fOut)
	}
	return
}
