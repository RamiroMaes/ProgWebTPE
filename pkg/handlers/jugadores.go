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
	"ejemplo.com/mi-proyecto-go/views"
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	"github.com/lib/pq"
)

// POST /jugadores
func CreateJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error al parsear el formulario: "+err.Error(), http.StatusBadRequest)
			return
		}

		nombre := r.FormValue("agregarNombre")
		posicion := r.FormValue("agregarPosicion")
		paisNombre := r.FormValue("agregarPais")
		idJugador, errID := strconv.Atoi(r.FormValue("agregarNumero"))
		altura, errAltura := strconv.Atoi(r.FormValue("agregarAltura"))
		
        // Para la fecha, el formato de un input type="date" es "YYYY-MM-DD"
		fechaNacimiento, errFecha := time.Parse("2006-01-02", r.FormValue("agregarFechaNacimiento"))
		if errID != nil || errAltura != nil || errFecha != nil {
			http.Error(w, "Error en la conversion de la fecha de nacimiento.", http.StatusBadRequest)
			return
		}

		params := db.CreateJugadorParams{
			Nombre:         nombre,
			IDJugador:      int32(idJugador),
			Posicion:       posicion,
			FechaNacimiento: fechaNacimiento,
			Altura:         int32(altura),
			PaisNombre:     paisNombre,
		}

		// Validación de que todos los campos obligatorios estén presentes
		if strings.TrimSpace(params.Nombre) == "" {
			http.Error(w, "nombre es obligatorio", http.StatusBadRequest)
			return
		}
		if params.Altura <= 0 {
			http.Error(w, "altura es obligatorio y debe ser > 0", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(params.Posicion) == "" {
			http.Error(w, "posición es obligatorio", http.StatusBadRequest)
			return
		}
		if strings.TrimSpace(params.PaisNombre) == "" {
			http.Error(w, "pais_nombre es obligatorio", http.StatusBadRequest)
			return
		}
		if time.Time(params.FechaNacimiento).IsZero() {
			http.Error(w, "fecha_nacimiento es obligatorio", http.StatusBadRequest)
			return
		}
		if params.IDJugador <= 0 {
			http.Error(w, "id_jugador debe ser > 0", http.StatusBadRequest)
			return
		}

		// El código de error estándar de PostgreSQL para una violación de unicidad es 23505.
		// El código de error de PostgreSQL para una violación de clave foránea es 23503.

		queries := db.New(dbConn)
		createdJugador, err := queries.CreateJugador(r.Context(), params)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
                switch pqErr.Code {
                case "23505":
                    http.Error(w, "Error: Ya existe un jugador con ese número.", http.StatusConflict)
                case "23503":
                    if pqErr.Constraint == "jugador_pais" { // Se verifica que el conflicto sea con la tabla de paises y no con otra.
                        http.Error(w, "Error: El país ingresado no es válido.", http.StatusBadRequest)
                    } else {
                        // Para otras violaciones de FK que no sean la de tabla país (en caso de futuras ampliaciones de la bdd).
                        http.Error(w, "Error de datos: "+pqErr.Message, http.StatusBadRequest)
                    }
                default:
                    // Para otros errores de base de datos no capturados específicamente
                    http.Error(w, "Error de base de datos: "+pqErr.Message, http.StatusInternalServerError)
                }
			} else {
                // Si el error no es de tipo *pq.Error
				http.Error(w, "Error interno al crear el jugador: "+err.Error(), http.StatusInternalServerError)
			}
			return // No continuamos.
		}
		w.Header().Set("Content-Type", "text/html")
		views.EntityRow(createdJugador).Render(r.Context(), w)
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

// GET /plantel
func ListPlantelHandler(dbConn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := db.New(dbConn)
		all, err := queries.ListPlantel(context.Background())
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

// DELETE /jugadores/{id}
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
        if err := queries.DeleteJugador(r.Context(), id); err != nil {
            if errors.Is(err, sql.ErrNoRows) {
                http.NotFound(w, r)
                return
            }
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusNoContent)
    }
}