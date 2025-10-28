#!/usr/bin/env bash

BASE_URL="http://localhost:8080"

echo "paises"
curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre": "Argentina"}' \
  "$BASE_URL/paises"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre": "Francia"}' \
  "$BASE_URL/paises"

curl -X GET "$BASE_URL/paises"


echo "\n lesiones"
curl -X POST -H "Content-Type: application/json" \
  --data '{"tipo_lesion": "Esguince", "descripcion": "Tobillo derecho"}' \
  "$BASE_URL/lesiones"

curl -X GET "$BASE_URL/lesiones"


echo "\n clubs"
curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre": "PSG", "ciudad": "Paris"}' \
  "$BASE_URL/clubs"

curl -X GET "$BASE_URL/clubs"

curl -X GET "$BASE_URL/clubs/PSG/Paris"

echo "\n actualizando club"
curl -X PUT -H "Content-Type: application/json" \
  --data '{"nombre": "Paris Saint-Germain", "ciudad": "Paris", "nombre_2": "PSG", "ciudad_2": "Paris"}' \
  "$BASE_URL/clubs/PSG/Paris"

echo "\n jugadores"
curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre": "Juan Perez", "id_jugador": 1, "posicion": "Delantero", "fecha_nacimiento": "1995-01-01T00:00:00Z", "altura": 180, "pais_nombre": "Argentina"}' \
  "$BASE_URL/jugadores"

curl -X GET "$BASE_URL/jugadores"

curl -X GET "$BASE_URL/jugadores/1"

curl -X PUT -H "Content-Type: application/json" \
  --data '{"nombre": "Juan Perez", "id_jugador": 1, "posicion": "Mediocampista", "fecha_nacimiento": "1995-01-01T00:00:00Z", "altura": 180, "pais_nombre": "Argentina"}' \
  "$BASE_URL/jugadores/1"

curl -X GET "$BASE_URL/plantel"

#echo "\n borramos todo"
#curl -X DELETE "$BASE_URL/jugadores/1/Juan%20Perez"
#curl -X DELETE "$BASE_URL/clubs/PSG/Paris"
#curl -X DELETE "$BASE_URL/lesiones/Esguince"
#curl -X DELETE "$BASE_URL/paises/Francia"

echo "\n estado final"
