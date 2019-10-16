package cdphelper

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

// ExpedirCRP init necesidad
func ExpedirCrp(id string) (crp map[string]interface{}, outErr map[string]interface{}) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/" + id
	var solicitud = make(map[string]map[string]interface{})
	crp = make(map[string]interface{})
	solicitud["infoCrp"] = make(map[string]interface{})
	solicitud["infoCrp"]["consecutivo"] = GetConsecutivoCRP()
	solicitud["infoCrp"]["fechaExpedicion"] = time.Now().Format(time.RFC3339)
	solicitud["infoCrp"]["estado"] = 1
	if err := request.SendJson(urlmongo, "PUT", &crp, &solicitud); err == nil {
		crp["data"] = solicitud
		return crp, nil
	} else {
		outErr = map[string]interface{}{"Function": "ExpedirCRP", "Error": err.Error()}
		return nil, outErr
	}

}

// SolicitarCRP init necesidad
func SolicitarCRP(solCrp map[string]interface{}) (solicitud map[string]interface{}, outErr map[string]interface{}) {
	var (
		okBeneficiario   bool
		okCdp            bool
		okVigencia       bool
		okValor          bool
		okNumCompromiso  bool
		okTipoCompromiso bool
		mongoData        map[string]interface{}
		compromiso       map[string]interface{}
	)

	solicitud = make(map[string]interface{})
	solicitud["consecutivo"] = GetConsecutivoSolicitudCRP()
	solicitud["consecutivoCdp"], okCdp = solCrp["consecutivoCdp"]
	solicitud["beneficiario"], okBeneficiario = solCrp["beneficiario"]
	solicitud["vigencia"], okVigencia = solCrp["vigencia"]
	solicitud["valor"], okValor = solCrp["monto"]
	compromiso = make(map[string]interface{})
	compromiso["numeroCompromiso"], okNumCompromiso = solCrp["numCompromiso"]
	compromiso["tipoCompromiso"], okTipoCompromiso = solCrp["compromiso"]
	solicitud["compromiso"] = compromiso
	solicitud["infoCrp"] = nil
	solicitud["activo"] = true
	solicitud["fechaCreacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaModificacion"] = time.Now().Format(time.RFC3339)
	beego.Info(solicitud)
	beego.Info(okBeneficiario, "la", okVigencia, "le", okNumCompromiso, "li", okValor, "Lo", okNumCompromiso, "Lu", okTipoCompromiso, "lad", okCdp)
	if okBeneficiario && okCdp && okVigencia && okValor && okNumCompromiso && okTipoCompromiso {
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

// GetCosecutivoSolicitudCDP Get Cosecutivo Solicitud crp
func GetConsecutivoSolicitudCRP() (consecutivo int64) {
	return 100
}

// GetCosecutivoSolicitudCDP Get Cosecutivo Solicitud crp
func GetConsecutivoCRP() (consecutivo int64) {
	return 666
}
