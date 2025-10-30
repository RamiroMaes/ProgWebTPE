#!/usr/bin/env bash

BASE_URL="${BASE_URL:-http://localhost:8080}"
JSON="Content-Type: application/json"

printf "\n sentencias de paises \n"
curl -sS -H "$JSON" --data '{"nombre":"Argentina"}' "$BASE_URL/paises"
echo
curl -sS -H "$JSON" --data '{"nombre":"Colombia"}' "$BASE_URL/paises"
echo
curl -sS -H "$JSON" --data '{"nombre":"Paraguay"}' "$BASE_URL/paises"
echo
curl -sS -H "$JSON" --data '{"nombre":"Uruguay"}' "$BASE_URL/paises"
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
curl -sS -H "$JSON" --data '{"nombre":"Franco Armani","id_jugador":1,"posicion":"Arquero","fecha_nacimiento":"1986-10-16T00:00:00Z","altura":189,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 1 (ID: 1) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Jeremias Ledesma","id_jugador":25,"posicion":"Arquero","fecha_nacimiento":"1993-02-13T00:00:00Z","altura":186,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 2 (ID: 25) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Santiago Beltrán","id_jugador":41,"posicion":"Arquero","fecha_nacimiento":"2004-10-04T00:00:00Z","altura":189,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 3 (ID: 41) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Federico Gattoni","id_jugador":2,"posicion":"Defensor","fecha_nacimiento":"1999-02-16T00:00:00Z","altura":183,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 4 (ID: 2) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Gonzalo Montiel","id_jugador":4,"posicion":"Defensor","fecha_nacimiento":"1997-01-01T00:00:00Z","altura":175,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 5 (ID: 4) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Juan Portillo","id_jugador":5,"posicion":"Defensor","fecha_nacimiento":"2000-05-18T00:00:00Z","altura":166,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 6 (ID: 5) cargado"

curl -sS -H "$JSON" --data '{"nombre":"German Pezzella","id_jugador":6,"posicion":"Defensor","fecha_nacimiento":"1991-06-27T00:00:00Z","altura":187,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 7 (ID: 6) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Lautaro Rivero","id_jugador":13,"posicion":"Defensor","fecha_nacimiento":"2003-11-01T00:00:00Z","altura":185,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 8 (ID: 13) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Sebastián Boselli","id_jugador":14,"posicion":"Defensor","fecha_nacimiento":"2003-12-04T00:00:00Z","altura":183,"pais_nombre":"Uruguay"}' "$BASE_URL/jugadores"
echo -e "\nJugador 9 (ID: 14) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Fabricio Bustos","id_jugador":16,"posicion":"Defensor","fecha_nacimiento":"1996-04-28T00:00:00Z","altura":167,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 10 (ID: 16) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Paulo Diaz","id_jugador":17,"posicion":"Defensor","fecha_nacimiento":"1994-08-25T00:00:00Z","altura":180,"pais_nombre":"Chile"}' "$BASE_URL/jugadores"
echo -e "\nJugador 11 (ID: 17) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Milton Casco","id_jugador":20,"posicion":"Defensor","fecha_nacimiento":"1988-04-11T00:00:00Z","altura":170,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 12 (ID: 20) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Marcos Acuña","id_jugador":21,"posicion":"Defensor","fecha_nacimiento":"1991-10-28T00:00:00Z","altura":172,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 13 (ID: 21) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Lucas Martínez Quarta","id_jugador":28,"posicion":"Defensor","fecha_nacimiento":"1996-05-10T00:00:00Z","altura":183,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 14 (ID: 28) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Maximiliano Meza","id_jugador":8,"posicion":"Mediocampista","fecha_nacimiento":"1992-12-15T00:00:00Z","altura":181,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 15 (ID: 8) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Juan Fernando Quintero","id_jugador":10,"posicion":"Mediocampista","fecha_nacimiento":"1993-01-18T00:00:00Z","altura":168,"pais_nombre":"Colombia"}' "$BASE_URL/jugadores"
echo -e "\nJugador 16 (ID: 10) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Gonzalo Martínez","id_jugador":18,"posicion":"Mediocampista","fecha_nacimiento":"1993-06-13T00:00:00Z","altura":172,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 17 (ID: 18) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Kevin Castaño","id_jugador":22,"posicion":"Mediocampista","fecha_nacimiento":"2000-09-29T00:00:00Z","altura":177,"pais_nombre":"Colombia"}' "$BASE_URL/jugadores"
echo -e "\nJugador 18 (ID: 22) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Matías Galarza Fonda","id_jugador":23,"posicion":"Mediocampista","fecha_nacimiento":"2002-02-11T00:00:00Z","altura":175,"pais_nombre":"Colombia"}' "$BASE_URL/jugadores"
echo -e "\nJugador 19 (ID: 23) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Enzo Pérez","id_jugador":24,"posicion":"Mediocampista","fecha_nacimiento":"1986-02-22T00:00:00Z","altura":178,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 20 (ID: 24) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Nacho Fernández","id_jugador":26,"posicion":"Mediocampista","fecha_nacimiento":"1990-01-12T00:00:00Z","altura":182,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 21 (ID: 26) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Giuliano Galoppo","id_jugador":34,"posicion":"Mediocampista","fecha_nacimiento":"1999-06-18T00:00:00Z","altura":179,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 22 (ID: 34) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Giorgio Costantini","id_jugador":35,"posicion":"Mediocampista","fecha_nacimiento":"2006-04-16T00:00:00Z","altura":185,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 23 (ID: 35) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Santiago Lencina","id_jugador":39,"posicion":"Mediocampista","fecha_nacimiento":"2005-09-04T00:00:00Z","altura":173,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 24 (ID: 39) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Maximiliano Salas","id_jugador":7,"posicion":"Delantero","fecha_nacimiento":"1997-12-01T00:00:00Z","altura":172,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 25 (ID: 7) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Miguel Borja","id_jugador":9,"posicion":"Delantero","fecha_nacimiento":"1993-01-26T00:00:00Z","altura":183,"pais_nombre":"Colombia"}' "$BASE_URL/jugadores"
echo -e "\nJugador 26 (ID: 9) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Facundo Colidio","id_jugador":11,"posicion":"Delantero","fecha_nacimiento":"2000-01-04T00:00:00Z","altura":175,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 27 (ID: 11) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Sebastián Driussi","id_jugador":15,"posicion":"Delantero","fecha_nacimiento":"1996-02-09T00:00:00Z","altura":179,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 28 (ID: 15) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Bautista Dadín","id_jugador":27,"posicion":"Delantero","fecha_nacimiento":"2006-05-20T00:00:00Z","altura":175,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 29 (ID: 27) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Alex Woiski","id_jugador":29,"posicion":"Delantero","fecha_nacimiento":"2006-03-17T00:00:00Z","altura":170,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 30 (ID: 29) cargado"

curl -sS -H "$JSON" --data '{"nombre":"Ian Subiabre","id_jugador":38,"posicion":"Delantero","fecha_nacimiento":"2007-01-01T00:00:00Z","altura":172,"pais_nombre":"Argentina"}' "$BASE_URL/jugadores"
echo -e "\nJugador 31 (ID: 38) cargado"
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