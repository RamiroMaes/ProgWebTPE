Nombre del contenedor de DOcker:
    some-postgres
Usuario de la base de datos:
    postgres
Nombre de la base de datos:
    tpespecial
Contraseña:
    XYZ

Pasos:
    1-Primeramente se debe contar con el lenguaje 'Go', la plataforma 'Docker' y 'sqlc' instalados en el sistema.
    3- Tener iniciado Docker en el sistema.
    3-correr por la consola del directorio del proyecto:
        a- docker run --name some-postgres -e POSTGRES_PASSWORD=XYZ -p 5432:5432 -d postgres:16
            -esto crea un contenedor, si este no existe previamente. Tener en cuenta que el numero de puerto debe estar libre.

        b- docker exec -it some-postgres psql -U postgres
            -Se conecta al servidor de Postgres.

        c- CREATE DATABASE tpespecial;
            -Se crea la base de datos.

        d-\q
            -Dentro del servidor de Postgres, para volver a la terminal del contenedor
        e- cat ./db/schema/schema.sql | docker exec -i -e PGPASSWORD=XYZ some-postgres psql -U postgres -d tpespecial
            -creamos las tablas que se encuentran en schema.sql
            
        f- go run .
            -Asi, se corre el programa, esperando como respuesta:
            
                Conexión exitosa a la base de datos

                Creando Países de prueba
                País 'Argentina' creado o ya existente.
                País 'Francia' creado o ya existente.

                Creando un jugador de prueba
                Jugador creado: {Nombre:Lionel Messi IDJugador:10 Posicion:Delantero FechaNacimiento:1987-06-24 00:00:00 +0000 +0000 Altura:1.70 PaisNombre:Argentina}

                Obteniendo el jugador por ID
                Jugador obtenido: {Nombre:Lionel Messi IDJugador:10 Posicion:Delantero FechaNacimiento:1987-06-24 00:00:00 +0000 +0000 Altura:1.70 PaisNombre:Argentina}

                Actualizando el jugador
                Jugador actualizado.
                Datos después de actualizar: Posición -> Mediapunta

                Listando todos los jugadores
                Jugadores encontrados (1):
                - ID: 10, Nombre: Lionel Messi, Posición: Mediapunta

                Probando crear un Club
                Club creado: {Nombre:River Plate Ciudad:Buenos Aires}
                Clubes encontrados (1):
                - Nombre: River Plate, Ciudad: Buenos Aires
                Relacion Jugador-Club creada: {FechaInicio:2020-01-01 00:00:00 +0000 +0000 FechaFin:2023-12-31 00:00:00 +0000 +0000 JugadorNombre:Lionel Messi JugadorIDJugador:10 ClubNombre:River Plate ClubCiudad:Buenos Aires}
                Relacion Jugador-Club (jugo) borrada.
                Club 'River Plate' borrado.

                Borrando el jugador
                Jugador borrado exitosamente.

                Borrando los países de prueba
                País 'Argentina' borrado o no existía.
                País 'Francia' borrado o no existía.

                Creando una lesion de prueba
                Lesion creada: {TipoLesion:Esguince Descripcion:Esguince de tobillo leve}

                Obteniendo la lesion por su tipo
                Lesion obtenida: {TipoLesion:Esguince Descripcion:Esguince de tobillo leve}

                Listando todas las lesiones
                Lesiones encontradas (1):
                -Tipo: Esguince, Descripción: Esguince de tobillo leve

                Borrando la lesion
                Lesion 'Esguince' borrado o no existía.



    -docker exec -it -e PGPASSWORD=XYZ some-postgres psql -U postgres -d tpespecial
        -Si se quiere reconectar a la base