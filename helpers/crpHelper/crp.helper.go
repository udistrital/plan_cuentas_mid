package crphelper

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// ExpedirCRP init necesidad
func ExpedirCrp(id string) (crp map[string]interface{}, outErr map[string]interface{}) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/" + id
	var resMongo = make(map[string]interface{})
	var solicitud = make(map[string]interface{})
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitud = resMongo["Body"].(map[string]interface{})
		crp = make(map[string]interface{})
		solicitud["infoCrp"] = make(map[string]interface{})
		solicitud["infoCrp"] = map[string]interface{}{"consecutivo": GetConsecutivoCRP(), "fechaExpedicion": time.Now().Format(time.RFC3339), "estado": 1}
		if err := request.SendJson(urlmongo, "PUT", &crp, &solicitud); err == nil {
			crp = solicitud
			return crp, nil
		} else {
			outErr = map[string]interface{}{"Function": "ExpedirCrp", "Error": err.Error()}
			return nil, outErr
		}
	} else {
		outErr = map[string]interface{}{"Function": "ExpedirCrp", "Error": err.Error()}
		return nil, outErr
	}
}

// SolicitarCRP init necesidad
func SolicitarCRP(solCrp map[string]interface{}) (solicitud map[string]interface{}, outErr map[string]interface{}) {
	var (
		okBeneficiario bool
		okCdp          bool
		okVigencia     bool
		okValor        bool
		okCompromiso   bool
		mongoData      map[string]interface{}
	)

	solicitud = make(map[string]interface{})
	solicitud["consecutivo"] = GetConsecutivoSolicitudCRP()
	solicitud["consecutivoCdp"], okCdp = solCrp["consecutivoCdp"]
	solicitud["beneficiario"], okBeneficiario = solCrp["beneficiario"]
	solicitud["vigencia"], okVigencia = solCrp["vigencia"]
	solicitud["valor"], okValor = solCrp["monto"]
	solicitud["compromiso"], okCompromiso = solCrp["compromiso"]
	solicitud["infoCrp"] = nil
	solicitud["activo"] = true
	solicitud["fechaCreacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaModificacion"] = time.Now().Format(time.RFC3339)
	beego.Info(solicitud)
	beego.Info(okBeneficiario, "la", okVigencia, "le", okCompromiso, "li", okValor, "lad", okCdp)
	if okBeneficiario && okCdp && okVigencia && okCompromiso {
		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/"
		if err := request.SendJson(urlmongo, "POST", &mongoData, &solicitud); err == nil {
			return solicitud, nil
		} else {
			outErr = map[string]interface{}{"Function": "SolicitarCRP", "Error": err.Error()}
			return nil, outErr
		}
	} else {
		outErr = map[string]interface{}{"Function": "SolicitarCRP", "Error": "Datos incompletos del CDP"}
		return nil, outErr
	}

}

// GetCosecutivoSolicitudCRP Get Cosecutivo Solicitud crp
func GetConsecutivoSolicitudCRP() (consecutivo int) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
		fmt.Println("solicitudes", solicitudes)
	}
	consecutivo = len(solicitudes) + 1
	return consecutivo
}

// GetCosecutivoSolicitudCRP Get Cosecutivo Solicitud crp
func GetConsecutivoCRP() (consecutivo int) {
	total := GetConsecutivoSolicitudCRP()
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
		for k, _ := range solicitudes {
			if solicitudes[k].(map[string]interface{})["infoCrp"] == nil {
				total = total - 1
			}
		}
	}
	consecutivo = total
	return consecutivo
}
