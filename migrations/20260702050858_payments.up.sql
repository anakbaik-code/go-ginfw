CREATE TABLE
    payments (
        id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        order_id BIGINT UNSIGNED NOT NULL,
        payment_code VARCHAR(100) NOT NULL UNIQUE,
        amount INT UNSIGNED NOT NULL,
        payment_method VARCHAR(50) NOT NULL,
        payment_provider VARCHAR(50) NOT NULL,
        provider_transaction_id VARCHAR(255) NULL,
        status ENUM (
            'pending',
            'paid',
            'failed',
            'expired',
            'cancelled',
            'refunded'
        ) NOT NULL DEFAULT 'pending',
        paid_at TIMESTAMP NULL,
        expired_at TIMESTAMP NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        deleted_at TIMESTAMP NULL DEFAULT NULL,
        CONSTRAINT fk_payment_order FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE
    );