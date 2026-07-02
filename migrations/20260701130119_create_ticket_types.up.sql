CREATE TABLE
    ticket_types (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        event_id BIGINT UNSIGNED NOT NULL,
        name VARCHAR(100) NOT NULL,
        price BIGINT UNSIGNED NOT NULL DEFAULT 0,
        quota INT UNSIGNED NOT NULL,
        max_per_transaction INT UNSIGNED DEFAULT 5,
        start_sale_at TIMESTAMP NULL,
        end_sale_at TIMESTAMP NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP DEFAULT NULL,
        CONSTRAINT fk_ticket_type_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
    );