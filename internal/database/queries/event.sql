-- name: CreateEvent :execresult
INSERT INTO
    events (
        category_id,
        user_id,
        title,
        description,
        location,
        start_time,
        end_time,
        price,
        quota,
        status
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListEventsActive :many
-- Hanya menampilkan event yang aktif, belum kedaluwarsa, dan belum di-soft delete
SELECT
    e.id,
    e.title,
    e.description,
    e.location,
    e.start_time,
    e.end_time,
    e.price,
    e.quota,
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    -- Menghitung sisa tiket secara real-time dari tiket yang berstatus 'paid'
    (
        e.quota - (
            SELECT
                COUNT(*)
            FROM
                tickets t
            WHERE
                t.event_id = e.id
                AND t.status = 'paid'
                AND t.deleted_at IS NULL
        )
    ) AS available_quota
FROM
    events e
    JOIN categories c ON e.category_id = c.id
    JOIN users u ON e.user_id = u.id
WHERE
    e.status = 'active'
    AND e.end_time > NOW()
    AND e.deleted_at IS NULL
    AND c.deleted_at IS NULL
ORDER BY
    e.start_time ASC;

-- name: GetEventByID :one
SELECT
    e.id,
    e.category_id,
    e.user_id,
    e.title,
    e.description,
    e.location,
    e.start_time,
    e.end_time,
    e.price,
    e.quota,
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    (
        e.quota - (
            SELECT
                COUNT(*)
            FROM
                tickets t
            WHERE
                t.event_id = e.id
                AND t.status = 'paid'
                AND t.deleted_at IS NULL
        )
    ) AS available_quota
FROM
    events e
    JOIN categories c ON e.category_id = c.id
    JOIN users u ON e.user_id = u.id
WHERE
    e.id = ?
    AND e.deleted_at IS NULL
LIMIT
    1;

-- name: UpdateEvent :exec
UPDATE events
SET
    category_id = ?,
    title = ?,
    description = ?,
    location = ?,
    start_time = ?,
    end_time = ?,
    price = ?,
    quota = ?,
    status = ?
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateEventStatus :exec
-- Dipakai khusus untuk pembatalan (cancel) event oleh admin/organizer
UPDATE events
SET
    status = ?
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteEvent :exec
-- Soft Delete Event
UPDATE events
SET
    deleted_at = NOW()
WHERE
    id = ?;