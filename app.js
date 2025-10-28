const jugadores = document.getElementById("tabla-jugadores");

async function cargarProductos() {
    try {
        const response = await fetch('/jugadoresCompleto');

        if (!response.ok) {
            throw new Error(`Error HTTP: ${response.status}`);
        }

        const jugadores = await response.json();
        
        construirTabla(jugadores);

    } catch (error) {
        console.error("Error al cargar jugadores:", error);
    }
}

function construirTabla(jugadores) {
    const tbody = document.getElementById("tabla-jugadores");
    tbody.innerHTML = ""; // Se limpia la tabla por si acaso

    // Itera sobre los productos y crea las filas
    for (const producto of productos) {
        const fila = `
            <tr>
                <td>${producto.id}</td>
                <td>${producto.name}</td>
                <td>${producto.price}</td>
            </tr>
        `;
        tbody.innerHTML += fila; // Agrega la fila al cuerpo de la tabla
    }
}

// Llama a la función cuando el DOM esté listo
document.addEventListener('DOMContentLoaded', cargarProductos);