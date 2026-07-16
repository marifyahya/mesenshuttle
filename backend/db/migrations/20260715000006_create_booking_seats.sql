-- +goose Up
CREATE TABLE IF NOT EXISTS booking_seats (
    id CHAR(36) PRIMARY KEY,
    booking_id CHAR(36) NOT NULL,
    seat_number VARCHAR(10) NOT NULL,
    CONSTRAINT fk_bookings_booking_seats FOREIGN KEY (booking_id) REFERENCES bookings(id) ON UPDATE CASCADE ON DELETE CASCADE,
    INDEX idx_booking_seats_booking_id (booking_id)
);

-- +goose Down
DROP TABLE IF EXISTS booking_seats;
