package crphelper

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// ExpedirCrp init necesidad
func ExpedirCrp(id string) (crp map[string]interface{}, outErr map[string]interface{}) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/" + id
	var resMongo = make(map[string]interface{})
	var solicitud = make(map[string]interface{})
	var err, errPut error
	if err = request.GetJson(urlmongo, &resMongo); err == nil {
		solicitud = resMongo["Body"].(map[string]interface{})
		crp = make(map[string]interface{})
		solicitud["estado"] = models.GetEstadoExpedidoCrp()
		if errPut = request.SendJson(urlmongo, "PUT", &crp, &solicitud); errPut == nil {
			crp = solicitud
			return crp, nil
		}
		outErr = map[string]interface{}{"Function": "ExpedirCrp", "Error": errPut.Error()}
		return nil, outErr

	}
	outErr = map[string]interface{}{"Function": "ExpedirCrp", "Error": err.Error()}
	return nil, outErr
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
	solicitud["estado"] = models.GetEstadoSolicitudCrp()
	solicitud["activo"] = true
	solicitud["fechaCreacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaModificacion"] = time.Now().Format(time.RFC3339)
	solicitud["fechaInicioVigencia"] = solCrp["fechaInicioVigencia"]
	solicitud["fechaFinalVigencia"] = solCrp["fechaFinalVigencia"]

	if okBeneficiario && okCdp && okVigencia && okCompromiso && okValor {

		urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/"
		if err := request.SendJson(urlmongo, "POST", &mongoData, &solicitud); err != nil {
			outErr = map[string]interface{}{"Function": "SolicitarCRP", "Error": err.Error()}
		}
	} else {
		outErr = map[string]interface{}{"Function": "SolicitarCRP", "Error": "Datos incompletos del CDP"}
	}

	return
}

// GetCosecutivoSolicitudCRP ...get Cosecutivo Solicitud crp
func GetConsecutivoSolicitudCRP() (consecutivo int) {
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP/"
	fmt.Println(urlmongo)
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		if _, ok := resMongo["Body"].([]interface{}); ok {
			solicitudes = resMongo["Body"].([]interface{})
		}
	}
	consecutivo = len(solicitudes) + 1
	return consecutivo
}

// GetConsecutivoCRP Get Cosecutivo Solicitud crp
func GetConsecutivoCRP() (consecutivo int) {
	total := GetConsecutivoSolicitudCRP()
	urlmongo := beego.AppConfig.String("financieraMongoCurdApiService") + "solicitudesCRP"
	var resMongo = make(map[string]interface{})
	var solicitudes []interface{}
	if err := request.GetJson(urlmongo, &resMongo); err == nil {
		solicitudes = resMongo["Body"].([]interface{})
		for k := range solicitudes {
			if solicitudes[k].(map[string]interface{})["infoCrp"] == nil {
				total = total - 1
			}
		}
	}
	consecutivo = total
	return consecutivo
}

// GetFullCRP ...
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
				urlcrud = beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal/" + strVig + "/1?query=consecutivo:" + strAux + ",tipo:cdp"

				if response2, errCdp := request.GetJsonTest(urlcrud, &auxObjCdp); errCdp == nil {
					objTmpCrp["solicitudCrp"] = solct["consecutivo"]
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
								necesidadID := bodyNecesidad["necesidad"]

								strnecesidadID := fmt.Sprintf("%v", necesidadID)
								urlcrud = beego.AppConfig.String("necesidadesCrudService") + "necesidad/" + strnecesidadID
								if response4, err := request.GetJsonTest(urlcrud, &objNecesidad); err == nil {
									tipoFinanciacion := objNecesidad["TipoFinanciacionNecesidadId"].(map[string]interface{})
									logs.Debug(objNecesidad["TipoFinanciacionNecesidadId"])
									logs.Debug(tipoFinanciacion["Nombre"])

									objTmpCrp["consecutivoCdp"] = aux
									objTmpCrp["vigencia"] = vig

									objTmpCrp["movimiento_cdp"] = solctCdp["AfectacionIds"]
									objTmpCrp["centroGestor"] = solctCdp["CentroGestor"] // objeto de objetos
									objTmpCrp["estado"] = solctCdp["Estado"]
									objTmpCrp["necesidadFinanciacion"] = tipoFinanciacion["Id"] // REEMPLAZAR POR NOMBRE CUANDO EXISTA
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
				objTmpCrp = make(map[string]interface{})
			}

		} else {
			logs.Info("Error (2) servicio caido")
			logs.Debug(err)
			outputError = map[string]interface{}{"Function": "GetCDP, ", "Error": err}
			return nil, outputError
		}

	} else {
		logs.Info("Error (1) servicio caido")
		logs.Debug(err, response)
		outputError = map[string]interface{}{"Function": "GetSolicitudesCRP", "Error": err}
		return nil, outputError
	}
	return consultaCrps, nil

}
