Para poder ejecutar se debera:
    -Tener el Lenguaje GO instalado e tu maquina.
    -Inicializar el modulo 'go' ingresando via la terminal del proyecto: "go mod init prueba.com/mi-proyecto" para crear asi el archivo go.mod si este no esta incluido.
    -Inicializar el servidor, ingresando via una terminal "go run ." esta debera estar situada en el directorio del proyecto.
    -Se mostrara: Servidor escuchando en http://localhost:8080. Se ingresa el link en un navegador web para poder visualizar la resolucion.


La estructura del proyecto esta conformada por un archivo main.go, encargado de tener los 'handler' de las rutas elegidas, sirviendo los archivos '.html' correspondientes a cada una. Define el puerto e inicializa el servidor HTTP. La estructura de la ruta jugador apunta a contener la informacion de cada uno encontrada en la base de datos para los practicos siguientes. La ruta raiz busca mostrar el listado de los jugadores del Plantel Profesional de River Plate.