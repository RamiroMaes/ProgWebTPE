const jugadores = document.getElementById("tabla-jugadores");

async function cargarPlantel() {
    try {
        const response = await fetch('/plantel');

        if (!response.ok) {
            throw new Error(`Error HTTP: ${response.status}`);
        }

        const jugadores = await response.json();
        
        construirTabla(jugadores);

    } catch (error) {
        console.error("Error al cargar jugadores:", error);
    }
}

// Funcion para eliminar la hora en la fecha de nacimiento.
function formatearFecha(fechaString) {
    return fechaString.split('T')[0]; // Toma solo la parte antes de la 'T'
}

function construirTabla(jugadores) {
    const tabla = document.getElementById("tabla-jugadores");
    const tbody = tabla.querySelector("tbody");
    
    tbody.innerHTML = ""; // Se limpia la tabla por si acaso

    // Itera sobre los jugadores y crea las filas
    for (const jugador of jugadores) {
        const fila = `
            <tr>
                <td>${jugador.posicion}</td>
                <td>${jugador.id_jugador}</td>
                <td>${jugador.nombre}</td>
                <td>${jugador.pais_nombre}</td>
                <td>${formatearFecha(jugador.fecha_nacimiento)}</td>
                <td>${jugador.altura}</td>
            </tr>
        `;
        tbody.innerHTML += fila; // Agrega la fila al cuerpo de la tabla
    }
}

// Llama a la función cuando el DOM esté listo
document.addEventListener('DOMContentLoaded', cargarPlantel);

