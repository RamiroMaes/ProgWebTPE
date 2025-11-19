#!/usr/bin/env bash

BASE_URL="${BASE_URL:-http://localhost:8080}"

printf "\n sentencias de paises \n"
paises=(
    "Afganistan" "Albania" "Alemania" "Andorra" "Angola" "Antigua y Barbuda" 
    "Arabia Saudita" "Argelia" "Argentina" "Armenia" "Australia" "Austria" 
    "Azerbaiyan" "Bahamas" "Banglades" "Barbados" "Barein" "Belgica" "Belice" 
    "Benin" "Bielorrusia" "Birmania" "Bolivia" "Bosnia y Herzegovina" "Botsuana" 
    "Brasil" "Brunei" "Bulgaria" "Burkina Faso" "Burundi" "Butan" "Cabo Verde" 
    "Camboya" "Camerun" "Canada" "Catar" "Chad" "Chile" "China" "Chipre" 
    "Ciudad del Vaticano" "Colombia" "Comoras" "Corea del Norte" "Corea del Sur" 
    "Costa de Marfil" "Costa Rica" "Croacia" "Cuba" "Dinamarca" "Dominica" 
    "Ecuador" "Egipto" "El Salvador" "Emiratos Arabes Unidos" "Eritrea" "Eslovaquia" 
    "Eslovenia" "Espana" "Estados Unidos" "Estonia" "Etiopia" "Filipinas" 
    "Finlandia" "Fiyi" "Francia" "Gabon" "Gambia" "Georgia" "Ghana" "Granada" 
    "Grecia" "Guatemala" "Guinea" "Guinea-Bisau" "Guinea Ecuatorial" "Guyana" 
    "Haiti" "Honduras" "Hungria" "India" "Indonesia" "Irak" "Iran" "Irlanda" 
    "Islandia" "Islas Marshall" "Islas Salomon" "Israel" "Italia" "Jamaica" 
    "Japon" "Jordania" "Kazajistan" "Kenia" "Kirguistan" "Kiribati" "Kuwait" 
    "Laos" "Lesoto" "Letonia" "Libano" "Liberia" "Libia" "Liechtenstein" 
    "Lituania" "Luxemburgo" "Macedonia del Norte" "Madagascar" "Malasia" "Malaui" 
    "Maldivas" "Mali" "Malta" "Marruecos" "Mauricio" "Mauritania" "Mexico" 
    "Micronesia" "Moldavia" "Monaco" "Mongolia" "Montenegro" "Mozambique" 
    "Namibia" "Nauru" "Nepal" "Nicaragua" "Niger" "Nigeria" "Noruega" 
    "Nueva Zelanda" "Oman" "Paises Bajos" "Pakistan" "Palaos" "Panama" 
    "Papua Nueva Guinea" "Paraguay" "Peru" "Polonia" "Portugal" "Reino Unido" 
    "Republica Centroafricana" "Republica Checa" "Republica del Congo" 
    "Republica Democratica del Congo" "Republica Dominicana" "Ruanda" "Rumania" 
    "Rusia" "Samoa" "San Cristobal y Nieves" "San Marino" 
    "San Vicente y las Granadinas" "Santa Lucia" "Santo Tome y Principe" "Senegal" 
    "Serbia" "Seychelles" "Sierra Leona" "Singapur" "Siria" "Somalia" "Sri Lanka" 
    "Suazilandia" "Sudafrica" "Sudan" "Sudan del Sur" "Suecia" "Suiza" "Surinam" 
    "Tailandia" "Tanzania" "Tayikistan" "Timor Oriental" "Togo" "Tonga" 
    "Trinidad y Tobago" "Tunez" "Turkmenistan" "Turquia" "Tuvalu" "Ucrania" 
    "Uganda" "Uruguay" "Uzbekistan" "Vanuatu" "Venezuela" "Vietnam" "Yemen" 
    "Yibuti" "Zambia" "Zimbabue"
)

# Ejecutamos el bucle para hacer un curl por cada país
for pais in "${paises[@]}"; do
  echo "Cargando: $pais"
  curl -sS -X POST -d "nombre=$pais" "$BASE_URL/paises"
  echo ""
done
curl -sS "$BASE_URL/paises"
echo


