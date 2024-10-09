// Fetch user tasks from the backend and display in the table
function fetchTasks() {
    fetch('http://localhost:8080/tasks', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        }
    })
    .then(response => response.json())
    .then(data => {
        const taskTableBody = document.getElementById('taskTable').querySelector('tbody');
        taskTableBody.innerHTML = ''; // Clear any existing rows
        
        data.forEach(task => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${task.id}</td>
                <td>${task.title}</td>
                <td>${task.description}</td>
                <td>${new Date(task.start_time).toLocaleString()}</td>
                <td>${new Date(task.end_time).toLocaleString()}</td>
                <td>${task.status === 'completed' ? '✅ Completed' : '⏳ Pending'}</td>
            `;
            taskTableBody.appendChild(row);
        });
    })
    .catch(error => {
        console.error('Error fetching tasks:', error);
    });
}

// Function to log the user out
function logout() {
    // Implement logout functionality (clear session, redirect, etc.)
    alert('You have been logged out!');
    window.location.href = 'login.html';
}

// Populate username and tasks on page load
document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('username').innerText = localStorage.getItem('username') || 'User';
    fetchTasks();
});