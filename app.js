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

// Función para agregar UN jugador a la tabla (sin recargar todo)
function agregarFilaATabla(jugador) {
    const tabla = document.getElementById("tabla-jugadores");
    const tbody = tabla.querySelector("tbody");
    
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
    tbody.innerHTML += fila;
}

// Función para manejar el envío del formulario
async function enviarFormulario(event) {
    event.preventDefault(); // Previene el envío por defecto (que recarga la página)

    // Capturar los valores del formulario
    const fechaInput = document.getElementById('fecha_nacimiento').value;
    // Convertir "2025-10-03" a "2025-10-03T00:00:00Z"
    const fechaISO = fechaInput + "T00:00:00Z";

    // Capturar los valores del formulario
    const nuevoJugador = {
        posicion: document.getElementById('posicion').value,
        id_jugador: parseInt(document.getElementById('numero').value),
        nombre: document.getElementById('nombre').value,
        pais_nombre: document.getElementById('pais').value,
        fecha_nacimiento: fechaISO,
        altura: parseInt(document.getElementById('altura').value)
    };
    
    // DEBUG: Ver qué estamos enviando
    console.log('Datos a enviar:', nuevoJugador);

    try {
        // Enviar POST a la API
        const response = await fetch('/jugadores', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(nuevoJugador)
        });

        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`Error HTTP: ${response.status} - ${errorText}`);
        }

        // Leer la respuesta UNA SOLA VEZ
        const jugadorCreado = await response.json();
        console.log('Jugador creado:', jugadorCreado);
        
        // Agregar el jugador a la tabla
        agregarFilaATabla(jugadorCreado);
        
        // Limpiar el formulario
        event.target.reset();
        
        alert('Jugador agregado exitosamente!');

    } catch (error) {
        console.error('Error al agregar jugador:', error);
        alert('Error al agregar el jugador. Ver consola para más detalles.');
    }
}

// Event listeners
function inicializar() {
    cargarPlantel();
    
    // Capturar el formulario y agregar el listener
    const form = document.querySelector('form');
    form.addEventListener('submit', enviarFormulario);
}

document.addEventListener('DOMContentLoaded', inicializar);