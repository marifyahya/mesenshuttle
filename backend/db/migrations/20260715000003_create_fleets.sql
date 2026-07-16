-- +goose Up
CREATE TABLE IF NOT EXISTS fleets (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    plate_number VARCHAR(20) NOT NULL UNIQUE,
    type VARCHAR(50) NOT NULL,
    total_seats BIGINT NOT NULL,
    active_status BOOLEAN DEFAULT TRUE,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL
);

-- +goose Down
DROP TABLE IF EXISTS fleets;