printf "\n sentencias de jugadores\n"
curl -X POST -d "agregarNombre=Franco Armani&agregarPosicion=Arquero&agregarPais=Argentina&agregarNumero=1&agregarAltura=189&agregarFechaNacimiento=1986-10-16" "$BASE_URL/jugadores"
echo -e "\nJugador 1 (ID: 1) cargado"

curl -X POST -d "agregarNombre=Jeremias Ledesma&agregarPosicion=Arquero&agregarPais=Argentina&agregarNumero=25&agregarAltura=186&agregarFechaNacimiento=1993-02-13" "$BASE_URL/jugadores"
echo -e "\nJugador 2 (ID: 25) cargado"

curl -X POST -d "agregarNombre=Santiago Beltrán&agregarPosicion=Arquero&agregarPais=Argentina&agregarNumero=41&agregarAltura=189&agregarFechaNacimiento=2004-10-04" "$BASE_URL/jugadores"
echo -e "\nJugador 3 (ID: 41) cargado"

curl -X POST -d "agregarNombre=Federico Gattoni&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=2&agregarAltura=183&agregarFechaNacimiento=1999-02-16" "$BASE_URL/jugadores"
echo -e "\nJugador 4 (ID: 2) cargado"

curl -X POST -d "agregarNombre=Gonzalo Montiel&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=4&agregarAltura=175&agregarFechaNacimiento=1997-01-01" "$BASE_URL/jugadores"
echo -e "\nJugador 5 (ID: 4) cargado"

curl -X POST -d "agregarNombre=Juan Portillo&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=5&agregarAltura=166&agregarFechaNacimiento=2000-05-18" "$BASE_URL/jugadores"
echo -e "\nJugador 6 (ID: 5) cargado"

curl -X POST -d "agregarNombre=German Pezzella&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=6&agregarAltura=187&agregarFechaNacimiento=1991-06-27" "$BASE_URL/jugadores"
echo -e "\nJugador 7 (ID: 6) cargado"

curl -X POST -d "agregarNombre=Lautaro Rivero&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=13&agregarAltura=185&agregarFechaNacimiento=2003-11-01" "$BASE_URL/jugadores"
echo -e "\nJugador 8 (ID: 13) cargado"

curl -X POST -d "agregarNombre=Sebastián Boselli&agregarPosicion=Defensor&agregarPais=Uruguay&agregarNumero=14&agregarAltura=183&agregarFechaNacimiento=2003-12-04" "$BASE_URL/jugadores"
echo -e "\nJugador 9 (ID: 14) cargado"

curl -X POST -d "agregarNombre=Fabricio Bustos&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=16&agregarAltura=167&agregarFechaNacimiento=1996-04-28" "$BASE_URL/jugadores"
echo -e "\nJugador 10 (ID: 16) cargado"

curl -X POST -d "agregarNombre=Paulo Diaz&agregarPosicion=Defensor&agregarPais=Chile&agregarNumero=17&agregarAltura=180&agregarFechaNacimiento=1994-08-25" "$BASE_URL/jugadores"
echo -e "\nJugador 11 (ID: 17) cargado"

curl -X POST -d "agregarNombre=Milton Casco&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=20&agregarAltura=170&agregarFechaNacimiento=1988-04-11" "$BASE_URL/jugadores"
echo -e "\nJugador 12 (ID: 20) cargado"

curl -X POST -d "agregarNombre=Marcos Acuña&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=21&agregarAltura=172&agregarFechaNacimiento=1991-10-28" "$BASE_URL/jugadores"
echo -e "\nJugador 13 (ID: 21) cargado"

curl -X POST -d "agregarNombre=Lucas Martínez Quarta&agregarPosicion=Defensor&agregarPais=Argentina&agregarNumero=28&agregarAltura=183&agregarFechaNacimiento=1996-05-10" "$BASE_URL/jugadores"
echo -e "\nJugador 14 (ID: 28) cargado"

curl -X POST -d "agregarNombre=Maximiliano Meza&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=8&agregarAltura=181&agregarFechaNacimiento=1992-12-15" "$BASE_URL/jugadores"
echo -e "\nJugador 15 (ID: 8) cargado"

curl -X POST -d "agregarNombre=Juan Fernando Quintero&agregarPosicion=Mediocampista&agregarPais=Colombia&agregarNumero=10&agregarAltura=168&agregarFechaNacimiento=1993-01-18" "$BASE_URL/jugadores"
echo -e "\nJugador 16 (ID: 10) cargado"

