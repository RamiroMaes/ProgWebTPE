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
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	"github.com/lib/pq"
)

// POST /jugadores
func CreateJugadorHandler(dbConn *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
            return
        }

        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
            return
        }

        numero, err1 := strconv.Atoi(r.FormValue("agregarNumero"))
        altura, err2 := strconv.Atoi(r.FormValue("agregarAltura"))
        fechaNacimiento, err3 := time.Parse("2006-01-02", r.FormValue("agregarFechaNacimiento"))

        if err1 != nil || err2 != nil || err3 != nil {
            http.Error(w, "Error en la conversión de campos numéricos o fecha.", http.StatusBadRequest)
            return
        }

        params := db.CreateJugadorParams{
            Nombre:          r.FormValue("agregarNombre"),
            IDJugador:       int32(numero),
            Posicion:        r.FormValue("agregarPosicion"),
            FechaNacimiento: fechaNacimiento,
            Altura:          int32(altura),
            PaisNombre:      r.FormValue("agregarPais"),
        }

        if strings.TrimSpace(params.Nombre) == "" {
            http.Error(w, "Nombre es obligatorio", http.StatusBadRequest)
            return
        }

        queries := db.New(dbConn)
        _, err := queries.CreateJugador(r.Context(), params)
        if err != nil {
            if pqErr, ok := err.(*pq.Error); ok {
                switch pqErr.Code {
                case "23505":
                    http.Error(w, "Error: Ya existe un jugador con ese número.", http.StatusConflict)
                case "23503":
                    http.Error(w, "Error: El país ingresado no es válido.", http.StatusBadRequest)
                default:
                    http.Error(w, "Error de base de datos: "+pqErr.Message, http.StatusInternalServerError)
                }
            } else {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
            return
        }

        http.Redirect(w, r, "/#tabla-jugadores", http.StatusSeeOther)	// Redirige a la sección de la tabla de jugadores
    }
}


// GET /jugadores
	// Devuelve texto plano con solo id y nombre de cada jugador
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
	// Devuelve texto plano con todos los datos de cada jugador
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
	// Devuelve texto plano con todos los datos del jugador solicitado
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
        // Se obtiene ID de la URL
        idStr := r.PathValue("id")
        id64, err := strconv.ParseInt(idStr, 10, 32)
        if err != nil {
            http.Error(w, "id inválido en la URL", http.StatusBadRequest)
            return
        }
        id := int32(id64)

        // Se parsea el formulario
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error al leer el formulario: "+err.Error(), http.StatusBadRequest)
            return
        }

        nombre := r.FormValue("nombre")
        posicion := r.FormValue("posicion")
        paisNombre := r.FormValue("pais")
        
        altura, errAlt := strconv.Atoi(r.FormValue("altura"))
        fechaNacimiento, errFecha := time.Parse("2006-01-02", r.FormValue("fecha_nacimiento"))

        if errAlt != nil || errFecha != nil {
            http.Error(w, "Error en formato de altura (número) o fecha (YYYY-MM-DD)", http.StatusBadRequest)
            return
        }

        params := db.UpdateJugadorParams{
            IDJugador:       id,
            Nombre:          nombre,
            Posicion:        posicion,
            PaisNombre:      paisNombre,
            Altura:          int32(altura),
            FechaNacimiento: fechaNacimiento,
        }

        queries := db.New(dbConn)
        if err := queries.UpdateJugador(r.Context(), params); err != nil {
            http.Error(w, "Error al actualizar: "+err.Error(), http.StatusInternalServerError)
            return
        }

        updated, _ := queries.GetJugador(r.Context(), id)
        
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Jugador actualizado: %v", updated) // %v imprime el struct
    }
}


// DELETE /jugadores/{id}
	// Elimina un jugador por ID (pensado para hacerlo desde CLI)
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


// POST /jugadores/{id}
	// Elimina un jugador desde el botón de la fila del jugador
func DeleteBotonJugador(queries *db.Queries) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
            return
        }

        if err := r.ParseForm(); err != nil {
            http.Error(w, "Error al parsear el formulario", http.StatusBadRequest)
            return
        }

        idStr := r.FormValue("id")
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "ID inválido", http.StatusBadRequest)
            return
        }

        err = queries.DeleteJugador(r.Context(), int32(id))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}