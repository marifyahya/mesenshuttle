-- +goose Up
CREATE TABLE IF NOT EXISTS routes (
    id CHAR(36) PRIMARY KEY,
    origin_city VARCHAR(100) NOT NULL,
    origin_pool VARCHAR(150) NOT NULL,
    destination_city VARCHAR(100) NOT NULL,
    destination_pool VARCHAR(150) NOT NULL,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    INDEX idx_routes_origin_city (origin_city),
    INDEX idx_routes_destination_city (destination_city)
);

-- +goose Down
DROP TABLE IF EXISTS routes;
