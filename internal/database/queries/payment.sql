-- name: CreatePayment :execresult
INSERT INTO payments (
    order_id,
    payment_code,
    amount,
    payment_method,
    payment_provider,
    provider_transaction_id,
    status,
    paid_at,
    expired_at
)
VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: GetPaymentByID :one
SELECT
    id,
    order_id,
    payment_code,
    amount,
    payment_method,
    payment_provider,
    provider_transaction_id,
    status,
    paid_at,
    expired_at,
    created_at,
    updated_at,
    deleted_at
FROM payments
WHERE
    id = ?
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetPaymentByCode :one
SELECT
    id,
    order_id,
    payment_code,
    amount,
    payment_method,
    payment_provider,
    provider_transaction_id,
    status,
    paid_at,
    expired_at,
    created_at,
    updated_at,
    deleted_at
FROM payments
WHERE
    payment_code = ?
    AND deleted_at IS NULL
LIMIT 1;

-- name: GetPaymentByOrderID :one
SELECT
    id,
    order_id,
    payment_code,
    amount,
    payment_method,
    payment_provider,
    provider_transaction_id,
    status,
    paid_at,
    expired_at,
    created_at,
    updated_at,
    deleted_at
FROM payments
WHERE
    order_id = ?
    AND deleted_at IS NULL
LIMIT 1;

-- name: ListPayments :many
SELECT
    id,
    order_id,
    payment_code,
    amount,
    payment_method,
    payment_provider,
    provider_transaction_id,
    status,
    paid_at,
    expired_at,
    created_at,
    updated_at,
    deleted_at
FROM payments
WHERE
    deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT ?
OFFSET ?;

-- name: ListPaymentsByUserID :many
SELECT
    p.id,
    p.order_id,
    p.payment_code,
    p.amount,
    p.payment_method,
    p.payment_provider,
    p.provider_transaction_id,
    p.status,
    p.paid_at,
    p.expired_at,
    p.created_at,
    p.updated_at,
    p.deleted_at
FROM payments p
JOIN orders o ON o.id = p.order_id
WHERE
    o.user_id = ?
    AND p.deleted_at IS NULL
ORDER BY
    p.created_at DESC
LIMIT ?
OFFSET ?;

-- name: UpdatePaymentStatus :exec
UPDATE payments
SET
    status = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: MarkPaymentAsPaid :exec
UPDATE payments
SET
    status = 'paid',
    paid_at = CURRENT_TIMESTAMP,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateProviderTransactionID :exec
UPDATE payments
SET
    provider_transaction_id = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: CountPayments :one
SELECT COUNT(*)
FROM payments
WHERE
    deleted_at IS NULL;

-- name: DeletePayment :exec
UPDATE payments
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;