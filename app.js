const jugadores = document.getElementById("tabla-jugadores");

// Funcion para cargar los jugadores desde la base de datos.
async function cargarPlantel() {
    try {
        const response = await fetch('/plantel');

        if (!response.ok) {
            throw new Error(`Error HTTP: ${response.status}`);
        }

        const jugadoresData = await response.json();

        construirTabla(jugadoresData);

    } catch (error) {
        console.error("Error al cargar jugadores:", error);
    }
}

// Funcion para formatear la fecha de nacimiento correctamente.
function formatearFecha(fechaString) {
    return fechaString.split('T')[0]; // Toma solo la parte antes de la 'T'.
}

function construirTabla(jugadores) {
    const tabla = document.getElementById("tabla-jugadores");
    const tbody = tabla.querySelector("tbody");
    
    tbody.innerHTML = ""; // Se limpia la tabla por si acaso.

    // Itera sobre los jugadores y crea las filas.
    for (const jugador of jugadores) {
        const fila = `
            <tr data-id="${jugador.id_jugador}">    <!-- Data lo utilizamos para luego obtener el jugador que se desea eliminar. -->
                <td>${jugador.posicion}</td>
                <td>${jugador.id_jugador}</td>
                <td>${jugador.nombre}</td>
                <td>${jugador.pais_nombre}</td>
                <td>${formatearFecha(jugador.fecha_nacimiento)}</td>
                <td>${jugador.altura}</td>
                <td class=celda-eliminar>
                    <button class= "eliminar" type="button">Eliminar</button>
                </td>
            </tr>
        `;
        tbody.innerHTML += fila; // Agrega la fila al cuerpo de la tabla.
    }
}

// Función para agregar UN jugador a la tabla (sin recargar todo).
function agregarFilaATabla(jugador) {
    const tabla = document.getElementById("tabla-jugadores");
    const tbody = tabla.querySelector("tbody");
    
    const fila = `
        <tr data-id="${jugador.id_jugador}"> 
            <td>${jugador.posicion}</td>
            <td>${jugador.id_jugador}</td>
            <td>${jugador.nombre}</td>
            <td>${jugador.pais_nombre}</td>
            <td>${formatearFecha(jugador.fecha_nacimiento)}</td>
            <td>${jugador.altura}</td>
            <td>
                <button class="eliminar" type="button">Eliminar</button>
            </td>
        </tr>
    `;
    tbody.innerHTML += fila;
}

// Función para manejar el envío del formulario.
async function enviarFormulario(event) {
    event.preventDefault(); // Previene el envío por defecto (que recarga la página).

    // Capturar los valores del formulario
    const fechaInput = document.getElementById('fecha_nacimiento').value;

    // Hacer lo inverso a lo hecho previamente con la fecha de nacimiento.
    const fechaISO = fechaInput + "T00:00:00Z";

    // Capturar los valores del formulario.
    const nuevoJugador = {
        posicion: document.getElementById('posicion').value,
        id_jugador: parseInt(document.getElementById('numero').value),
        nombre: document.getElementById('nombre').value,
        pais_nombre: document.getElementById('pais').value,
        fecha_nacimiento: fechaISO,
        altura: parseInt(document.getElementById('altura').value)
    };
    
    try {
        // Enviar con POST a la API.
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

        // Leer la respuesta UNA SOLA VEZ.
        const jugadorCreado = await response.json();
        console.log('Jugador creado:', jugadorCreado);
        
        agregarFilaATabla(jugadorCreado);
        
        // Limpiar el formulario.
        event.target.reset();
        
        alert('Jugador agregado exitosamente!');

    } catch (error) {
        console.error('Error al agregar jugador:', error);
        alert('Error al agregar el jugador. Ver consola para más detalles.');
    }
}

// Funcion para eliminar un jugador de la tabla.
async function eliminarEntrada(e) {
    if (!e.target.closest('.eliminar')) return; // Condicion aceptada si se hace click en boton eliminar.

    const row = e.target.closest('tr'); // Para esto usamos data en la fila.
    const id = row.dataset.id; 
    if (!id) return;

    if (!confirm(`Desea eliminar al jugador número ${id}?`)) return;

    try {
        const res = await fetch(`/jugadores/${encodeURIComponent(id)}`, { method: 'DELETE' });
        if (!res.ok) throw new Error(`HTTP ${res.status}`);
        row.remove();
    } catch (err) {
        console.error('Error eliminando jugador:', err);
        alert('No se pudo eliminar el jugador.');
    }
}

// Event listeners.
function inicializar() {
    cargarPlantel();

    // Listener para agregar un jugador via formulario.
    const form = document.querySelector('form');
    form.addEventListener('submit', enviarFormulario);

    // Listener para eliminar un jugador via click en el boton.
    const tbody = jugadores.querySelector('tbody');
    tbody.addEventListener('click', eliminarEntrada);
}

// Se ejecuta la funcion inicializar una vez se haya cargado el DOM.
document.addEventListener('DOMContentLoaded', inicializar);