// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/plan_cuentas_mid/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/rubro",
			beego.NSInclude(
				&controllers.RubroController{},
			),
		),
		beego.NSNamespace("/apropiacion",
			beego.NSInclude(
				&controllers.ApropiacionController{},
			),
		),
		beego.NSNamespace("/aprobacion_apropiacion",
			beego.NSInclude(
				&controllers.AprobacionController{},
			),
		),
		beego.NSNamespace("/fuente_financiamiento_apropiacion",
			beego.NSInclude(
				&controllers.FuenteFinanciamientoApropiacionController{},
			),
		),
		beego.NSNamespace("/movimiento",
			beego.NSInclude(
				&controllers.MovimientoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
