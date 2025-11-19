package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	db "ejemplo.com/mi-proyecto-go/db/sqlc"
)

// POST /paises
func CreatePaisHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nombre := r.FormValue("nombre")

		if strings.TrimSpace(nombre) == "" {
			http.Error(w, "nombre es obligatorio", http.StatusBadRequest)
			return
		}

		queries := db.New(dbConn)

		created, err := queries.CreatePais(r.Context(), nombre)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

		// La respuesta es texto plano y no mas JSON.
		w.WriteHeader(http.StatusCreated)
        w.Write([]byte("País creado: " + created))
	}
		
}

// GET /paises
func ListPaisesHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListPaises(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
		for _, pais := range all {
            fmt.Fprintf(w, "%v | ", pais) // %v imprime el struct
        }
	}
}

// DELETE /paises/{nombre}
func DeletePaisHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nombre := r.PathValue("nombre")
		queries := db.New(dbConn)
		err := queries.DeletePais(context.Background(), nombre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}