curl -X POST -d "agregarNombre=Gonzalo Martínez&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=18&agregarAltura=172&agregarFechaNacimiento=1993-06-13" "$BASE_URL/jugadores"
echo -e "\nJugador 17 (ID: 18) cargado"

curl -X POST -d "agregarNombre=Kevin Castaño&agregarPosicion=Mediocampista&agregarPais=Colombia&agregarNumero=22&agregarAltura=177&agregarFechaNacimiento=2000-09-29" "$BASE_URL/jugadores"
echo -e "\nJugador 18 (ID: 22) cargado"

curl -X POST -d "agregarNombre=Matías Galarza Fonda&agregarPosicion=Mediocampista&agregarPais=Colombia&agregarNumero=23&agregarAltura=175&agregarFechaNacimiento=2002-02-11" "$BASE_URL/jugadores"
echo -e "\nJugador 19 (ID: 23) cargado"

curl -X POST -d "agregarNombre=Enzo Pérez&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=24&agregarAltura=178&agregarFechaNacimiento=1986-02-22" "$BASE_URL/jugadores"
echo -e "\nJugador 20 (ID: 24) cargado"

curl -X POST -d "agregarNombre=Nacho Fernández&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=26&agregarAltura=182&agregarFechaNacimiento=1990-01-12" "$BASE_URL/jugadores"
echo -e "\nJugador 21 (ID: 26) cargado"

curl -X POST -d "agregarNombre=Giuliano Galoppo&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=34&agregarAltura=179&agregarFechaNacimiento=1999-06-18" "$BASE_URL/jugadores"
echo -e "\nJugador 22 (ID: 34) cargado"

curl -X POST -d "agregarNombre=Giorgio Costantini&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=35&agregarAltura=185&agregarFechaNacimiento=2006-04-16" "$BASE_URL/jugadores"
echo -e "\nJugador 23 (ID: 35) cargado"

curl -X POST -d "agregarNombre=Santiago Lencina&agregarPosicion=Mediocampista&agregarPais=Argentina&agregarNumero=39&agregarAltura=173&agregarFechaNacimiento=2005-09-04" "$BASE_URL/jugadores"
echo -e "\nJugador 24 (ID: 39) cargado"

curl -X POST -d "agregarNombre=Maximiliano Salas&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=7&agregarAltura=172&agregarFechaNacimiento=1997-12-01" "$BASE_URL/jugadores"
echo -e "\nJugador 25 (ID: 7) cargado"

curl -X POST -d "agregarNombre=Miguel Borja&agregarPosicion=Delantero&agregarPais=Colombia&agregarNumero=9&agregarAltura=183&agregarFechaNacimiento=1993-01-26" "$BASE_URL/jugadores"
echo -e "\nJugador 26 (ID: 9) cargado"

curl -X POST -d "agregarNombre=Facundo Colidio&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=11&agregarAltura=175&agregarFechaNacimiento=2000-01-04" "$BASE_URL/jugadores"
echo -e "\nJugador 27 (ID: 11) cargado"

curl -X POST -d "agregarNombre=Sebastián Driussi&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=15&agregarAltura=179&agregarFechaNacimiento=1996-02-09" "$BASE_URL/jugadores"
echo -e "\nJugador 28 (ID: 15) cargado"

curl -X POST -d "agregarNombre=Bautista Dadín&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=27&agregarAltura=175&agregarFechaNacimiento=2006-05-20" "$BASE_URL/jugadores"
echo -e "\nJugador 29 (ID: 27) cargado"

curl -X POST -d "agregarNombre=Alex Woiski&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=29&agregarAltura=170&agregarFechaNacimiento=2006-03-17" "$BASE_URL/jugadores"
echo -e "\nJugador 30 (ID: 29) cargado"

curl -X POST -d "agregarNombre=Ian Subiabre&agregarPosicion=Delantero&agregarPais=Argentina&agregarNumero=38&agregarAltura=172&agregarFechaNacimiento=2007-01-01" "$BASE_URL/jugadores"
echo -e "\nJugador 31 (ID: 38) cargado"
echo


printf "\n vemos como quedan finalmente las tablas\n"
printf "\n tabla paises \n"
curl -sS "$BASE_URL/paises"; echo
printf "\n tabla jugadores \n"
curl -sS "$BASE_URL/jugadores"; echo