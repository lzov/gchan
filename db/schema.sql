-- Tabla de usuarios
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username TEXT NOT NULL UNIQUE CHECK(length(username) BETWEEN 3 AND 50)
);

-- Tabla de boards
CREATE TABLE boards (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT
);

-- Tabla de threads
CREATE TABLE threads (
    id INTEGER PRIMARY KEY,
    subject TEXT,
    board_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (board_id) REFERENCES boards(id)
);

-- Tabla de posts
CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    comment TEXT NOT NULL,
    file TEXT,  -- path o nombre de archivo (opcional)
    creation_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER,  -- relación opcional con la tabla users
    thread_id INTEGER NOT NULL,
    FOREIGN KEY (thread_id) REFERENCES threads(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Índices recomendados para rendimiento
CREATE INDEX idx_threads_board_id ON threads(board_id);
CREATE INDEX idx_posts_thread_id ON posts(thread_id);

