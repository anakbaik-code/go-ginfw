-- name: CreateOrderItem :execresult
INSERT INTO
    order_items (
        order_id,
        event_id,
        ticket_type_id,
        quantity,
        price
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetOrderItemByID :one
SELECT
    id,
    order_id,
    event_id,
    ticket_type_id,
    quantity,
    price
FROM
    order_items
WHERE
    id = ?
LIMIT
    1;

-- name: ListOrderItemsByOrderID :many
SELECT
    id,
    order_id,
    event_id,
    ticket_type_id,
    quantity,
    price
FROM
    order_items
WHERE
    order_id = ?
ORDER BY
    id ASC;

-- name: DeleteOrderItem :exec
DELETE FROM order_items
WHERE
    id = ?;

-- name: DeleteOrderItemsByOrderID :exec
DELETE FROM order_items
WHERE
    order_id = ?;

-- name: CountOrderItemsByOrderID :one
SELECT
    COUNT(*)
FROM
    order_items
WHERE
    order_id = ?;