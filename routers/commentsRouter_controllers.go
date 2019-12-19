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

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CdpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CdpController"],
        beego.ControllerComments{
            Method: "ExpedirCdp",
            Router: `/expedirCDP/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CdpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CdpController"],
        beego.ControllerComments{
            Method: "SolicitarCdp",
            Router: `/solicitarCDP`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"],
        beego.ControllerComments{
            Method: "ExpedirCrp",
            Router: `/expedirCRP/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"],
        beego.ControllerComments{
            Method: "GetInfoCrp",
            Router: `/getFullCrp`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:CrpController"],
        beego.ControllerComments{
            Method: "SolicitarCrp",
            Router: `/solicitarCRP`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id/:vigencia/:unidadEjecutora`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"],
        beego.ControllerComments{
            Method: "RegistrarModificacion",
            Router: `/modificacion`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"],
        beego.ControllerComments{
            Method: "GetRubrosbyFuente",
            Router: `/plan_adquisiciones_rubros_fuente/:vigencia/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:FuenteFinanciamientoApropiacionController"],
        beego.ControllerComments{
            Method: "SimulacionAfectacion",
            Router: `/simulacion_afectacion_modificacion/:centroGestor/:vigencia`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"],
        beego.ControllerComments{
            Method: "GetAllModificacionPresupuestalByVigenciaAndCG",
            Router: `/:vigencia/:CG`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:ModificacionPresupuestalController"],
        beego.ControllerComments{
            Method: "SimulacionAfectacion",
            Router: `/simulacion_afectacion_modificacion/:centroGestor/:vigencia`,
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
            Method: "GetFullNecesidad",
            Router: `/getfullnecesidad/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:NecesidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:NecesidadController"],
        beego.ControllerComments{
            Method: "PostFullNecesidad",
            Router: `/post_full_necesidad`,
            AllowHTTPMethods: []string{"post"},
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

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:VigenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:VigenciaController"],
        beego.ControllerComments{
            Method: "CerrarVigencia",
            Router: `/cerrar_vigencia/:vigencia/:area`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:VigenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/plan_cuentas_mid/controllers:VigenciaController"],
        beego.ControllerComments{
            Method: "GetCierreVigencia",
            Router: `/get_cierre/:vigencia/:area`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
