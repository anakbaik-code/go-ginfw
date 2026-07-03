CREATE TABLE
    tickets (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        order_item_id BIGINT UNSIGNED NOT NULL,
        ticket_code VARCHAR(100) NOT NULL UNIQUE,
        qr_code VARCHAR(255) NOT NULL ,
        status ENUM ('active', 'used', 'cancelled') NOT NULL DEFAULT 'active',
        checked_in_at TIMESTAMP NULL,
        CONSTRAINT fk_ticket_order_item FOREIGN KEY (order_item_id) REFERENCES order_items (id)
    );