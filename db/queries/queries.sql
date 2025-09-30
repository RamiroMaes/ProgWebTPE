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

--Una consulta para listar todos los registros (List...).

-- name: ListJugadores :many
SELECT * FROM Jugador;
-- name: ListClubs :many
SELECT * FROM Club;
-- name: ListLesiones :many
SELECT * FROM Lesion;

--Una consulta para actualizar un registro (Update...).

-- name: UpdateJugador :exec
UPDATE Jugador
SET Posicion = $1, Fecha_Nacimiento = $2, Altura = $3, Pais_Nombre = $4
WHERE iD_Jugador = $5 AND Nombre = $6;

--Preguntar
-- name: UpdateClub :exec
UPDATE Club
SET Nombre = $1, Ciudad = $2
WHERE Nombre = $3 AND Ciudad = $4;
-- name: UpdateLesion :exec
UPDATE Lesion
SET Descripcion = $1
WHERE Tipo_Lesion = $2;


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
