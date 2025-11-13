package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	h "ejemplo.com/mi-proyecto-go/pkg/handlers"
	_ "github.com/lib/pq"
)

// Inicia el servidor HTTP
func StartServer(connStr string, addr string) error {
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("db open: %w", err)
	}
	if err := dbConn.Ping(); err != nil {
		return fmt.Errorf("db ping: %w", err)
	}
	fmt.Println("HTTP server: conectado a la base de datos")

	// Usamos el mux estándar de Go 1.22+
	mux := http.NewServeMux()

	// --- JUGADORES ---
	mux.HandleFunc("GET /jugadores", h.ListJugadoresHandler(dbConn))
	mux.HandleFunc("GET /plantel", h.ListPlantelHandler(dbConn))
	mux.HandleFunc("POST /jugadores", h.CreateJugadorHandler(dbConn))
	mux.HandleFunc("GET /jugadores/{id}", h.GetJugadorHandler(dbConn))
	mux.HandleFunc("PUT /jugadores/{id}", h.UpdateJugadorHandler(dbConn))
	mux.HandleFunc("DELETE /jugadores/{id}", h.DeleteJugadorHandler(dbConn))

	// --- CLUBS ---
	mux.HandleFunc("GET /clubs", h.ListClubsHandler(dbConn))
	mux.HandleFunc("POST /clubs", h.CreateClubHandler(dbConn))
	mux.HandleFunc("GET /clubs/{nombre}/{ciudad}", h.GetClubHandler(dbConn))
	mux.HandleFunc("PUT /clubs/{nombre}/{ciudad}", h.UpdateClubHandler(dbConn))
	mux.HandleFunc("DELETE /clubs/{nombre}/{ciudad}", h.DeleteClubHandler(dbConn))

	// --- PAISES ---
	mux.HandleFunc("GET /paises", h.ListPaisesHandler(dbConn))
	mux.HandleFunc("POST /paises", h.CreatePaisHandler(dbConn))
	mux.HandleFunc("DELETE /paises/{nombre}", h.DeletePaisHandler(dbConn))

	// --- LESIONES ---
	mux.HandleFunc("GET /lesiones", h.ListLesionesHandler(dbConn))
	mux.HandleFunc("POST /lesiones", h.CreateLesionHandler(dbConn))
	mux.HandleFunc("DELETE /lesiones/{tipo_lesion}", h.DeleteLesionHandler(dbConn))

	// --- RELACIONES ---
	// mux.HandleFunc("GET /jugadores/{id}/clubs", h.ListJugosForJugadorHandler(dbConn))
	// mux.HandleFunc("POST /jugadores/{id}/clubs", h.CreateJugoHandler(dbConn))
	// mux.HandleFunc("GET /jugadores/{id}/lesiones", h.ListLesionesForJugadorHandler(dbConn))
	// mux.HandleFunc("POST /jugadores/{id}/lesiones", h.CreateTieneHandler(dbConn))

	/*
	// --- ARCHIVOS ESTÁTICOS ---
	// Servir index.html en la raíz exacta
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	// Servir carpeta de imágenes
    mux.Handle("GET /img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	// Servir otros archivos estáticos (JS, CSS, etc.)
	mux.HandleFunc("GET /{file}", func(w http.ResponseWriter, r *http.Request) {
		file := r.PathValue("file")
		http.ServeFile(w, r, "./"+file)
	})
	*/

    mux.Handle("GET /img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))

	q := db.New(dbConn)
    JugadorHandler := h.NewJugadoresHandler(q)


	mux.Handle("GET /", JugadorHandler)

	// Inicia el servidor
	log.Printf("Servidor escuchando en %s\n", addr)
	return http.ListenAndServe(addr, mux)
}
