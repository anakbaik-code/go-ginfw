-- name: CountEvents :one
SELECT
    COUNT(*)
FROM
    events
WHERE
    deleted_at IS NULL;

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
        status
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListEventsByStatus :many
SELECT
    e.id,
    e.title,
    e.description,
    e.location,
    e.start_time,
    e.end_time,
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    -- 1. Hitung total kuota asli dari semua ticket_types di event ini
    COALESCE(
        (
            SELECT
                SUM(tt.quota)
            FROM
                ticket_types tt
            WHERE
                tt.event_id = e.id
        ),
        0
    ) -
    -- 2. Dikurangi jumlah tiket yang sudah sukses terjual (paid)
    COALESCE(
        (
            SELECT
                COUNT(*)
            FROM
                payments p
            WHERE
                p.event_id = e.id
                AND p.status = 'paid'
                AND p.deleted_at IS NULL
        ),
        0
    ) AS available_quota
FROM
    events e
    JOIN categories c ON e.category_id = c.id
    JOIN users u ON e.user_id = u.id
WHERE
    e.status = ?
    AND e.deleted_at IS NULL
    AND c.deleted_at IS NULL
ORDER BY
    e.start_time ASC
LIMIT
    ?
OFFSET
    ?;

-- name: ListMyEvents :many
SELECT
    e.id,
    e.title,
    e.description,
    e.location,
    e.start_time,
    e.end_time,
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    -- 1. Hitung total kuota asli dari semua ticket_types di event milik saya ini
    COALESCE(
        (
            SELECT
                SUM(tt.quota)
            FROM
                ticket_types tt
            WHERE
                tt.event_id = e.id
        ),
        0
    ) -
    -- 2. Dikurangi jumlah tiket yang sudah sukses terjual (paid)
    COALESCE(
        (
            SELECT
                COUNT(*)
            FROM
                payments p
            WHERE
                p.event_id = e.id
                AND p.status = 'paid'
                AND p.deleted_at IS NULL
        ),
        0
    ) AS available_quota
FROM
    events e
    JOIN categories c ON e.category_id = c.id
    JOIN users u ON e.user_id = u.id
WHERE
    e.user_id = ?
    AND e.deleted_at IS NULL
    AND c.deleted_at IS NULL
ORDER BY
    e.created_at DESC
LIMIT
    ?
OFFSET
    ?;

-- name: CountMyEvents :one
SELECT
    COUNT(*)
FROM
    events e
WHERE
    e.user_id = ?
    AND e.deleted_at IS NULL;

-- name: GetMyEventByID :one
SELECT
    e.id,
    e.title,
    e.description,
    e.location,
    e.start_time,
    e.end_time,
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    -- 1. Hitung total kuota dari semua jenis tiket di event ini
    COALESCE(
        (
            SELECT
                SUM(tt.quota)
            FROM
                ticket_types tt
            WHERE
                tt.event_id = e.id
        ),
        0
    ) -
    -- 2. Dikurangi total tiket yang sudah terjual (paid)
    COALESCE(
        (
            SELECT
                COUNT(*)
            FROM
                payments p
            WHERE
                p.event_id = e.id
                AND p.status = 'paid'
                AND p.deleted_at IS NULL
        ),
        0
    ) AS available_quota
FROM
    events e
    JOIN categories c ON e.category_id = c.id
    JOIN users u ON e.user_id = u.id
WHERE
    e.id = ?
    AND e.user_id = ?
    AND e.deleted_at IS NULL
    AND c.deleted_at IS NULL;

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
    e.status,
    e.created_at,
    c.name AS category_name,
    u.name AS organizer_name,
    -- 1. Hitung total kuota asli dari semua ticket_types di event ini
    COALESCE(
        (
            SELECT
                SUM(tt.quota)
            FROM
                ticket_types tt
            WHERE
                tt.event_id = e.id
        ),
        0
    ) -
    -- 2. Dikurangi jumlah tiket yang sudah sukses terjual (paid)
    COALESCE(
        (
            SELECT
                COUNT(*)
            FROM
                payments p
            WHERE
                p.event_id = e.id
                AND p.status = 'paid'
                AND p.deleted_at IS NULL
        ),
        0
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
    deleted_at = NOW
WHERE
    id = ?;

-- name: CountEventsByStatus :one
SELECT
    COUNT(*)
FROM
    events e
    JOIN categories c ON e.category_id = c.id
WHERE
    e.status = ?
    AND e.deleted_at IS NULL
    AND c.deleted_at IS NULL;