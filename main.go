package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	// OJO: Reemplaza "PROGWEBTPE" con el nombre de tu módulo (ver go.mod)
	sqlc "ejemplo.com/mi-proyecto-go/db/sqlc"
)

func main() {
	// OJO: Reemplaza estos valores con tus credenciales reales.
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

	// --- 1. PREPARACIÓN: Crear un País (debido a la clave foránea) ---
	// El jugador que vamos a crear es de "Francia", así que esa fila debe existir primero.
	fmt.Println("\n--- Creando País de prueba... ---")
	_, err = queries.CreatePais(ctx, "Francia")
	if err != nil {
		// No usamos log.Fatalf para que no se detenga si el país ya existe
		fmt.Printf("Advertencia al crear país (puede que ya existiera): %v\n", err)
	} else {
		fmt.Println("País 'Francia' creado o ya existente.")
	}

	// --- 2. CREAR un nuevo jugador ---
	fmt.Println("\n--- Creando un jugador... ---")

	// Usamos time.Date para crear la fecha de nacimiento
	fechaNac := time.Date(1987, time.June, 24, 0, 0, 0, 0, time.UTC)

	createJugadorParams := sqlc.CreateJugadorParams{
		Nombre:          "Lionel",
		IDJugador:       10, // Corresponde a iD_Jugador
		Posicion:        "Delantero",
		FechaNacimiento: fechaNac,
		Altura:          "1.70", // sqlc usa string para el tipo DECIMAL para mantener la precisión
		PaisNombre:      "Argentina",
	}

	createdJugador, err := queries.CreateJugador(ctx, createJugadorParams)
	if err != nil {
		log.Fatalf("Error al crear jugador: %v", err)
	}
	fmt.Printf("Jugador creado: %+v\n", createdJugador)

	// --- 3. LISTAR todos los jugadores ---
	fmt.Println("\n--- Listando todos los jugadores... ---")
	jugadores, err := queries.ListJugadores(ctx)
	if err != nil {
		log.Fatalf("Error al listar jugadores: %v", err)
	}
	fmt.Printf("Jugadores encontrados (%d):\n", len(jugadores))
	for _, jugador := range jugadores {
		fmt.Printf(" - ID: %d, Nombre: %s, Posición: %s, Altura: %s\n",
			jugador.IDJugador, jugador.Nombre, jugador.Posicion, jugador.Altura)
	}

	// --- 4. ACTUALIZAR la posición del jugador creado ---
	fmt.Println("\n--- Actualizando el jugador... ---")
	err = queries.UpdateJugador(ctx, sqlc.UpdateJugadorParams{
		Nombre:          createdJugador.Nombre,
		IDJugador:       createdJugador.IDJugador,
		FechaNacimiento: createdJugador.FechaNacimiento,
		Altura:          createdJugador.Altura,
		PaisNombre:      createdJugador.PaisNombre,
		Posicion:        "Mediapunta", // Nueva posición
	})
	if err != nil {
		log.Fatalf("Error al actualizar jugador: %v", err)
	}
	fmt.Printf("Jugador actualizado")

	// --- 5. BORRAR el jugador ---
	fmt.Println("\n--- Borrando el jugador... ---")
	err = queries.DeleteJugador(ctx, sqlc.DeleteJugadorParams{
		Nombre:    createdJugador.Nombre,
		IDJugador: createdJugador.IDJugador,
	})
	if err != nil {
		log.Fatalf("Error al borrar jugador: %v", err)
	}
	fmt.Println("Jugador borrado exitosamente.")
}
