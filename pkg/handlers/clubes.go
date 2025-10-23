package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	db "ejemplo.com/mi-proyecto-go/db/sqlc"
)

// POST /clubs
func CreateClubHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p db.CreateClubParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(p.Nombre) == "" || strings.TrimSpace(p.Ciudad) == "" {
			http.Error(w, "nombre y ciudad son requeridos", http.StatusBadRequest)
			return
		}

		queries := db.New(dbConn)
		created, err := queries.CreateClub(context.Background(), p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

// GET /clubs
func ListClubsHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListClubs(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
	}
}

// GET /clubs/{nombre}/{ciudad}
func GetClubHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		nombre := r.PathValue("nombre")
		ciudad := r.PathValue("ciudad")

		queries := db.New(dbConn)
		c, err := queries.GetClub(context.Background(), db.GetClubParams{
			Nombre: nombre,
			Ciudad: ciudad,
		})
		if err != nil {
			if err == sql.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(c)
	}
}

// PUT /clubs/{nombre}/{ciudad}
func UpdateClubHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// --- CÓDIGO SIMPLIFICADO ---
		nombreViejo := r.PathValue("nombre")
		ciudadVieja := r.PathValue("ciudad")

		var p db.UpdateClubParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Asignamos los valores de la URL a los parámetros de la query
		p.Nombre_2 = nombreViejo
		p.Ciudad_2 = ciudadVieja

		queries := db.New(dbConn)
		if err := queries.UpdateClub(context.Background(), p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p) // Devuelve el body que se usó para actualizar
	}
}

// DELETE /clubs/{nombre}/{ciudad}
func DeleteClubHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nombre := r.PathValue("nombre")
		ciudad := r.PathValue("ciudad")

		queries := db.New(dbConn)
		err := queries.DeleteClub(context.Background(), db.DeleteClubParams{
			Nombre: nombre,
			Ciudad: ciudad,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusNoContent)
	}
}