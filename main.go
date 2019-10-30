package main

import (
	_ "github.com/udistrital/plan_cuentas_mid/routers"
	//"github.com/udistrital/utils_oas/customerror"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	apistatus "github.com/udistrital/utils_oas/apiStatusLib"
)

func init() {
}

func main() {
	// beego.BConfig.RecoverFunc = responseformat.GlobalResponseHandler
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//beego.ErrorController(&customerror.CustomErrorController{})
	apistatus.Init()

	//mongoProcess.PresupuestoMongoJobInit()
	beego.Run()

}
