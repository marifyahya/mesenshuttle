-- +goose Up
CREATE TABLE IF NOT EXISTS bookings (
    id CHAR(36) PRIMARY KEY,
    schedule_id CHAR(36) NOT NULL,
    booking_code VARCHAR(20) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    status TINYINT NOT NULL DEFAULT 0,
    total_price BIGINT NOT NULL,
    payment_url VARCHAR(255) NULL,
    created_at DATETIME(3) NULL,
    updated_at DATETIME(3) NULL,
    CONSTRAINT fk_bookings_schedule FOREIGN KEY (schedule_id) REFERENCES schedules(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    INDEX idx_bookings_schedule_id (schedule_id),
    INDEX idx_bookings_email (email),
    INDEX idx_bookings_phone (phone),
    INDEX idx_bookings_status (status),
    INDEX idx_bookings_created_at (created_at)
);

-- +goose Down
DROP TABLE IF EXISTS bookings;
