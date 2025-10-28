package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	db "ejemplo.com/mi-proyecto-go/db/sqlc"
)

// POST /jugadores
func CreateJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p db.CreateJugadorParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "invalid json body: "+err.Error(), http.StatusBadRequest)
			return
		}
		
		// Validación de que todos los campos obligatorios estén presentes
		if strings.TrimSpace(p.Nombre) == "" {
			http.Error(w, "nombre es obligatorio", http.StatusBadRequest)
			return
		}
		if p.Altura <= 0 {
			http.Error(w, "altura es obligatorio y debe ser > 0", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(p.Posicion) == "" {
			http.Error(w, "posición es obligatorio", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(p.PaisNombre) == "" {
			http.Error(w, "pais_nombre es obligatorio", http.StatusBadRequest)
			return
		}
		if time.Time(p.FechaNacimiento).IsZero() {
			http.Error(w, "fecha_nacimiento es obligatorio", http.StatusBadRequest)
			return
		}
		if p.IDJugador <= 0 {
			http.Error(w, "id_jugador debe ser > 0", http.StatusBadRequest)
			return
		}
		
		// Lógica de DB
		queries := db.New(dbConn)
		created, err := queries.CreateJugador(context.Background(), p)
		if err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

// GET /jugadores
func ListJugadoresHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListJugadores(context.Background())
		if err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
	}
}

// GET /jugadoresCompleto
func ListJugadoresCompletoHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListJugadoresCompleto(context.Background())
		if err != nil {
			http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(all)
	}
}


// GET /jugadores/{id}
func GetJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id64, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, "id inválido en la URL", http.StatusBadRequest)
			return
		}
		id := int32(id64)

		// Lógica de DB
		queries := db.New(dbConn)
		jugador, err := queries.GetJugador(context.Background(), id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				http.NotFound(w, r)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jugador)
	}
}

// PUT /jugadores/{id}
func UpdateJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id64, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, "id inválido en la URL", http.StatusBadRequest)
			return
		}
		id := int32(id64)
		var p db.UpdateJugadorParams
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, "invalid json body: "+err.Error(), http.StatusBadRequest)
			return
		}
		
		// Validación. mismo que antes
		if p.IDJugador != id {
			http.Error(w, "id en el cuerpo debe coincidir con el id en la ruta", http.StatusBadRequest)
			return
		}

		// Lógica de DB
		queries := db.New(dbConn)
		if err := queries.UpdateJugador(context.Background(), p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		updated, _ := queries.GetJugador(context.Background(), id)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updated)
	}
}

// DELETE /jugadores/{id}/{nombre}
func DeleteJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		
		id64, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		
		id := int32(id64)

		queries := db.New(dbConn)
		err = queries.DeleteJugador(context.Background(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		w.WriteHeader(http.StatusNoContent)
	}
}