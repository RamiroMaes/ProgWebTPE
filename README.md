## Requisitos Previos

Es requisito obligatorio tener instalados en tu sistema:
*   [Go]https://go.dev/doc/install (versión 1.25.0 o superior)
*   [Docker]https://docs.docker.com/get-docker/
*   [sqlc]https://docs.sqlc.dev/en/latest/overview/install.html

Usando el Makefile:

### 1: Probar la API

* El proyecto incluye un script de prueba para verificar que este funcione correctamente. Este comando configurará la base de datos, iniciará el servidor, ejecutará las pruebas y luego lo detendrá todo.

* Notese que las pruebas son cargadas desde un archivo llamada requests.bash. Este contiene sentencias CURL.

* Ejecutar por consola:
        make test

### 2: Limpieza

* Para detener y eliminar el contenedor de la base de datos creado mediante el comando 'make test', podes usar:
        make clean

### 3: Extras

* En el archivo de Makefile se ven ademas posibles comandos make para realizar instancias temporales de 'make test'