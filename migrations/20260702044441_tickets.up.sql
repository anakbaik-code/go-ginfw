CREATE TABLE
    tickets (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        order_item_id BIGINT UNSIGNED NOT NULL,
        ticket_code VARCHAR(100) UNIQUE,
        qr_code VARCHAR(255),
        status VARCHAR(20) DEFAULT 'active',
        checked_in_at TIMESTAMP NULL,
        CONSTRAINT fk_ticket_order_item FOREIGN KEY (order_item_id) REFERENCES order_items (id)
    );