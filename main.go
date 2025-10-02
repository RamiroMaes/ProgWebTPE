package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	sqlc "ejemplo.com/mi-proyecto-go/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres password=XYZ dbname=tpespecial sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}
	fmt.Println("¡Conexión exitosa a la base de datos!")

	queries := sqlc.New(db)
	ctx := context.Background()

	// --- 1. PREPARACIÓN: Crear Países ---
	fmt.Println("\n--- Creando Países de prueba... ---")
	paises := []string{"Argentina", "Francia"}
	for _, p := range paises {
		_, err := queries.CreatePais(ctx, p)
		if err != nil {
			fmt.Printf("Advertencia al crear país '%s' (puede que ya existiera): %v\n", p, err)
		} else {
			fmt.Printf("País '%s' creado o ya existente.\n", p)
		}
	}

	// --- 2. CREAR un nuevo jugador ---
	fmt.Println("\n--- Creando un jugador... ---")
	fechaNac := time.Date(1987, time.June, 24, 0, 0, 0, 0, time.UTC)
	createJugadorParams := sqlc.CreateJugadorParams{
		Nombre:          "Lionel Messi",
		IDJugador:       10,
		Posicion:        "Delantero",
		FechaNacimiento: fechaNac,
		Altura:          "1.70",
		PaisNombre:      "Argentina",
	}

	createdJugador, err := queries.CreateJugador(ctx, createJugadorParams)
	if err != nil {
		log.Fatalf("Error al crear jugador: %v", err)
	}
	fmt.Printf("Jugador creado: %+v\n", createdJugador)

	// --- 3. OBTENER el jugador recién creado ---
	fmt.Println("\n--- Obteniendo el jugador por ID... ---")
	fetchedJugador, err := queries.GetJugador(ctx, createdJugador.IDJugador)
	if err != nil {
		log.Fatalf("Error al obtener jugador: %v", err)
	}
	fmt.Printf("Jugador obtenido: %+v\n", fetchedJugador)

	// --- 4. ACTUALIZAR la posición del jugador ---
	fmt.Println("\n--- Actualizando el jugador... ---")
	// La firma correcta es (Posicion, FechaNacimiento, Altura, PaisNombre, IDJugador, Nombre)
	err = queries.UpdateJugador(ctx, sqlc.UpdateJugadorParams{
		Posicion:        "Mediapunta", // Nueva posición
		FechaNacimiento: createdJugador.FechaNacimiento,
		Altura:          createdJugador.Altura,
		PaisNombre:      createdJugador.PaisNombre,
		IDJugador:       createdJugador.IDJugador, // WHERE
		Nombre:          createdJugador.Nombre,    // WHERE
	})
	if err != nil {
		log.Fatalf("Error al actualizar jugador: %v", err)
	}
	fmt.Println("Jugador actualizado. Verificando el cambio...")

	// Verificamos el cambio obteniendo el jugador de nuevo
	updatedJugador, err := queries.GetJugador(ctx, createdJugador.IDJugador)
	if err != nil {
		log.Fatalf("Error al obtener jugador actualizado: %v", err)
	}
	fmt.Printf("Datos después de actualizar: Posición -> %s\n", updatedJugador.Posicion)

	// --- 5. LISTAR todos los jugadores ---
	fmt.Println("\n--- Listando todos los jugadores... ---")
	jugadores, err := queries.ListJugadores(ctx)
	if err != nil {
		log.Fatalf("Error al listar jugadores: %v", err)
	}
	fmt.Printf("Jugadores encontrados (%d):\n", len(jugadores))
	for _, jugador := range jugadores {
		fmt.Printf(" - ID: %d, Nombre: %s, Posición: %s\n",
			jugador.IDJugador, jugador.Nombre, jugador.Posicion)
	}

	/*
		// --- 6. BORRAR el jugador ---
		fmt.Println("\n--- Borrando el jugador... ---")
		err = queries.DeleteJugador(ctx, sqlc.DeleteJugadorParams{
			IDJugador: createdJugador.IDJugador,
			Nombre:    createdJugador.Nombre,
		})
		if err != nil {
			log.Fatalf("Error al borrar jugador: %v", err)
		}
		fmt.Println("Jugador borrado exitosamente.")
	*/
	// --- 7. Probar CRUD para Club ---
	fmt.Println("\n--- Probando CRUD para Club... ---")
	club, err := queries.CreateClub(ctx, sqlc.CreateClubParams{Nombre: "River Plate", Ciudad: "Buenos Aires"})
	if err != nil {
		fmt.Printf("Advertencia al crear club (puede que ya existiera): %v\n", err)
	} else {
		fmt.Printf("Club creado: %+v\n", club)
	}

	// Listar clubs
	clubs, err := queries.ListClubs(ctx)
	if err != nil {
		log.Fatalf("Error al listar clubs: %v", err)
	}
	fmt.Printf("Clubs encontrados (%d):\n", len(clubs))
	for _, c := range clubs {
		fmt.Printf(" - Nombre: %s, Ciudad: %s\n", c.Nombre, c.Ciudad)
	}

	/*
	   // Borrar club
	   err = queries.DeleteClub(ctx, sqlc.DeleteClubParams{Nombre: "River Plate", Ciudad: "Buenos Aires"})

	   	if err != nil {
	   		log.Fatalf("Error al borrar club: %v", err)
	   	}

	   fmt.Println("Club 'River Plate' borrado.")
	*/
}
