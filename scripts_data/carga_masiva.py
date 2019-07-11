import csv
import requests 
import json
  

# Se abre el archivo CSV u se imprimen los registros del mismo.
PLAN_CUENTAS_MID = "http://localhost:8080/v1/rubro/"
with open('rubros.csv', 'r') as csvFile:
    reader = csv.reader(csvFile)
    for row in reader:
        arreglo=row[0].split(';')

        # Diccionario de datos con la estructura del arbol de rubros que será enviada para el POST
        data = {'Id' : 1,
        'Organizacion' : arreglo[0],
        'Codigo' : arreglo[3],
        'Descripcion' : arreglo[4],
        'UnidadEjecutora' : arreglo[1],
        'Nombre' : arreglo[4]
        }
        
        r_mid = requests.post(url = PLAN_CUENTAS_MID, data = json.dumps(data))
        
        response_mid = r_mid.text 
        status_mid = r_mid.status_code
        print(response_mid)
        print(status_mid)
        
csvFile.close()



