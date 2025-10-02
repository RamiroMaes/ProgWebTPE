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
	fmt.Println("Conexión exitosa a la base de datos")

	queries := sqlc.New(db)
	ctx := context.Background()

	//1. Creamos un pais. Tener un pas es requisito para poder luego crear un jugador
	fmt.Println("\nCreando Países de prueba")
	paises := []string{"Argentina", "Francia"}
	for _, p := range paises {
		_, err := queries.CreatePais(ctx, p)
		if err != nil {
			fmt.Printf("Advertencia al crear país '%s' (puede que ya existiera): %v\n", p, err)
		} else {
			fmt.Printf("País '%s' creado o ya existente.\n", p)
		}
	}

	//2. Creamos un nuevo jugador
	fmt.Println("\nCreando un jugador de prueba")
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

	//3. Obtenemos el jugador recién creado
	fmt.Println("\nObteniendo el jugador por ID")
	fetchedJugador, err := queries.GetJugador(ctx, createdJugador.IDJugador)
	if err != nil {
		log.Fatalf("Error al obtener jugador: %v", err)
	}
	fmt.Printf("Jugador obtenido: %+v\n", fetchedJugador)

	//4.Actualizamos la posición del jugador
	fmt.Println("\nActualizando el jugador")
	err = queries.UpdateJugador(ctx, sqlc.UpdateJugadorParams{
		Posicion:        "Mediapunta", // Nueva posición
		FechaNacimiento: createdJugador.FechaNacimiento,
		Altura:          createdJugador.Altura,
		PaisNombre:      createdJugador.PaisNombre,
		IDJugador:       createdJugador.IDJugador, //PK
		Nombre:          createdJugador.Nombre,    //PK
	})
	if err != nil {
		log.Fatalf("Error al actualizar jugador: %v", err)
	}
	fmt.Println("Jugador actualizado.")

	//Verificamos el cambio obteniendo el jugador de nuevo
	updatedJugador, err := queries.GetJugador(ctx, createdJugador.IDJugador)
	if err != nil {
		log.Fatalf("Error al obtener jugador actualizado: %v", err)
	}
	fmt.Printf("Datos después de actualizar: Posición -> %s\n", updatedJugador.Posicion)

	//5. Listamos todos los jugadores ---
	fmt.Println("\nListando todos los jugadores")
	jugadores, err := queries.ListJugadores(ctx)
	if err != nil {
		log.Fatalf("Error al listar jugadores: %v", err)
	}
	fmt.Printf("Jugadores encontrados (%d):\n", len(jugadores))
	for _, jugador := range jugadores {
		fmt.Printf("- ID: %d, Nombre: %s, Posición: %s\n",
			jugador.IDJugador, jugador.Nombre, jugador.Posicion)
	}

	//6. Probamos crear un club
	fmt.Println("\nProbando crear un Club")
	club, err := queries.CreateClub(ctx, sqlc.CreateClubParams{Nombre: "River Plate", Ciudad: "Buenos Aires"})
	if err != nil {
		fmt.Printf("Advertencia al crear club (puede que ya existiera): %v\n", err)
	} else {
		fmt.Printf("Club creado: %+v\n", club)
	}

	//Listamos clubes
	clubes, err := queries.ListClubs(ctx)
	if err != nil {
		log.Fatalf("Error al listar clubes: %v", err)
	}
	fmt.Printf("Clubes encontrados (%d):\n", len(clubes))
	for _, c := range clubes {
		fmt.Printf("- Nombre: %s, Ciudad: %s\n", c.Nombre, c.Ciudad)
	}

	//7. Probamos generar una relacion Jugador-Club

	CreateJugoParams := sqlc.CreateJugoParams{
		JugadorIDJugador: createdJugador.IDJugador,
		JugadorNombre:    createdJugador.Nombre,
		ClubNombre:       club.Nombre,
		ClubCiudad:       club.Ciudad,
		FechaInicio:      time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
		FechaFin:         time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
	}

	jugo, err := queries.CreateJugo(ctx, CreateJugoParams)
	if err != nil {
		log.Fatalf("Error al crear relacion Jugador-Club: %v", err)
	}
	fmt.Printf("Relacion Jugador-Club creada: %+v\n", jugo)

	//8. Probamos borrar la relacion Jugador-Club
	err = queries.DeleteJugo(ctx, sqlc.DeleteJugoParams{
		JugadorIDJugador: createdJugador.IDJugador,
		JugadorNombre:    createdJugador.Nombre,
		ClubNombre:       club.Nombre,
		ClubCiudad:       club.Ciudad,
	})

	if err != nil {
		log.Fatalf("Error al borrar relacion Jugador-Club: %v", err)
	}

	fmt.Println("Relacion Jugador-Club (jugo) borrada.")

	//9. Borrar club
	err = queries.DeleteClub(ctx, sqlc.DeleteClubParams{
		Nombre: club.Nombre,
		Ciudad: club.Ciudad,
	})
	if err != nil {
		fmt.Printf("Advertencia al borrar club '%s' (puede que no exista o tenga jugadores asociados): %v\n", club.Nombre, err)
	}
	fmt.Println("Club 'River Plate' borrado.")

	//10. BORRAR el jugador ---
	fmt.Println("\nBorrando el jugador")
	err = queries.DeleteJugador(ctx, sqlc.DeleteJugadorParams{
		IDJugador: createdJugador.IDJugador,
		Nombre:    createdJugador.Nombre,
	})
	if err != nil {
		fmt.Printf("Advertencia al borrar jugador '%s' (puede que no exista o tenga relaciones activas): %v\n", createdJugador.Nombre, err)
	}
	fmt.Println("Jugador borrado exitosamente.")

	//11. Borramos pais
	fmt.Println("\nBorrando los países de prueba")
	for _, p := range paises {
		err := queries.DeletePais(ctx, p)
		if err != nil {
			fmt.Printf("Advertencia al borrar país '%s' (puede que no exista o tenga jugadores asociados): %v\n", p, err)
		} else {
			fmt.Printf("País '%s' borrado o no existía.\n", p)
		}
	}

	//12. Creamos una lesión
	fmt.Println("\nCreando una lesion de prueba")
	createLesionParams := sqlc.CreateLesionParams{
		TipoLesion:  "Esguince",
		Descripcion: "Esguince de tobillo leve",
	}

	createdLesion, err := queries.CreateLesion(ctx, createLesionParams)
	if err != nil {
		log.Fatalf("Error al crear lesion: %v", err)
	}
	fmt.Printf("Lesion creada: %+v\n", createdLesion)

	//13. Obtenemos la lesion recién creada
	fmt.Println("\nObteniendo la lesion por su tipo")
	LesionElegida, err := queries.GetLesion(ctx, createdLesion.TipoLesion)
	if err != nil {
		log.Fatalf("Error al obtener lesion: %v", err)
	}
	fmt.Printf("Lesion obtenida: %+v\n", LesionElegida)

	//14. Listamos todas las lesiones
	fmt.Println("\nListando todas las lesiones")
	lesiones, err := queries.ListLesiones(ctx)
	if err != nil {
		log.Fatalf("Error al listar lesiones: %v", err)
	}
	fmt.Printf("Lesiones encontradas (%d):\n", len(lesiones))
	for _, lesion := range lesiones {
		fmt.Printf("-Tipo: %s, Descripción: %s\n",
			lesion.TipoLesion, lesion.Descripcion)
	}

	/*//15. Probamos generar una relacion Jugador-Lesion
	CreateTieneParams := sqlc.CreateTieneParams{
		JugadorIDJugador: createdJugador.IDJugador,
		JugadorNombre:    createdJugador.Nombre,
		LesionTipoLesion: createdLesion.TipoLesion,
		FechaInicio:      time.Date(2023, time.January, 15, 0, 0, 0, 0, time.UTC),
		FechaFin:         time.Date(2023, time.February, 15, 0, 0, 0, 0, time.UTC),
	}

	tuvo, err := queries.CreateTiene(ctx, CreateTieneParams)
	if err != nil {
		log.Fatalf("Error al crear relacion JUgador-Lesion: %v", err)
	}
	fmt.Printf("Relacion Jugador-Lesion creada: %+v\n", tuvo)*/

	//16. BORRAR la lesion ---
	fmt.Println("\nBorrando la lesion")
	for _, l := range lesiones {
		err := queries.DeleteLesion(ctx, l.TipoLesion)
		if err != nil {
			fmt.Printf("Advertencia al borrar lesion '%s' (puede que no exista o tenga jugadores asociados): %v\n", l.TipoLesion, err)
		} else {
			fmt.Printf("Lesion '%s' borrado o no existía.\n", l.TipoLesion)
		}
	}
}
