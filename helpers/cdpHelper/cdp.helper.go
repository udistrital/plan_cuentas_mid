package cdphelper

import (
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// ExpedirCdp init necesidad
func ExpedirCdp(id string) (cdp map[string]interface{}, outErr map[string]interface{}) {
	var resMongo, solicitud map[string]interface{}
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/" + id

	if err := request.GetJson(urlmongo, &resMongo); err == nil {

		solicitud = resMongo["Body"].(map[string]interface{})
		solicitud["estado"] = models.GetEstadoExpedidoCdp()
		if err := request.SendJson(urlmongo, "PUT", &cdp, &solicitud); err == nil {
			cdp = solicitud
		} else {
			outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": err.Error()}
		}

	} else {

		outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": err.Error()}
	}

	return
}

// SolicitarCDP init necesidad
func SolicitarCDP(necesidad map[string]interface{}) (solicitud map[string]interface{}, outErr map[string]interface{}) {
	var (
		okAreaFuncional bool
		okId            bool
		okVigencia      bool
		mongoData       map[string]interface{}
	)

	solicitud = make(map[string]interface{})
	solicitud["consecutivo"] = GetConsecutivoSolicitudCDP()
	solicitud["entidad"] = 1
	solicitud["centroGestor"], okAreaFuncional = necesidad["AreaFuncional"]
	solicitud["necesidad"], okId = necesidad["Id"]
	solicitud["vigencia"], okVigencia = necesidad["Vigencia"]
	solicitud["fechaRegistro"] = time.Now().Format(time.RFC3339)
	solicitud["estado"] = models.GetEstadoSolicitudCdp()
	solicitud["activo"] = true
	solicitud["fechaCreacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaModificacion"] = time.Now().Format(time.RFC3339)
	if okAreaFuncional && okId && okVigencia {
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/"
		if err := request.SendJson(urlmongo, "POST", &mongoData, &solicitud); err == nil {
			return solicitud, nil
		} else {
			outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": err.Error()}
			return nil, outErr
		}
	} else {
		outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": "Datos incompletos en necesidad"}
		return nil, outErr
	}

}

// AprobarCdp
func AprobarCdp(id, vigencia, areaFuncional string) (docPresupuestal map[string]interface{}, outErr map[string]interface{}) {
	var response map[string]interface{}

	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal/"

	if err := request.GetJson(urlmongo+"documento/" + vigencia + "/"+ areaFuncional + "/" + id, &response); err != nil {
		outErr = map[string]interface{}{"Function": "AprobarCdp", "Error": "No se pudo obtener el documento"}

	} else {
		docPresupuestal = response["Body"].(map[string]interface{})
		docPresupuestal["Estado"] = "aprobado"

		if err = request.SendJson(urlmongo + vigencia + "/" + areaFuncional + "/" + id, "PUT", &response, &docPresupuestal); err != nil {
			outErr = map[string]interface{}{"Function": "AprobarCdp", "Error": "No se actualizar el documento"}
		} 
	}
	
	return
}

// GetConsecutivoSolicitudCDP Get Cosecutivo Solicitud CDP
func GetConsecutivoSolicitudCDP() (consecutivo int) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
	}
	consecutivo = len(solicitudes) + 1
	return consecutivo
}

// GetConsecutivoCDP Get Cosecutivo Solicitud CDP
func GetConsecutivoCDP() (consecutivo int) {
	total := GetConsecutivoSolicitudCDP()
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
		for k, _ := range solicitudes {
			if solicitudes[k].(map[string]interface{})["infoCdp"] == nil {
				total = total - 1
			}
		}
	}
	consecutivo = total
	return consecutivo
}
