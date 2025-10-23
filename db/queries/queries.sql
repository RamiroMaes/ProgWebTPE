--Una consulta para crear un nuevo registro (Create...) usando el csv "plantelRiverPlate.csv"

-- name: CreateJugador :one
INSERT INTO Jugador (Nombre, iD_Jugador, Posicion,Fecha_Nacimiento, Altura, Pais_Nombre)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING Nombre, iD_Jugador, Posicion,Fecha_Nacimiento, Altura, Pais_Nombre;
-- name: CreatePais :one
INSERT INTO Pais(Nombre)
VALUES ($1)
RETURNING Nombre;
-- name: CreateClub :one
INSERT INTO Club (Nombre, Ciudad)
VALUES ($1, $2)
RETURNING Nombre, Ciudad;
-- name: CreateLesion :one
INSERT INTO Lesion (Tipo_Lesion, Descripcion)
VALUES ($1, $2)
RETURNING Tipo_Lesion, Descripcion;
-- name: CreateJugo :one
INSERT INTO Jugo (fecha_inicio, fecha_fin, Jugador_Nombre, Jugador_iD_Jugador, Club_Nombre, Club_Ciudad)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: CreateTiene :one
INSERT INTO Tiene (fecha_inicio, fecha_fin, Jugador_Nombre, Jugador_iD_Jugador, Lesion_Tipo_Lesion)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

--Una consulta para obtener un registro por su ID (Get...).

-- name: GetJugador :one
SELECT * FROM Jugador
WHERE iD_Jugador = $1;

-- name: GetClub :one
SELECT * FROM Club
WHERE Nombre = $1 AND Ciudad = $2;

-- name: GetLesion :one
SELECT * FROM Lesion
WHERE Tipo_Lesion = $1;

-- name: GetPais :one
SELECT * FROM Pais
WHERE Nombre = $1;

-- name: GetJugo :one
SELECT * FROM Jugo
WHERE Jugador_iD_Jugador = $1 AND Jugador_Nombre = $2 AND Club_Nombre = $3 AND Club_Ciudad = $4;

-- name: GetTiene :one
SELECT * FROM Tiene
WHERE Jugador_iD_Jugador = $1 AND Jugador_Nombre = $2 AND Lesion_Tipo_Lesion = $3;

--Una consulta para listar todos los registros (List...).

-- name: ListJugadores :many
SELECT iD_Jugador, Nombre FROM Jugador;
-- name: ListJugadoresCompleto :many
SELECT * FROM Jugador;
-- name: ListClubs :many
SELECT * FROM Club;
-- name: ListLesiones :many
SELECT * FROM Lesion;
-- name: ListPaises :many
SELECT * FROM Pais;
-- name: ListJugos :many
SELECT * FROM Jugo;
-- name: ListTienes :many
SELECT * FROM Tiene;

--Una consulta para actualizar un registro (Update...).

-- name: UpdateJugador :exec
UPDATE Jugador
SET Posicion = $1, Fecha_Nacimiento = $2, Altura = $3, Pais_Nombre = $4
WHERE iD_Jugador = $5 AND Nombre = $6;
-- name: UpdateClub :exec
UPDATE Club
SET Nombre = $1, Ciudad = $2
WHERE Nombre = $3 AND Ciudad = $4;
-- name: UpdateLesion :exec
UPDATE Lesion
SET Descripcion = $1
WHERE Tipo_Lesion = $2;
-- name: UpdatePais :exec
UPDATE Pais
SET Nombre = $1
WHERE Nombre = $2;
-- name: UpdateJugo :exec
UPDATE Jugo
SET fecha_inicio = $1, fecha_fin = $2
WHERE Jugador_iD_Jugador = $3 AND Jugador_Nombre = $4 AND Club_Nombre = $5 AND Club_Ciudad = $6;
-- name: UpdateTiene :exec
UPDATE Tiene
SET fecha_inicio = $1, fecha_fin = $2
WHERE Jugador_iD_Jugador = $3 AND Jugador_Nombre = $4 AND Lesion_Tipo_Lesion = $5;

--Una consulta para borrar un registro (Delete...).
-- name: DeleteJugador :exec
DELETE FROM Jugador
WHERE iD_Jugador = $1 AND Nombre = $2;
-- name: DeleteClub :exec
DELETE FROM Club
WHERE Nombre = $1 AND Ciudad = $2;
-- name: DeleteLesion :exec
DELETE FROM Lesion
WHERE Tipo_Lesion = $1;
-- name: DeletePais :exec
DELETE FROM Pais
WHERE Nombre = $1;
-- name: DeleteJugo :exec
DELETE FROM Jugo
WHERE Jugador_iD_Jugador = $1 AND Jugador_Nombre = $2 AND Club_Nombre = $3 AND Club_Ciudad = $4;
-- name: DeleteTiene :exec
DELETE FROM Tiene
WHERE Jugador_iD_Jugador = $1 AND Jugador_Nombre = $2 AND Lesion_Tipo_Lesion = $3;