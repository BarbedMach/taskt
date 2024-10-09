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

function loadTasks() {
    const userID = localStorage.getItem('userID'); // Get the logged-in user's ID

    fetch(`http://localhost:8080/tasks?userID=${userID}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => response.json())
    .then(data => {
        const taskTableBody = document.querySelector('#taskTable tbody');
        taskTableBody.innerHTML = ''; // Clear existing tasks

        data.forEach(task => {
            const row = document.createElement('tr');
            row.innerHTML = `
                <td>${task.id}</td>
                <td>${task.title}</td>
                <td>${task.description || ''}</td>
                <td>${new Date(task.start_time).toLocaleString()}</td>
                <td>${new Date(task.end_time).toLocaleString()}</td>
                <td>${task.status ? 'Completed' : 'Not Completed'}</td>
            `;
            taskTableBody.appendChild(row);
        });
    })
    .catch((error) => {
        console.error('Error fetching tasks:', error);
    });
}

function addTask() {
    document.getElementById('addTaskForm').addEventListener('submit', function(event) {
        event.preventDefault();
        
        const title = document.getElementById('taskTitle').value;
        const description = document.getElementById('taskDescription').value;
        const startTime = document.getElementById('startTime').value;
        const endTime = document.getElementById('endTime').value;
        const userID = localStorage.getItem('userID');
    
        fetch('http://localhost:8080/tasks', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ title, description, startTime, endTime, status: false, userID }),
        })
        .then(response => response.json())
        .then(data => {
            if (data.success) {
                localStorage.setItem('userID', data.user.id);
                loadTasks();
            } else {
                alert('Failed to add task: ' + data.message);
            }
        })
        .catch((error) => {
            console.error('Error:', error);
        });
    });
}

// Function to log the user out
function logout() {
    // Implement logout functionality (clear session, redirect, etc.)
    window.location.href = 'login.html';
}

// Populate username and tasks on page load
document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('username').innerText = localStorage.getItem('username') || 'User';
    fetchTasks();
});

document.addEventListener('DOMContentLoaded', function() {
    loadTasks(); // Load tasks when the page is loaded
});