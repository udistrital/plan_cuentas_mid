package necesidadhelper

import (
	"fmt"

	"github.com/astaxie/beego"
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
	if len(*nec) == 0 {
		return trnecesidad, err
	}
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
func getNecesidadCrud(id string) (necesidad *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad/?query=Id:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getNecesidadCrud", "Error": err.Error()}
		return nil, outErr
	}
	necesidad = &res[0]
	return necesidad, nil
}

// TODO se pueden generalizar las funciones que traen valores y arreglos con un closure para reducir lineas de codigo
// traer detalle servicio asociado a la necesidad
func getDetalleServicio(id string) (ds *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_servicio_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetalleServicio", "Error": err.Error()}
		return nil, outErr
	}
	if len(res[0]) > 0 {
		res[0]["NecesidadId"] = nil
	}
	ds = &res[0]
	return ds, nil
}

// traer detalle prestacion servicio asociado a la necesidad
func getDetallePrestacionServicio(id string) (dps *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_prestacion_servicio/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetallePrestacionServicio", "Error": err.Error()}
		return nil, outErr
	}
	if len(res[0]) > 0 {
		res[0]["NecesidadId"] = nil
	}
	dps = &res[0]
	return dps, nil
}

// traer productos catalogo asociados a la necesidad
func getProductosCatalogo(id string) (productos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_catalogo_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	productos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getProductosCatalogo", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			if res[k]["RequisitosMinimos"], outErr = getRequisitosProducto(fmt.Sprintf("%.0f", res[k]["Id"].(float64))); outErr != nil {
				return nil, outErr
			} else {
				productos = append(productos, &res[k])
			}
		}

	}
	return productos, nil
}

// traer req minimos de un producto catalogo
func getRequisitosProducto(id string) (requisitos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo/?query=ProductoCatalogoNecesidadId:" + id
	var res []map[string]interface{}
	requisitos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosProducto", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["ProductoCatalogoNecesidadId"] = nil
			requisitos = append(requisitos, &res[k])
		}
	}
	return requisitos, nil
}

// get marco legal de la necesidad
func getMarcoLegal(id string) (ml []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "marco_legal_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	ml = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getMarcoLegal", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			ml = append(ml, &res[k])
		}

	}
	return ml, nil
}

// get actividad especifica de la necesidad
func getActividadEspecifica(id string) (ae []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_especifica_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	ae = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEspecifica", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			ae = append(ae, &res[k])
		}

	}
	return ae, nil
}

// get requisitos minimos de la necesidad
func getRequisitosNecesidad(id string) (rm []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	rm = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosNecesidad", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			rm = append(rm, &res[k])
		}

	}
	return rm, nil
}

// get actividad economica de la necesidad
func getActividadEconomica(id string) (aec []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_economica_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	aec = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEconomica", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			aec = append(aec, &res[k])
		}

	}
	return aec, nil
}

// get rubros de la necesidad
func getRubrosNecesidad(id string, vigencia string, areafuncional string) (rubros []*models.RubroNecesidad, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "rubro_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	rubros = make([]*models.RubroNecesidad, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["NecesidadId"] = nil
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "arbol_rubro_apropiacion/" + res[k]["RubroId"].(string) + "/" + vigencia + "/" + areafuncional
			var resMongo map[string]interface{}
			if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
				res[k]["InfoRubro"] = resMongo["Body"]
			}
			var rubro models.RubroNecesidad
			if err := formatdata.FillStruct(res[k], &rubro); err != nil {
				outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
				return nil, outErr
			}
			if rubro.Fuentes, outErr = getFuentesRubro(fmt.Sprintf("%.0f", res[k]["Id"].(float64))); outErr != nil {
				return nil, outErr
			}
			if rubro.Productos, outErr = getProductosRubro(fmt.Sprintf("%.0f", res[k]["Id"].(float64))); outErr != nil {
				return nil, outErr
			}
			if rubro.Metas, outErr = getMetasRubro(fmt.Sprintf("%.0f", res[k]["Id"].(float64))); outErr != nil {
				return nil, outErr
			} else {
				rubros = append(rubros, &rubro)
			}
		}

	}
	return rubros, nil
}

// get fuentes rubro
func getFuentesRubro(id string) (fuentes []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_rubro_necesidad/?query=RubroNecesidadId:" + id
	var res []map[string]interface{}
	fuentes = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getFuentesRubro", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			var resMongo map[string]interface{}
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento/" + res[k]["FuenteId"].(string) + "/" + vigencia + "/" + unidadEjecutora
			if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
				res[k]["InfoFuente"] = resMongo["Body"]
			}
			res[k]["RubroNecesidadId"] = nil
			fuentes = append(fuentes, &res[k])
		}
	}
	return fuentes, nil
}

// get productos rubro
func getProductosRubro(id string) (productos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_rubro_necesidad/?query=RubroNecesidadId:" + id
	var res []map[string]interface{}
	productos = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getProductosRubro", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			var resMongo map[string]interface{}
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "producto/" + res[k]["ProductoId"].(string)
			if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
				res[k]["InfoProducto"] = resMongo["Body"]
			}
			res[k]["RubroNecesidadId"] = nil
			productos = append(productos, &res[k])
		}
	}
	return productos, nil
}

// get metas rubro
func getMetasRubro(id string) (metas []*models.MetaRubroNecesidad, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "meta_rubro_necesidad/?query=RubroNecesidadId:" + id
	var res []map[string]interface{}
	metas = make([]*models.MetaRubroNecesidad, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getMetasRubro", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["RubroNecesidadId"] = nil
			var meta models.MetaRubroNecesidad
			if err := formatdata.FillStruct(res[k], &meta); err != nil {
				outErr = map[string]interface{}{"Function": "getMetasRubro", "Error": err.Error()}
				return nil, outErr
			}
			if meta.Actividades, outErr = getActividadesMeta(fmt.Sprintf("%.0f", res[k]["Id"].(float64))); outErr != nil {
				return nil, outErr
			} else {
				metas = append(metas, &meta)
			}
		}
	}
	return metas, nil

}

// get actividades meta
func getActividadesMeta(id string) (actividades []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_meta/?query=MetaRubroNecesidadId:" + id
	var res []map[string]interface{}
	actividades = make([]*map[string]interface{}, 0)
	// logs.Debug(urlcrud)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadesMeta", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			res[k]["MetaRubroNecesidadId"] = nil
			if res[k]["FuentesActividad"], outErr = getFuentesActividad(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
				return nil, outErr
			} else {
				actividades = append(actividades, &res[k])

			}
		}
		fmt.Println("iterac: ", k)

	}

	return actividades, nil
}

// getFuentesActividad
func getFuentesActividad(id string) (fuentes []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_actividad/?query=ActividadMetaNecesidadId:" + id
	var res []map[string]interface{}
	fuentes = make([]*map[string]interface{}, 0)
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosProducto", "Error": err.Error()}
		return nil, outErr
	}
	for k, value := range res {
		if len(value) > 0 {
			var resMongo map[string]interface{}
			urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "fuente_financiamiento/" + value["FuenteId"].(string) + "/" + vigencia + "/" + unidadEjecutora
			if err := request.GetJson(urlmongo, &resMongo); err == nil && len(resMongo) > 0 {
				res[k]["InfoFuente"] = resMongo["Body"]
			}
			res[k]["ActividadMetaNecesidadId"] = nil
			fuentes = append(fuentes, &res[k])
		}
	}
	return fuentes, nil
}

// fin funciones get necesidad

//funciones post necesidad

