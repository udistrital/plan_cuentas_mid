package cdphelper

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// ExpedirCDP init necesidad
func ExpedirCdp(id string) (cdp map[string]interface{}, outErr map[string]interface{}) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/" + id
	var resMongo = make(map[string]interface{})
	var solicitud = make(map[string]interface{})
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitud = resMongo["Body"].(map[string]interface{})
		cdp = make(map[string]interface{})
		solicitud["infoCdp"] = make(map[string]interface{})
		solicitud["infoCdp"] = map[string]interface{}{"consecutivo": GetConsecutivoCDP(), "fechaExpedicion": time.Now().Format(time.RFC3339), "estado": 1}
		if err := request.SendJson(urlmongo, "PUT", &cdp, &solicitud); err == nil {
			cdp = solicitud
			return cdp, nil
		} else {
			outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": err.Error()}
			return nil, outErr
		}
	} else {
		outErr = map[string]interface{}{"Function": "SolicitarCDP", "Error": err.Error()}
		return nil, outErr
	}

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
	solicitud["estado"] = 1
	solicitud["justificacionRechazo"] = ""
	solicitud["infoCdp"] = nil
	solicitud["activo"] = true
	solicitud["fechaCreacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaModificacion"] = time.Now().Format(time.RFC3339)
	beego.Info(solicitud)
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

// GetCosecutivoSolicitudCDP Get Cosecutivo Solicitud CDP
func GetConsecutivoSolicitudCDP() (consecutivo int) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
		fmt.Println("solicitudes", solicitudes)
	}
	consecutivo = len(solicitudes) + 1
	return consecutivo
}

// GetCosecutivoSolicitudCDP Get Cosecutivo Solicitud CDP
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
