CREATE TABLE attendances (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id BIGINT UNSIGNED NOT NULL,
    event_id BIGINT UNSIGNED NOT NULL,
    user_id BIGINT  UNSIGNED NOT NULL,
    ticket_type_id BIGINT UNSIGNED NOT NULL,
    check_in_time DATETIME,
    check_in_method ENUM('qr_code', 'manual', 'online') DEFAULT 'qr_code',
    status ENUM('pending', 'checked_in', 'cancelled') DEFAULT 'pending',
    checked_by BIGINT UNSIGNED,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (event_id) REFERENCES events(id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (ticket_type_id) REFERENCES ticket_types(id),
    FOREIGN KEY (checked_by) REFERENCES users(id),
    
    UNIQUE KEY unique_attendance (order_id, event_id),
    INDEX idx_event_checkin (event_id, status),
    INDEX idx_user_attendance (user_id, event_id)
);