// PostTrNecesidad insertar la necesidad completa
func PostTrNecesidad(trnecesidad models.TrNecesidad) (out models.TrNecesidad, outErr map[string]interface{}) {
	var (
		resDependencia map[string]interface{}
		errOut         map[string]interface{}
	)
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
			if out.DetalleServicioNecesidad, errOut = postDetalleServicio(trnecesidad.DetalleServicioNecesidad, out.Necesidad); errOut == nil {
			}
			if out.DetallePrestacionServicioNecesidad, errOut = postDetallePrestacionServicio(trnecesidad.DetallePrestacionServicioNecesidad, out.Necesidad); errOut == nil {
			}
			if out.ProductosCatalogoNecesidad, errOut = postProductosCatalogo(trnecesidad.ProductosCatalogoNecesidad, out.Necesidad); errOut == nil {
			}
			if out.MarcoLegalNecesidad, errOut = postMarcoLegal(trnecesidad.MarcoLegalNecesidad, out.Necesidad); errOut == nil {
			}
			if out.ActividadEconomicaNecesidad, errOut = postActividadesEconomicas(trnecesidad.ActividadEconomicaNecesidad, out.Necesidad); errOut == nil {
			}
			if out.ActividadEspecificaNecesidad, errOut = postActividadesEspecificas(trnecesidad.ActividadEspecificaNecesidad, out.Necesidad); errOut == nil {

			}
			if out.RequisitosMinimos, errOut = postRequisitosNecesidad(trnecesidad.RequisitosMinimos, out.Necesidad); errOut == nil {
			}
			if out.Rubros, errOut = postRubros(trnecesidad.Rubros, out.Necesidad); errOut == nil {

			} else {
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
func agregarConsecutivoSolicitiud() int {
	var necesidades []map[string]interface{}
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad?limit=-1"
	if err := request.GetJson(urlcrud, &necesidades); err != nil {
		return 0
	}
	if len(necesidades[0]) == 0 {
		return 1
	}
	//fmt.Println(necesidades[0].(map[string]interface{}), necesidades[0].(map[string]interface{}) == nil)
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
	if productos == nil || len(productos) == 0 {
		return nil, nil
	}
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
	return out, nil

}

// post marco legal
func postMarcoLegal(marcolegal []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if marcolegal == nil || len(marcolegal) == 0 {
		return nil, nil
	}
	for _, value := range marcolegal {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "marco_legal_necesidad/"
		var mlOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &mlOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postMarcoLegal", "Error": err.Error()}
		}
		mlOut["NecesidadId"] = nil
		out = append(out, &mlOut)
	}
	return out, nil

}

// post actividad especifica
func postActividadesEspecificas(ae []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if ae == nil || len(ae) == 0 {
		return nil, nil
	}
	for _, value := range ae {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_especifica_necesidad/"
		var aeOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &aeOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postActividadesEspecificas", "Error": err.Error()}
		}
		aeOut["NecesidadId"] = nil
		out = append(out, &aeOut)
	}
	return out, nil

}

// post requisitos minimos necesidad
func postRequisitosNecesidad(rm []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if rm == nil || len(rm) == 0 {
		return nil, nil
	}
	for _, value := range rm {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo_necesidad/"
		var rmOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &rmOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postpostRequisitosNecesidad", "Error": err.Error()}
		}
		rmOut["NecesidadId"] = nil
		out = append(out, &rmOut)
	}
	return out, nil

}

// post actividad economica

func postActividadesEconomicas(ae []*map[string]interface{}, necesidad *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if ae == nil || len(ae) == 0 {
		return nil, nil
	}
	for _, value := range ae {
		(*value)["NecesidadId"] = necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_economica_necesidad/"
		var aeOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &aeOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postActividadesEspecificas", "Error": err.Error()}
		}
		aeOut["NecesidadId"] = nil
		out = append(out, &aeOut)
	}
	return out, nil

}

