# plan_cuentas_mid
Middleware para modelo de negocio plan_cuentas, proporciona a distintos clientes datos  para la gestión de la información del sistemas financiero, la presente api manipula los datos obtenidos de diferentes bases de datos, relacionales y no relacionales.

Clientes que consumen esta api:
- [Presupuesto](https://github.com/udistrital/presupuesto_cliente)
- [Necesidades](https://github.com/udistrital/necesidades_cliente)
- [Contabilidad](https://github.com/udistrital/contabilidad_cliente)

APIs CRUD consumidas por esta api:
- [Plan cunetas mongo](https://github.com/udistrital/plan_cuentas_mongo_crud)
- [Necesidades Crud](https://github.com/udistrital/necesidades_crud)


## Especificaciones Técnicas

### Variables de Entorno

```shell
# Ejemplo que se debe actualizar acorde al proyecto
FINANCIERA_MONGO_CRUD_PORT = [descripción]
FINANCIERA_MONGO_CRUD_DB_URL = [descripción]
```

### Instalación

***Requisitos:***
* [DOCKER](https://docs.docker.com/)
* [DOCKER-COMPOSE](https://docs.docker.com/compose/)

```shell
# Clonar el proyecto de github y ubicarse en la carpeta del proyecto:
git clone https://github.com/udistrital/plan_cuentas_mid.git

# Ingresar al directorio
cd plan_cuentas_mid

#Crear red de contenedores denominada back_end con el comando (si ya esta creada no es necesario crearla):
docker network create back_end

# Construir y correr los contenedores
docker-compose up --build

#Bajar los servicios de los contenedores
docker-compose down

# Subir los servicios de los contenedores ya construidos previamente
docker-compose up
```

### Arquitectura del proyecto
![](arquitectura.png)


## Licencia
[licencia](LICENSE)

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
