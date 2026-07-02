-- name: CreateTicket :execresult
INSERT INTO
    tickets (
        order_item_id,
        ticket_code,
        qr_code,
        status,
        checked_in_at
    )
VALUES
    (?, ?, ?, ?, ?);

-- name: GetTicketByID :one
SELECT
    id,
    order_item_id,
    ticket_code,
    qr_code,
    status,
    checked_in_at
FROM
    tickets
WHERE
    id = ?
LIMIT
    1;

-- name: GetTicketByCode :one
SELECT
    id,
    order_item_id,
    ticket_code,
    qr_code,
    status,
    checked_in_at
FROM
    tickets
WHERE
    ticket_code = ?
LIMIT
    1;

-- name: ListTicketsByOrderItemID :many
SELECT
    id,
    order_item_id,
    ticket_code,
    qr_code,
    status,
    checked_in_at
FROM
    tickets
WHERE
    order_item_id = ?
ORDER BY
    id ASC;

-- name: UpdateTicketQRCode :exec
UPDATE tickets
SET
    qr_code = ?
WHERE
    id = ?;

-- name: UpdateTicketStatus :exec
UPDATE tickets
SET
    status = ?
WHERE
    id = ?;

-- name: CheckInTicket :exec
UPDATE tickets
SET
    status = 'used',
    checked_in_at = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: CountTicketsByOrderItemID :one
SELECT
    COUNT(*)
FROM
    tickets
WHERE
    order_item_id = ?;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE
    id = ?;