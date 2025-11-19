package handlers

import (
	"net/http"
	
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	"ejemplo.com/mi-proyecto-go/views"
)

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
        views.Layout(views.EntityList(jugadores), views.EntityForm(paises)).Render(r.Context(), w)
    }
}