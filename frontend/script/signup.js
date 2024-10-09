document.getElementById('signupForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, email, password }),
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            alert('Signup successful! Please log in.');
            window.location.href = "login.html";
        } else {
            alert('Signup failed: ' + data.message);
        }
    })
    .catch((error) => {
        console.error('Error:', error);
    });
});