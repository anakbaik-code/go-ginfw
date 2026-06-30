CREATE TABLE
    events (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        category_id INT UNSIGNED NOT NULL,
        user_id BIGINT UNSIGNED NOT NULL,
        title VARCHAR(255) NOT NULL,
        description TEXT NOT NULL,
        location VARCHAR(255) NOT NULL,
        start_time TIMESTAMP NOT NULL,
        end_time TIMESTAMP NOT NULL,
        price INT UNSIGNED NOT NULL DEFAULT 0,
        quota INT UNSIGNED NOT NULL DEFAULT 0,
        status ENUM ('active', 'inactive', 'cancelled') DEFAULT 'active',
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL DEFAULT NULL, -- 👈 Soft Delete Event
        CONSTRAINT fk_event_category FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE RESTRICT,
        CONSTRAINT fk_event_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );