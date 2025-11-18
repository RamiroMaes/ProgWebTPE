-- name: CreateJugador :one
INSERT INTO Jugador (Nombre, iD_Jugador, Posicion,Fecha_Nacimiento, Altura, Pais_Nombre)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING Nombre, iD_Jugador, Posicion,Fecha_Nacimiento, Altura, Pais_Nombre;
-- name: CreatePais :one
INSERT INTO Pais(Nombre)
VALUES ($1)
RETURNING Nombre;

-- name: GetJugador :one
SELECT * FROM Jugador
WHERE iD_Jugador = $1;

-- name: GetPais :one
SELECT * FROM Pais
WHERE Nombre = $1;


-- name: ListJugadores :many
SELECT iD_Jugador, Nombre FROM Jugador;
-- name: ListPlantel :many
SELECT * FROM Jugador;
-- name: ListPaises :many
SELECT * FROM Pais;


-- name: UpdateJugador :exec
UPDATE Jugador
SET Nombre = $1, Posicion = $2, Fecha_Nacimiento = $3, Altura = $4, Pais_Nombre = $5
WHERE iD_Jugador = $6;
-- name: UpdatePais :exec
UPDATE Pais
SET Nombre = $1
WHERE Nombre = $2;


-- name: DeleteJugador :exec
DELETE FROM Jugador
WHERE iD_Jugador = $1;
-- name: DeletePais :exec
DELETE FROM Pais
WHERE Nombre = $1;