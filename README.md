# ProgWebTPE — Guía rápida de ejecución
# Requisitos
    - Linux, bash y make
    - Docker
    - Go 1.22+ (o compatible)
    - sqlc (solo en el caso de que se quieran modificar las queries)
    - templ (solo en el caso de que se quiera modificar algo del DOM)

# Variables (por defecto en el Makefile)
- Contenedor: `some-postgres`
- Usuario DB: `postgres`
- Password: `XYZ`
- DB: `tpespecial`
- Imagen: `postgres:16`
- Puerto HTTP: `8080`

Se podrian sobreescribir al invocar make, por ejemplo:
`make DB_PASSWORD=otra PASS DB_NAME=otra_db setup`




# PASOS:


# 1- Preparar la base de datos (2 opciones) mediante consola linux:
Crear o arrancar el contenedor, esperar disponibilidad, crear la DB si falta y aplicar el schema.
```bash
make setup
```

sino mas simple (tambien se levanta el servidor y se ejecuta el requests.bash):
```bash
    make test
```

## 2- Levantar el servidor:
En una terminal:
```bash
go run ./cmd/server
```
La API quedará en `http://localhost:8080`.

## 3- Probar la API (2 opciones):
    - Ejecutar el script de peticiones:
```bash
        chmod +x ./requests.bash
        ./requests.bash
```
    - O usar la tarea que arranca server en background, corre el script y lo apaga:
```bash
    make test
```

## 4 - Ver resolucion: 
    Abre en el navegador:

    http://localhost:8080/

## 5- Limpieza de puerto y contenedor
    Detener y eliminar el contenedor:
```     bash
        make clean
    ```
