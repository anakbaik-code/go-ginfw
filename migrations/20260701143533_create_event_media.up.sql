CREATE TABLE
    event_media (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        event_id BIGINT UNSIGNED NOT NULL,
        image_path VARCHAR(255) NOT NULL, -- Nyimpen path file/URL foto
        is_primary BOOLEAN DEFAULT FALSE, -- Penanda poster utama buat di homepage
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_image_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
    );