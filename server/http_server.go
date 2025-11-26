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

	// --- PAISES ---
	mux.HandleFunc("GET /paises", h.ListPaisesHandler(dbConn))
	mux.HandleFunc("POST /paises", h.CreatePaisHandler(dbConn))
	mux.HandleFunc("DELETE /paises/{nombre}", h.DeletePaisHandler(dbConn))

	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	q := db.New(dbConn)

	// --- JUGADORES ---
	mux.HandleFunc("GET /", h.ListJugadoresPage(q))
	mux.HandleFunc("POST /jugadores", h.CreateJugadorHandler(dbConn))

	mux.HandleFunc("GET /jugadores", h.GetJugadoresHandler(q, dbConn))
	mux.HandleFunc("GET /plantel", h.ListPlantelHandler(dbConn))
	
	mux.HandleFunc("GET /jugadores/{id}", h.GetJugadorHandler(dbConn))
	mux.HandleFunc("PUT /jugadores/{id}", h.UpdateJugadorHandler(dbConn))
	mux.HandleFunc("DELETE /jugadores/{id}", h.DeleteJugadorHandler(dbConn))

	// Inicia el servidor
	log.Printf("Servidor escuchando en %s\n", addr)
	return http.ListenAndServe(addr, mux)
}
