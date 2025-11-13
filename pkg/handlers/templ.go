package handlers

import (
	"net/http"
	
	db "ejemplo.com/mi-proyecto-go/db/sqlc"
	"ejemplo.com/mi-proyecto-go/views"
)

type JugadoresHandler struct {
	queries *db.Queries
}

func NewJugadoresHandler(q *db.Queries) *JugadoresHandler {
	return &JugadoresHandler{queries: q}
}

func (h *JugadoresHandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.ListJugadores(w, r)
	default:
		http.NotFound(w, r)
	}
}


func (h *JugadoresHandler) ListJugadores(w http.ResponseWriter, r *http.Request) {
	jugadores, err := h.queries.ListJugadores(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	views.Layout(views.EntityList(jugadores), views.EntityForm()).Render(r.Context(), w)
}