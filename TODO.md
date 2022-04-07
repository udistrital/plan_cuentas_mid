# Por hacer

## Limpiar y/o actualizar dependencias de `planCuentasApiService = ${PLAN_CUENTAS_CRUD}`

- [ ] [conf/app.conf](conf/app.conf)
- [ ] `AddApropiacion()` en [helpers/apropiacionHelper/apropiacion.helper.go](helpers/apropiacionHelper/apropiacion.helper.go)
- [ ] `PutApropiacion()` en [helpers/apropiacionHelper/apropiacion.helper.go](helpers/apropiacionHelper/apropiacion.helper.go)
- [ ] `AprobarPresupuesto()` en [helpers/apropiacionHelper/apropiacion.helper.go](helpers/apropiacionHelper/apropiacion.helper.go)
- [ ] `PresupuestoAprobado()` en [helpers/apropiacionHelper/apropiacion.helper.go](helpers/apropiacionHelper/apropiacion.helper.go)
- [ ] `RegistrarMultipleFuenteApropiacion()` en [helpers/fuenteApropiacionHelper/fuente_apropiacion.helper.go](helpers/fuenteApropiacionHelper/fuente_apropiacion.helper.go)
- [ ] `RegistrarFuenteHelper()` en [helpers/fuenteFinanciamientoHelper/fuente_financiamiento.helper.go](helpers/fuenteFinanciamientoHelper/fuente_financiamiento.helper.go)
- [ ] `AddRubro()` en [helpers/rubroHelper/rubro.helper.go](helpers/rubroHelper/rubro.helper.go)
- [ ] `DeleteRubro()` en [helpers/rubroHelper/rubro.helper.go](helpers/rubroHelper/rubro.helper.go)
- [x] `GetAprByCodigoAndVigencia()` en [managers/apropiacionManager/apropiacion.manager.go](managers/apropiacionManager/apropiacion.manager.go)

Consecuencia de lo anterior, revisar/ajustar donde se consuman las anteriores:

- [ ] `Post()` en [controllers/apropiacion.go](controllers/apropiacion.go)
- [ ] `Put()` en [controllers/apropiacion.go](controllers/apropiacion.go)
- [ ] `AprobacionAsignacionInicial()` en [controllers/aprobacion_apropiacion.go](controllers/aprobacion_apropiacion.go)
- [ ] `Aprobado()` en [controllers/aprobacion_apropiacion.go](controllers/aprobacion_apropiacion.go)
- [ ] `RegistrarFuenteConApropiacion()` en [controllers/fuente_financiamiento_apropiacion.go](controllers/fuente_financiamiento_apropiacion.go)
- [ ] `RegistrarRubro()` en [controllers/rubro.go](controllers/rubro.go)
- [ ] `EliminarRubro()` en [controllers/rubro.go](controllers/rubro.go)

No se recomienda inmediatamente eliminar los controladores anteriores, sino hacer
que retornen un HTTP 424 y actualizar la documentación que genera el Swagger para
indicar que dependen de PLAN_CUENTAS_CRUD, la cual no está en servicio actualmente.
