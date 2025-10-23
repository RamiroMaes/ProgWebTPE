package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	db "ejemplo.com/mi-proyecto-go/db/sqlc"
)

// POST /paises
func CreatePaisHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p struct {
			Nombre string `json:"nombre"`
		}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(p.Nombre) == "" {
			http.Error(w, "nombre es requerido", http.StatusBadRequest)
			return
		}

		queries := db.New(dbConn)
		created, err := queries.CreatePais(context.Background(), p.Nombre)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"nombre": created})
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
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