-- name: CreateOrder :execresult
INSERT INTO orders (
    user_id,
    total_amount,
    status
)
VALUES (
    ?, ?, ?
);

-- name: GetOrderByID :one
SELECT
    id,
    user_id,
    total_amount,
    status,
    created_at,
    updated_at,
    deleted_at
FROM orders
WHERE
    id = ?
    AND deleted_at IS NULL
LIMIT 1;

-- name: ListOrdersByUserID :many
SELECT
    id,
    user_id,
    total_amount,
    status,
    created_at,
    updated_at,
    deleted_at
FROM orders
WHERE
    user_id = ?
    AND deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT ?
OFFSET ?;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET
    status = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteOrder :exec
UPDATE orders
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: CountOrdersByUserID :one
SELECT COUNT(*)
FROM orders
WHERE
    user_id = ?
    AND deleted_at IS NULL;

-- name: CountOrders :one
SELECT COUNT(*)
FROM orders
WHERE
    deleted_at IS NULL;