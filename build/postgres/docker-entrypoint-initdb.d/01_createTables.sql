CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(256) NOT NULL,
    mail VARCHAR(256) NOT NULL,
    password VARCHAR(512) NOT NULL,
    role VARCHAR(256) NOT NULL
);

CREATE TABLE IF NOT EXISTS notes (
    id SERIAL PRIMARY KEY,
    owner INT NOT NULL,
    title VARCHAR(256),
    body TEXT,
    date DATE,
    FOREIGN KEY (owner) REFERENCES users(id)
);