package crphelper

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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

// GetFullCRP...
func GetFullCrp() (consultaCrps []map[string]interface{}, outputError map[string]interface{}) {
	var (
		urlcrud           string
		resSolicitudesCrp map[string]interface{}
		resNecesidad      map[string]interface{}
		//solCrps           []map[string]interface{}
		objNecesidad map[string]interface{}
		auxObjCdp    map[string]interface{}
		objTmpCrp    map[string]interface{}
	)

	// Solicita información de solicitudes de CRP
	urlcrud = beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP"
	if response, err := request.GetJsonTest(urlcrud, &resSolicitudesCrp); err == nil { // (2) error servicio caido
		aux := resSolicitudesCrp["Body"]
		objTmpCrp = make(map[string]interface{})
		if val, ok := aux.([]interface{}); ok {
			for _, sol := range val {

				solct := sol.(map[string]interface{})
				aux := solct["consecutivoCdp"]
				vig := solct["vigencia"]

				strAux := fmt.Sprintf("%v", aux)
				strVig := fmt.Sprintf("%v", vig)

				// consulta los CDP expedidos a los que van ligados esas solicitudes dea CRP
				urlcrud = beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal/" + strVig + "/1?query=data.consecutivo_cdp:" + strAux + ",tipo:cdp,estado:expedido"

				if response2, err := request.GetJsonTest(urlcrud, &auxObjCdp); err == nil {

					auxCdpInterface := auxObjCdp["Body"]

					if auxCdpInterface != nil {
						auxCdp := auxCdpInterface.([]interface{})

						for _, solCdp := range auxCdp {
							solctCdp := solCdp.(map[string]interface{})

							// Consulta ID necesidad

							idCDP := solctCdp["Data"].(map[string]interface{})
							idCDPAux := idCDP["solicitud_cdp"]

							strCDP := fmt.Sprintf("%v", idCDPAux)

							urlcrud = beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCDP/" + strCDP

							if response3, err := request.GetJsonTest(urlcrud, &resNecesidad); err == nil {

								// Consulta necesidad
								bodyNecesidad := resNecesidad["Body"].(map[string]interface{})
								necesidadId := bodyNecesidad["necesidad"]

								strnecesidadId := fmt.Sprintf("%v", necesidadId)
								urlcrud = beego.AppConfig.String("necesidadesCrudService") + "necesidad/" + strnecesidadId
								if response4, err := request.GetJsonTest(urlcrud, &objNecesidad); err == nil {
									logs.Info(response4)
									tipoFinanciacion := objNecesidad["TipoFinanciacionNecesidadId"].(map[string]interface{})
									logs.Debug(objNecesidad["TipoFinanciacionNecesidadId"])
									logs.Debug(tipoFinanciacion["Nombre"])

									objTmpCrp["consecutivoCdp"] = aux
									objTmpCrp["vigencia"] = vig

									objTmpCrp["centroGestor"] = solctCdp["CentroGestor"] // objeto de objetos
									objTmpCrp["estado"] = solctCdp["Estado"]
									objTmpCrp["necesidad"] = tipoFinanciacion["Id"] // REEMPLAZAR POR NOMBRE CUANDO EXISTA
									logs.Info(objTmpCrp, "banderita")
									consultaCrps = append(consultaCrps, objTmpCrp)

								} else {
									logs.Info("Error (6) Necesidad")
									outputError = map[string]interface{}{"Function": "GetFullCrp:getFullCrp at documento pres", "Error": response4.Status}
									return nil, outputError
								}

							} else {
								logs.Info("Error (5) Necesidad")
								outputError = map[string]interface{}{"Function": "GetFullCrp:getFullCrp at documento pres", "Error": response3.Status}
								return nil, outputError
							}

						}

					}

				} else {
					logs.Info("Error (4) Documento Presupuestal")
					outputError = map[string]interface{}{"Function": "GetFullCrp:getFullCrp at documento pres", "Error": response2.Status}
					return nil, outputError
				}
			}
			return consultaCrps, nil

		} else {
			logs.Info("Error (2) servicio caido")
			logs.Debug(err)
			outputError = map[string]interface{}{"Function": "GetEntrada, Es aca?ya", "Error": err}
			return nil, outputError
		}

	} else {
		logs.Info("Error (1) servicio caido")
		logs.Debug(err)
		logs.Info(response)
		outputError = map[string]interface{}{"Function": "GetSolicitudesCRP, Es aca?ya", "Error": err}
		return nil, outputError
	}

}
