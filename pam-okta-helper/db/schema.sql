CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    session_token TEXT,
    is_sudoer BOOLEAN NOT NULL DEFAULT 0,
    last_login DATETIME
);

