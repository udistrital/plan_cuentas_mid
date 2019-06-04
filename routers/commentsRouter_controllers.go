package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"],
        beego.ControllerComments{
            Method: "ArbolRubros",
            Router: `/ArbolRubros/:unidadEjecutora`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"],
        beego.ControllerComments{
            Method: "EliminarRubro",
            Router: `/EliminarRubro/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:RubroController"],
        beego.ControllerComments{
            Method: "RegistrarRubro",
            Router: `/RegistrarRubro/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
