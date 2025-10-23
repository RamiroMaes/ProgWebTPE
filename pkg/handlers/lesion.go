package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	db "ejemplo.com/mi-proyecto-go/db/sqlc"
)

// POST /lesiones
func CreateLesionHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p db.CreateLesionParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Validación mejorada
		if strings.TrimSpace(p.TipoLesion) == "" || strings.TrimSpace(p.Descripcion) == "" {
			http.Error(w, "tipo_lesion y descripcion son requeridos", http.StatusBadRequest)
			return
		}

		queries := db.New(dbConn)
		created, err := queries.CreateLesion(context.Background(), p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

// GET /lesiones
func ListLesionesHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListLesiones(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
	}
}

// DELETE /lesiones/{tipo_lesion}
func DeleteLesionHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tipoLesion := r.PathValue("tipo_lesion")
		queries := db.New(dbConn)
		err := queries.DeleteLesion(context.Background(), tipoLesion)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}