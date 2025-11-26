package handlers

import (
	"net/http"
    "sort"
	
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	"ejemplo.com/mi-proyecto-go/views"
)

var posicionOrden = map[string]int{
    "Arquero":       1,
    "Defensor":      2,
    "Mediocampista": 3,
    "Delantero":     4,
}

func ordenarJugadores(jugadores []db.Jugador, sortColumn string) {
    switch sortColumn {
    case "posicion":
        sort.Slice(jugadores, func(i, j int) bool {
            return posicionOrden[jugadores[i].Posicion] < posicionOrden[jugadores[j].Posicion]
        })
    case "id_jugador":
        sort.Slice(jugadores, func(i, j int) bool {
            return jugadores[i].IDJugador < jugadores[j].IDJugador
        })
    case "nombre":
        sort.Slice(jugadores, func(i, j int) bool {
            return jugadores[i].Nombre < jugadores[j].Nombre
        })
    case "pais_nombre":
        sort.Slice(jugadores, func(i, j int) bool {
            return jugadores[i].PaisNombre < jugadores[j].PaisNombre
        })
    case "fecha_nacimiento":
        sort.Slice(jugadores, func(i, j int) bool {
            return jugadores[i].FechaNacimiento.Before(jugadores[j].FechaNacimiento)
        })
    case "altura":
        sort.Slice(jugadores, func(i, j int) bool {
            return jugadores[i].Altura < jugadores[j].Altura
        })
    }
}

func ListJugadoresPage(queries *db.Queries) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Nos aseguramos de que este handler solo responda a la ruta raíz exacta.
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        
        // Se obtienen los jugadores desde la base de datos.
        jugadores, err := queries.ListPlantel(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Se obtienen los países desde la base de datos para el datalist en el formulario
        paises, err := queries.ListPaises(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Renderiza el layout completo con la lista y el formulario.
        views.Layout(views.EntityList(jugadores, ""), views.EntityForm(paises)).Render(r.Context(), w)
    }
}

func SortJugadoresHandler(queries *db.Queries) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        sortColumn := r.URL.Query().Get("sort")
        
        // Validar columna
        validColumns := map[string]bool{
            "posicion": true, "id_jugador": true, "nombre": true,
            "pais_nombre": true, "fecha_nacimiento": true, "altura": true,
        }
        
        if !validColumns[sortColumn] {
            sortColumn = "id_jugador"
        }
        
        jugadores, err := queries.ListPlantel(r.Context())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        ordenarJugadores(jugadores, sortColumn)
        
        w.Header().Set("Content-Type", "text/html")
        views.EntityList(jugadores, sortColumn).Render(r.Context(), w)
    }
}