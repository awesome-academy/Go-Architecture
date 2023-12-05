CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    created_on INT,
    modified_on INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id VARCHAR(255)
);