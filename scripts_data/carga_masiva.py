import csv
import requests 
  

# Se abre el archivo CSV u se imprimen los registros del mismo.

with open('datos.csv', 'r') as csvFile:
    reader = csv.reader(csvFile)
    for row in reader:
        print(row)
csvFile.close()

# Procesamiento para dar formato al árbol

# Diccionario de datos con la estructura del arbol de rubros que será enviada para el POST
data = {'CODIGO' : CODIGO,
        'NOMBRE' : NOMBRE,
        'ALGO' : 'ALGO'
        }

# Endpoints
PLAN_CUENTAS_CRUD = "http://maps.googleapis.com/maps/api/geocode/json"
PLAN_CUENTAS_MONGO_CRUD = "http://maps.googleapis.com/maps/api/geocode/json"

r_crud = requests.post(url = PLAN_CUENTAS_CRUD, data = data) 
r_mongo_crud = requests.post(url = PLAN_CUENTAS_MONGO_CRUD, data = data)