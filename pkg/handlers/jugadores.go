package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
	"fmt"
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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		for _, jugador := range all {
            fmt.Fprintf(w, "%v, ", jugador) // %v imprime el struct
        }
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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		for _, jugador := range all {
            fmt.Fprintf(w, "%v, ", jugador) // %v imprime el struct
        }
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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", jugador) // %v imprime el struct
	}
}


// PUT /jugadores/{id}
func UpdateJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 1. Obtener ID de la URL
        idStr := r.PathValue("id")
        id64, err := strconv.ParseInt(idStr, 10, 32)
        if err != nil {
            http.Error(w, "id inválido en la URL", http.StatusBadRequest)
            return
        }
        id := int32(id64)

        // 2. Parsear el formulario (x-www-form-urlencoded)
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error al leer el formulario: "+err.Error(), http.StatusBadRequest)
            return
        }

        // 3. Obtener y convertir valores
        // Asumimos nombres de campos estándar: nombre, posicion, pais, altura, fecha_nacimiento
        nombre := r.FormValue("nombre")
        posicion := r.FormValue("posicion")
        paisNombre := r.FormValue("pais")
        
        altura, errAlt := strconv.Atoi(r.FormValue("altura"))
        fechaNacimiento, errFecha := time.Parse("2006-01-02", r.FormValue("fecha_nacimiento"))

        if errAlt != nil || errFecha != nil {
            http.Error(w, "Error en formato de altura (número) o fecha (YYYY-MM-DD)", http.StatusBadRequest)
            return
        }

        // 4. Construir parámetros
        params := db.UpdateJugadorParams{
            IDJugador:       id,
            Nombre:          nombre,
            Posicion:        posicion,
            PaisNombre:      paisNombre,
            Altura:          int32(altura),
            FechaNacimiento: fechaNacimiento,
        }

        // 5. Ejecutar Update en BD
        queries := db.New(dbConn)
        if err := queries.UpdateJugador(r.Context(), params); err != nil {
            http.Error(w, "Error al actualizar: "+err.Error(), http.StatusInternalServerError)
            return
        }

        // 6. Responder con Texto Plano
        updated, _ := queries.GetJugador(r.Context(), id)
        
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        // Devuelve la representación en texto del struct actualizado
        fmt.Fprintf(w, "Jugador actualizado: %v", updated)
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