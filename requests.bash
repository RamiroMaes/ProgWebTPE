#!/usr/bin/env bash

BASE_URL="http://localhost:8080"

echo "\n Paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Brasil"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Uruguay"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"España"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Portugal"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Alemania"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Italia"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Inglaterra"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Países Bajos"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Colombia"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Bélgica"}' "$BASE_URL/paises"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Polonia"}' "$BASE_URL/paises"
curl -X GET "$BASE_URL/paises"

echo "\n Lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Contractura","descripcion":"Isquiotibial"}' "$BASE_URL/lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Desgarro","descripcion":"Aductor izquierdo"}' "$BASE_URL/lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Rotura de ligamentos","descripcion":"LCA rodilla derecha"}' "$BASE_URL/lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Tendinitis","descripcion":"Tendón de Aquiles"}' "$BASE_URL/lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Fractura","descripcion":"Metatarsiano"}' "$BASE_URL/lesiones"
curl -X POST -H "Content-Type: application/json" --data '{"tipo_lesion":"Pubalgia","descripcion":"Dolor inguinal"}' "$BASE_URL/lesiones"
curl -X GET "$BASE_URL/lesiones"

echo "\n Clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Real Madrid","ciudad":"Madrid"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Barcelona","ciudad":"Barcelona"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Bayern Munich","ciudad":"Munich"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Juventus","ciudad":"Turin"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Manchester City","ciudad":"Manchester"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Manchester United","ciudad":"Manchester"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Liverpool","ciudad":"Liverpool"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"River Plate","ciudad":"Buenos Aires"}' "$BASE_URL/clubs"
curl -X POST -H "Content-Type: application/json" --data '{"nombre":"Boca Juniors","ciudad":"Buenos Aires"}' "$BASE_URL/clubs"
curl -X GET "$BASE_URL/clubs"
curl -X GET "$BASE_URL/clubs/Real%20Madrid/Madrid"
curl -X GET "$BASE_URL/clubs/Manchester%20City/Manchester"
curl -X GET "$BASE_URL/clubs/River%20Plate/Buenos%20Aires"

echo "\n Jugadores"
curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Lionel Messi","id_jugador":2,"posicion":"Delantero","fecha_nacimiento":"1987-06-24T00:00:00Z","altura":170,"pais_nombre":"Argentina"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Kylian Mbappé","id_jugador":3,"posicion":"Delantero","fecha_nacimiento":"1998-12-20T00:00:00Z","altura":178,"pais_nombre":"Francia"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Neymar Jr","id_jugador":4,"posicion":"Delantero","fecha_nacimiento":"1992-02-05T00:00:00Z","altura":175,"pais_nombre":"Brasil"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Cristiano Ronaldo","id_jugador":5,"posicion":"Delantero","fecha_nacimiento":"1985-02-05T00:00:00Z","altura":187,"pais_nombre":"Portugal"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Kevin De Bruyne","id_jugador":6,"posicion":"Mediocampista","fecha_nacimiento":"1991-06-28T00:00:00Z","altura":181,"pais_nombre":"Bélgica"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Robert Lewandowski","id_jugador":7,"posicion":"Delantero","fecha_nacimiento":"1988-08-21T00:00:00Z","altura":185,"pais_nombre":"Polonia"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Virgil van Dijk","id_jugador":8,"posicion":"Defensor","fecha_nacimiento":"1991-07-08T00:00:00Z","altura":193,"pais_nombre":"Países Bajos"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Manuel Neuer","id_jugador":9,"posicion":"Arquero","fecha_nacimiento":"1986-03-27T00:00:00Z","altura":193,"pais_nombre":"Alemania"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Gianluigi Donnarumma","id_jugador":10,"posicion":"Arquero","fecha_nacimiento":"1999-02-25T00:00:00Z","altura":196,"pais_nombre":"Italia"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"Luis Suárez","id_jugador":11,"posicion":"Delantero","fecha_nacimiento":"1987-01-24T00:00:00Z","altura":182,"pais_nombre":"Uruguay"}' \
  "$BASE_URL/jugadores"

curl -X POST -H "Content-Type: application/json" \
  --data '{"nombre":"James Rodríguez","id_jugador":12,"posicion":"Mediocampista","fecha_nacimiento":"1991-07-12T00:00:00Z","altura":181,"pais_nombre":"Colombia"}' \
  "$BASE_URL/jugadores"



curl -X GET "$BASE_URL/jugadores"


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