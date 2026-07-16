-- +goose Up
CREATE TABLE IF NOT EXISTS schedules (
    id CHAR(36) PRIMARY KEY,
    route_id CHAR(36) NOT NULL,
    fleet_id CHAR(36) NOT NULL,
    departure_time DATETIME(3) NOT NULL,
    arrival_time DATETIME(3) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    status VARCHAR(20) DEFAULT 'Scheduled',
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    CONSTRAINT fk_schedules_route FOREIGN KEY (route_id) REFERENCES routes(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT fk_schedules_fleet FOREIGN KEY (fleet_id) REFERENCES fleets(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    INDEX idx_schedules_route_id (route_id),
    INDEX idx_schedules_fleet_id (fleet_id)
);

-- +goose Down
DROP TABLE IF EXISTS schedules;
