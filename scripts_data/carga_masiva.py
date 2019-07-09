import csv
import requests 
import json
  

# Se abre el archivo CSV u se imprimen los registros del mismo.
PLAN_CUENTAS_MID = "https://localhost:8080/v1/rubro/RegistrarRubro/"
with open('rubros.csv', 'r') as csvFile:
    reader = csv.reader(csvFile)
    for row in reader:
        arreglo=row[0].split(';')
        data = {'Id' : 1,
        'Organizacion' : arreglo[0],
        'Codigo' : arreglo[3],
        'Descripcion' : arreglo[4],
        'UnidadEjecutora' : arreglo[1],
        'Nombre' : arreglo[4]
        }
        #print(data)
        r_mid = requests.post(url = PLAN_CUENTAS_MID, data = json.dumps(data))
        response_mid = r_mid.text 
        status_mid = r_mid.status_code
        print(response_mid)
        print(status_mid)

csvFile.close()

# Procesamiento para dar formato al árbol

# Diccionario de datos con la estructura del arbol de rubros que será enviada para el POST
data = {'Id' : 1,
        'Organizacion' : 1,
        'Codigo' : '1',
        'Descripcion' : 'AA',
        'UnidadEjecutora' : 1,
        'Nombre' : 'AA'
        }

# Endpoints
# PLAN_CUENTAS_CRUD = "https://"
# PLAN_CUENTAS_MONGO_CRUD = "https://"

# r_crud = requests.post(url = PLAN_CUENTAS_CRUD, data = json.dumps(data)) 
# r_mongo_crud = requests.post(url = PLAN_CUENTAS_MONGO_CRUD, data = data)

# JSON Respuesta
# response_crud = r_crud.text
# response_mongo_crud = r_mongo_crud.text

# Status response
# status_crud = r_crud.status_code
# status_mongo_crud = r_mongo_crud.status_code