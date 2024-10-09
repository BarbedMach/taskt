CREATE TABLE IF NOT EXISTS tasks (
    task_id SERIAL PRIMARY KEY,            -- Unique identifier for each task
    title VARCHAR(100) NOT NULL,           -- Task title, required
    description TEXT,                      -- Task description
    start_time TIMESTAMP NOT NULL,         -- Start time of the task
    end_time TIMESTAMP NOT NULL,           -- End time of the task
    status BOOLEAN NOT NULL DEFAULT FALSE, -- Status of the task (completed or not)
    user_id INT NOT NULL,                  -- User ID from users table (foreign key)
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE  -- Foreign key to users table
);