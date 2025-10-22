package handlers

import (
)

func JugadoresHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getJugadores(w)
    case http.MethodPost:
        createJugador(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func getJugadores(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(repository.GetAll())
}

func createJugador(w http.ResponseWriter, r *http.Request) {
    var newJugador db.sqlc.Jugador

    if err := json.NewDecoder(r.Body).Decode(&newJugador); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := logic.ValidateJugador(newJugador); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	createdJugador, err := queries.CreateJugador(ctx, db.sqlc.createJugadorParams{
		fmt.Printf("Jugador creado: %+v\n", createdJugador)
	}
}