# PLAN CUENTAS MID
Middleware para modelo de negocio plan_cuentas, el proyecto está escrito en el lenguaje Go, generado mediante el **[framework beego](https://beego.me/)**
***
### Arquitectura del proyecto
![](arquitectura.png)
## Opciones de instalación 
> ### En Ambiente Dockerizado :whale:

**para usar esta opcion es necesario contar con [DOCKER](https://docs.docker.com/) y [DOCKER-COMPOSE](https://docs.docker.com/compose/) en cualquier SO compatible**

- Clonar el proyecto de github y ubicarse en la carpeta del proyecto:
```shell
git clone https://github.com/udistrital/plan_cuentas_mid.git
cd plan_cuentas_mid
```

- Correr el proyecto por docker compose 
1. Crear red de contenedores denominada back_end con el comando (si ya esta creada no es necesario crearla):

```sh
docker network create back_end
```

2. Para construir y correr los contenedores:
```sh
docker-compose up --build
```
- Bajar los servicios de los contenedores
```sh
docker-compose down
```
- Subir los servicios de los contenedores ya construidos previamente
```sh
docker-compose up
```
# Archivos para variables de entorno: 

- para definir puertos, dns y configuraciones internas dentro del archivo **.env**
- para definir conexiones externas a otros apis se debe crear el archivo **custom.env** en la raiz del proyecto

# Licencia
[licencia](LICENSE)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
