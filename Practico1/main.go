package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Registramos un handler para la ruta raíz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 - Página no encontrada")
			return
		}
		//Establecemos la cabecera Content-Type
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//Servimos el archivo index.html
		http.ServeFile(w, r, "index.html")
	})

	//Registramos un manejador para la ruta /jugador
	http.HandleFunc("/jugador", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "jugador.html")
	})

	//Definimos el puerto y muestra un mensaje
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	//Iniciamos el servidor HTTP
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
