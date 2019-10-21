package necesidadhelper

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// GetTrNecesidad obtiene necesidad de crud apí con sus objetos relacionados
func GetTrNecesidad(id string) (trnecesidad models.TrNecesidad, outErr map[string]interface{}) {
	var err map[string]interface{}
	if trnecesidad.Necesidad, err = getNecesidadCrud(id); err != nil {
		return trnecesidad, err
	} else {
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
		if trnecesidad.Rubros, err = getRubrosNecesidad(id); err != nil {
			return trnecesidad, err
		}
		return trnecesidad, nil
	}

}

// traer la necesidad
func getNecesidadCrud(id string) (necesidad *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "necesidad/?query=Id:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getNecesidadCrud", "Error": err.Error()}
		return nil, outErr
	} else {
		necesidad := &res[0]
		return necesidad, nil
	}
}

// TODO se pueden generalizar las funciones que traen valores y arreglos con un closure para reducir lineas de codigo
// traer detalle servicio asociado a la necesidad
func getDetalleServicio(id string) (ds *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_servicio_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetalleServicio", "Error": err.Error()}
		return nil, outErr
	} else {
		if len(res[0]) > 0 {
			res[0]["NecesidadId"] = nil
		}
		ds := &res[0]
		return ds, nil
	}
}

// traer detalle prestacion servicio asociado a la necesidad
func getDetallePrestacionServicio(id string) (dps *map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "detalle_prestacion_servicio/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getDetallePrestacionServicio", "Error": err.Error()}
		return nil, outErr
	} else {
		if len(res[0]) > 0 {
			res[0]["NecesidadId"] = nil
		}
		dps := &res[0]
		return dps, nil
	}
}

// traer productos catalogo asociados a la necesidad
func getProductosCatalogo(id string) (productos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "producto_catalogo_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getProductosCatalogo", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["NecesidadId"] = nil
				if value["RequisitosMinimos"], outErr = getRequisitosProducto(fmt.Sprintf("%.0f", value["Id"].(float64))); outErr != nil {
					return nil, outErr
				} else {
					productos = append(productos, &value)
				}
			}

		}
		return productos, nil
	}
}

// traer req minimos de un producto catalogo
func getRequisitosProducto(id string) (requisitos []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "requisito_minimo/?query=ProductoCatalogoNecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRequisitosProducto", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["ProductoCatalogoNecesidadId"] = nil
				requisitos = append(requisitos, &value)
			}
		}
		return requisitos, nil
	}
}

// get marco legal de la necesidad
func getMarcoLegal(id string) (ml []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "marco_legal_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getMarcoLegal", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["NecesidadId"] = nil
				ml = append(ml, &value)
			}

		}
		return ml, nil
	}
}

// get actividad especifica de la necesidad
func getActividadEspecifica(id string) (ae []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_especifica_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEspecifica", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["NecesidadId"] = nil
				ae = append(ae, &value)
			}

		}
		return ae, nil
	}
}

// get actividad economica de la necesidad
func getActividadEconomica(id string) (aec []*map[string]interface{}, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "actividad_economica_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getActividadEconomica", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["NecesidadId"] = nil
				aec = append(aec, &value)
			}

		}
		return aec, nil
	}
}

// get rubros de la necesidad
func getRubrosNecesidad(id string) (rubros []*models.RubroNecesidad, outErr map[string]interface{}) {
	urlcrud := beego.AppConfig.String("necesidadesCrudService") + "rubro_necesidad/?query=NecesidadId:" + id
	var res []map[string]interface{}
	if err := request.GetJson(urlcrud, &res); err != nil {
		outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
		return nil, outErr
	} else {
		for _, value := range res {
			if len(value) > 0 {
				value["NecesidadId"] = nil
				var rubro models.RubroNecesidad
				if err = formatdata.FillStruct(value, &rubro); err != nil {
					outErr = map[string]interface{}{"Function": "getRubrosNecesidad", "Error": err.Error()}
					return nil, outErr
				} else {
					fmt.Println(rubros)
					rubros = append(rubros, &rubro)
				}
			}

		}

		return rubros, nil
	}
}
