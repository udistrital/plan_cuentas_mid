package fuenteapropiacionhelper

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

// URLCRUD Ruta de plan_cuentas_crud
var URLCRUD = beego.AppConfig.String("planCuentasApiService") + "fuente_financiamiento_apropiacion/registrar_multiple"

// RegistrarMultipleFuenteApropiacion utiliza la transacción de fuente_financiamiento_apropiacion/registrar_multiple
// para registrar multiples datos en fuente_financiamiento_apropiacion
func RegistrarMultipleFuenteApropiacion(fuentesApropiacion []*models.FuenteFinanciamientoApropiacion) (idRegistrados []int) {
	var (
		res          map[string]interface{}
		bodyResponse map[string][]int
	)

	if err := request.SendJson(URLCRUD, "POST", &res, &fuentesApropiacion); err != nil {
		log.Panicln(err.Error())
		return
	}

	if err := formatdata.FillStruct(res["Body"], &bodyResponse); err != nil {
		log.Panicln(err.Error())
		return
	}
	idRegistrados = bodyResponse["Ids"]
	return
}

// ConcatenarFuente recibe un arreglo de interfaces y las contatena para formar un arreglo de models.FuenteFinanciamientoApropiacion
func ConcatenarFuente(fuenteFinanciamiento *models.FuenteFinanciamiento, fuenteApropiaciones ...interface{}) []*models.FuenteFinanciamientoApropiacion {
	var (
		apropiacion        *models.Apropiacion
		fuenteApropiacion  *models.FuenteFinanciamientoApropiacion   // una sola fuente
		fuentesApropiacion []*models.FuenteFinanciamientoApropiacion // Un conjunto de fuentes apropiacion
	)

	for _, value := range fuenteApropiaciones {
		// // Obtiene la información de la apropición
		if err := formatdata.FillStruct(value.(map[string]interface{})["Apropiacion"], &apropiacion); err != nil {
			log.Panicln(err.Error())
			return nil
		}

		/* Convierte toda la información de value a tipo fuenteFinanciamiento
		(no se pueden convertir automáticamente structs definidos, como apropiacion y fuenteFinanciamiento) */
		if err := formatdata.FillStruct(value, &fuenteApropiacion); err != nil {
			log.Panicln(err.Error())
			return nil
		}
		fuenteApropiacion.ApropiacionId = apropiacion
		fuenteApropiacion.FuenteFinanciamientoId = fuenteFinanciamiento

		fuentesApropiacion = append(fuentesApropiacion, fuenteApropiacion)
	}

	return fuentesApropiacion
}
