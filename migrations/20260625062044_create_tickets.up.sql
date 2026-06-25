CREATE TABLE
    tickets (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        event_id BIGINT UNSIGNED NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        ticket_code VARCHAR(100) NOT NULL UNIQUE,
        status ENUM ('booked', 'paid', 'cancelled', 'used') DEFAULT 'booked',
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL DEFAULT NULL, -- 👈 Soft Delete Tiket
        CONSTRAINT fk_ticket_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE,
        CONSTRAINT fk_ticket_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );  