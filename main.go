package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Define el contenido HTML para la ruta principal
	htmlContent := `<!DOCTYPE html><html><head>
	 <title>Plantilla River Plate</title></head><body>
	 <h1>Plantilla River Plate</h1>
	 <p>A continuación se muestra la plantilla del equipo de fútbol River Plate:</p></body>
	 </html>`
	// 2. Define el contenido HTML para /about
	aboutContent := `<!DOCTYPE html><html><head>
	 <title>Acerca de</title></head><body>
	 <h1>Info del Servidor</h1>
	 <p>Servidor HTTP ompleto</p></body></html>`

	// 3. Registra un manejador (handler) para la ruta raíz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 - Página no encontrada")
			return
		}
		// 4. Establece la cabecera Content-Type
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 5. Escribe el HTML en la respuesta
		fmt.Fprint(w, htmlContent)
	})

	// 6. Registra un manejador para la ruta /about
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, aboutContent)
	})

	// 8. Define el puerto y muestra un mensaje
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// 9. Inicia el servidor HTTP
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
