// Refer to internal/handlers/todo.go for the backend side
const API_URL = `tasks`

export async function getTodos() {
    const response = await fetch(API_URL);
    if (!response.ok) throw new Error('Failed to fetch todos');
    return response.json();
}

export async function addTodo(task) {
    const response = await fetch(API_URL, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(task),
    });
    if (!response.ok) throw new Error('Failed to add todo');
}

export async function updateTodo(task) {
    const response = await fetch(`${API_URL}/${task.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(task),
    });
    if (!response.ok) throw new Error('Failed to add todo');
}

export async function deleteTodo(id) {
    const response = await fetch(`${API_URL}/${id}`, {
      method: 'DELETE',
    });
    if (!response.ok) throw new Error('Failed to delete todo');
}