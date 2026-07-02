-- name: CreateTicketType :execresult
INSERT INTO
    ticket_types (
        event_id,
        name,
        price,
        quota,
        max_per_transaction,
        start_sale_at,
        end_sale_at
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: GetTicketTypeByID :one
SELECT
    id,
    event_id,
    name,
    price,
    quota,
    max_per_transaction,
    start_sale_at,
    end_sale_at,
    created_at,
    updated_at
FROM
    ticket_types
WHERE
    id = ?
    AND deleted_at IS NULL
LIMIT
    1;

-- name: ListTicketTypesByEventID :many
SELECT
    id,
    event_id,
    name,
    price,
    quota,
    max_per_transaction,
    start_sale_at,
    end_sale_at,
    created_at,
    updated_at
FROM
    ticket_types
WHERE
    event_id = ?
    AND deleted_at IS NULL
ORDER BY
    created_at ASC;

-- name: UpdateTicketType :exec
UPDATE ticket_types
SET
    name = ?,
    price = ?,
    quota = ?,
    max_per_transaction = ?,
    start_sale_at = ?,
    end_sale_at = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteTicketType :exec
UPDATE ticket_types
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: CountTicketTypesByEventID :one
SELECT
    COUNT(*)
FROM
    ticket_types
WHERE
    event_id = ?
    AND deleted_at IS NULL;

-- name: ListActiveTicketTypes :many
SELECT
    id,
    event_id,
    name,
    price,
    quota,
    max_per_transaction,
    start_sale_at,
    end_sale_at,
    created_at,
    updated_at
FROM
    ticket_types
WHERE
    event_id = ?
    AND deleted_at IS NULL
    AND (
        start_sale_at IS NULL
        OR start_sale_at <= ?
    )
    AND (
        end_sale_at IS NULL
        OR end_sale_at >= ?
    )
ORDER BY
    price ASC;

-- name: GetAvailableTicketTypeByID :one
SELECT
    id,
    event_id,
    name,
    price,
    quota,
    max_per_transaction,
    start_sale_at,
    end_sale_at,
    created_at,
    updated_at
FROM
    ticket_types
WHERE
    id = ?
    AND deleted_at IS NULL
    AND (
        start_sale_at IS NULL
        OR start_sale_at <= ?
    )
    AND (
        end_sale_at IS NULL
        OR end_sale_at >= ?
    )
LIMIT
    1;