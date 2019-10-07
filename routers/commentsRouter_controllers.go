package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"],
        beego.ControllerComments{
            Method: "AprobacionAsignacionInicial",
            Router: `/AprobacionAsignacionInicial/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"],
        beego.ControllerComments{
            Method: "Aprobado",
            Router: `/Aprobado`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:AprobacionController"],
        beego.ControllerComments{
            Method: "InformacionAsignacionInicial",
            Router: `/InformacionAsignacionInicial/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id/:valor/:vigencia`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "ArbolApropiaciones",
            Router: `/ArbolApropiaciones/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ApropiacionController"],
        beego.ControllerComments{
            Method: "SaldoApropiacion",
            Router: `/SaldoApropiacion/:rubro/:unidadEjecutora/:vigencia`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"],
        beego.ControllerComments{
            Method: "RegistrarFuenteConApropiacion",
            Router: `registrar_fuentes_con_apropiacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:MovimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:MovimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "InitNecesidad",
            Router: `/initnecesidad/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
