-- +goose Up
CREATE TABLE IF NOT EXISTS admins (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL
);

-- +goose Down
DROP TABLE IF EXISTS admins;
