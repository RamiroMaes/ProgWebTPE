#!/usr/bin/env bash

BASE_URL="${BASE_URL:-http://localhost:8080}"
JSON="Content-Type: application/json"

printf "\n sentencias de paises \n"
curl -sS -H "$JSON" --data '{"nombre":"Argentina"}' "$BASE_URL/paises"
echo
curl -sS -H "$JSON" --data '{"nombre":"Francia"}' "$BASE_URL/paises"
echo
curl -sS "$BASE_URL/paises"
echo

printf "\n sentencias de lesiones \n"
curl -sS -H "$JSON" --data '{"tipo_lesion":"Esguince","descripcion":"Tobillo derecho"}' "$BASE_URL/lesiones"
echo
curl -sS "$BASE_URL/lesiones"
echo

printf "\n sentencias de club \n"
curl -sS -H "$JSON" --data '{"nombre":"PSG","ciudad":"Paris"}' "$BASE_URL/clubs"
echo
curl -sS "$BASE_URL/clubs"
echo
curl -sS "$BASE_URL/clubs/PSG/Paris"
echo

printf "\n actualizamos club psg\n"
curl -sS -X PUT -H "$JSON" --data '{"nombre":"Paris Saint-Germain","ciudad":"Paris","nombre_2":"PSG","ciudad_2":"Paris"}' \
  "$BASE_URL/clubs/PSG/Paris"
echo

printf "\n sentencias de jugadores\n"
curl -sS -H "$JSON" --data '{"nombre":"Juan Perez","id_jugador":1,"posicion":"Delantero","fecha_nacimiento":"1995-01-01T00:00:00Z","altura":180,"pais_nombre":"Argentina"}' \
  "$BASE_URL/jugadores"
echo
curl -sS -H "$JSON" --data '{"nombre":"Pity Martinez","id_jugador":2,"posicion":"Delantero","fecha_nacimiento":"1994-02-03T00:00:00Z","altura":170,"pais_nombre":"Argentina"}' \
  "$BASE_URL/jugadores"
echo
curl -sS "$BASE_URL/jugadores"
echo
curl -sS "$BASE_URL/jugadores/1"
echo
curl -sS -X PUT -H "$JSON" --data '{"nombre":"Juan Perez","id_jugador":1,"posicion":"Mediocampista","fecha_nacimiento":"1995-01-01T00:00:00Z","altura":180,"pais_nombre":"Argentina"}' \
  "$BASE_URL/jugadores/1"
echo
curl -sS "$BASE_URL/plantel"
echo

#printf "\n borramos ciertos atributos de las tablas\n"
#curl -sS -X DELETE "$BASE_URL/jugadores/1"
#echo
#curl -sS -X DELETE "$BASE_URL/clubs/PSG/Paris"
#echo
#curl -sS -X DELETE "$BASE_URL/lesiones/Esguince"
#echo
#curl -sS -X DELETE "$BASE_URL/paises/Francia"
#echo

printf "\n vemos como quedan finalmente las tablas\n"
curl -sS "$BASE_URL/paises"; echo
curl -sS "$BASE_URL/lesiones"; echo
curl -sS "$BASE_URL/clubs"; echo
curl -sS "$BASE_URL/jugadores"; echo