package apropiacionmanager

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// apropiacionURI ...
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
var apropiacionURI = beego.AppConfig.String("planCuentasApiService") + "apropiacion" + "/"

// GetAprByCodigoAndVigencia ... Return Apropiation Info by rubro's COde and vigencia.
//
// Deprecated: Depende de PLAN_CUENTAS_CRUD (ya no está en servicio)
func GetAprByCodigoAndVigencia(codigo string, vigencia int) (aprComp []models.Apropiacion, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if err = request.GetJson(apropiacionURI+"?query=RubroId.Codigo:"+codigo+",Vigencia:"+strconv.Itoa(vigencia), &aprComp); err != nil {
		if len(aprComp) > 0 {
			err = errors.New("No Data Found")
		}
	}

	return

}
