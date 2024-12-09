const API_BASE_URL = "http://localhost:8080";

// Función para obtener todos los usuarios
async function getUsers() {
    try {
        const response = await fetch(`${API_BASE_URL}/users`);
        const data = await response.json();
        document.getElementById('response').textContent = JSON.stringify(data, null, 2);
    } catch (error) {
        document.getElementById('response').textContent = `Error: ${error.message}`;
    }
}

// Función para crear un nuevo usuario
async function createUser() {
    const username = document.getElementById('username').value;
    if (!username) {
        alert("Por favor, escribe un nombre");
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/users/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: username }),
        });
        const data = await response.json();
        document.getElementById('response').textContent = JSON.stringify(data, null, 2);
    } catch (error) {
        document.getElementById('response').textContent = `Error: ${error.message}`;
    }
}