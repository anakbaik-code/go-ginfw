CREATE TABLE order_items (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT UNSIGNED NOT NULL,
    event_id BIGINT UNSIGNED NOT NULL,
    ticket_type_id BIGINT UNSIGNED NOT NULL,
    quantity INT UNSIGNED NOT NULL,
    price INT UNSIGNED NOT NULL,
    CONSTRAINT fk_order_item_order
        FOREIGN KEY (order_id) REFERENCES orders(id),

    CONSTRAINT fk_order_item_event
        FOREIGN KEY (event_id) REFERENCES events(id),

    CONSTRAINT fk_order_item_ticket
        FOREIGN KEY (ticket_type_id) REFERENCES ticket_types(id)
);