// post rubro
func postRubros(rubros []*models.RubroNecesidad, necesidad *map[string]interface{}) (out []*models.RubroNecesidad, outErr map[string]interface{}) {
	var errOut map[string]interface{}
	if rubros == nil || len(rubros) == 0 {
		return nil, nil
	}
	for _, rubro := range rubros {
		rubro.NecesidadId = *necesidad
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "rubro_necesidad/"
		var rOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &rOut, rubro); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postRubros", "Error": err.Error()}
		}
		rOut["NecesidadId"] = nil
		if rOut["Fuentes"], errOut = postFuentesRubro(rubro.Fuentes, &rOut); errOut == nil {

		}
		if rOut["Productos"], errOut = postProductosRubro(rubro.Productos, &rOut); errOut == nil {

		}
		if rOut["Metas"], errOut = postMetasRubro(rubro.Metas, &rOut); errOut == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postRubros", "Error": errOut}
		}
		var rubroOut models.RubroNecesidad
		if errConvert := formatdata.FillStruct(rOut, &rubroOut); errConvert == nil {
			out = append(out, &rubroOut)
		}

	}
	return out, nil
}

// post fuentes

func postFuentesRubro(f []*map[string]interface{}, rubro *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if f == nil || len(f) == 0 {
		return nil, nil
	}
	for _, value := range f {
		(*value)["RubroNecesidadId"] = rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_rubro_necesidad/"
		var fOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &fOut, value); err == nil {
		} else {
			return nil, map[string]interface{}{"Function": "postFuentesRubro", "Error": err.Error()}
		}
		fOut["RubroNecesidadId"] = nil
		out = append(out, &fOut)
	}
	return out, nil
}

// post productos

func postProductosRubro(p []*map[string]interface{}, rubro *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if p == nil || len(p) == 0 {
		return nil, nil
	}
	for _, value := range p {
		(*value)["RubroNecesidadId"] = rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_rubro_necesidad/"
		var pOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &pOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postProductosRubro", "Error": err.Error()}
		}
		pOut["RubroNecesidadId"] = nil
		out = append(out, &pOut)
	}
	return out, nil

}

// post metas

func postMetasRubro(m []*models.MetaRubroNecesidad, rubro *map[string]interface{}) (out []*models.MetaRubroNecesidad, outErr map[string]interface{}) {
	var errOut map[string]interface{}
	if m == nil || len(m) == 0 {
		return nil, nil
	}
	for _, meta := range m {
		meta.RubroNecesidadId = *rubro
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "meta_rubro_necesidad/"
		var mOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &mOut, meta); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postMetasRubro", "Error": err.Error()}
		}
		mOut["RubroNecesidadId"] = nil
		if mOut["Actividades"], errOut = postActividadesMeta(meta.Actividades, &mOut); errOut == nil {

		}
		var metaOut models.MetaRubroNecesidad
		if errConvert := formatdata.FillStruct(mOut, &metaOut); errConvert == nil {
			out = append(out, &metaOut)
		}

	}
	return out, nil

}

// post actividades
func postActividadesMeta(act []*map[string]interface{}, meta *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if act == nil || len(act) == 0 {
		return nil, nil
	}
	for _, value := range act {
		(*value)["MetaRubroNecesidadId"] = meta
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_meta/"
		var actOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &actOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postActividadesMeta", "Error": err.Error()}
		}
		actOut["MetaRubroNecesidadId"] = nil
		postFuentesActividad((*value)["FuentesActividad"].([]interface{}), &actOut)
		out = append(out, &actOut)
	}
	return out, nil

}

// post fuentes actividades

func postFuentesActividad(fuentes []interface{}, act *map[string]interface{}) (out []*map[string]interface{}, outErr map[string]interface{}) {
	if fuentes == nil || len(fuentes) == 0 {
		return nil, nil
	}
	for _, value := range fuentes {
		value.(map[string]interface{})["ActividadMetaNecesidadId"] = act
		urlcrud := beego.AppConfig.String("necesidadesCrudService") + "fuente_actividad/"
		var fOut map[string]interface{}
		if err := request.SendJson(urlcrud, "POST", &fOut, value); err == nil {

		} else {
			return nil, map[string]interface{}{"Function": "postFuentesActividad", "Error": err.Error()}
		}
		fOut["ActividadMetaNecesidadId"] = nil
		out = append(out, &fOut)
	}
	return out, nil

}
