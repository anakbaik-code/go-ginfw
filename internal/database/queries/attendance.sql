-- ============================================
-- CREATE / CHECK-IN
-- ============================================
-- name: CheckIn :exec
INSERT INTO
    attendances (
        order_id,
        event_id,
        user_id,
        ticket_type_id,
        check_in_time,
        check_in_method,
        status,
        checked_by
    )
VALUES
    (?, ?, ?, ?, NOW(), ?, 'checked_in', ?) ON DUPLICATE KEY
UPDATE check_in_time = NOW(),
check_in_method =
VALUES
    (check_in_method),
    status = 'checked_in',
    checked_by =
VALUES
    (checked_by),
    updated_at = NOW();

-- name: CheckInByQR :exec
INSERT INTO
    attendances (
        order_id,
        event_id,
        user_id,
        ticket_type_id,
        check_in_time,
        check_in_method,
        status,
        checked_by
    )
VALUES
    (?, ?, ?, ?, NOW(), 'qr_code', 'checked_in', ?) ON DUPLICATE KEY
UPDATE check_in_time = NOW(),
check_in_method = 'qr_code',
status = 'checked_in',
checked_by =
VALUES
    (checked_by),
    updated_at = NOW();

-- name: ManualCheckIn :exec
INSERT INTO
    attendances (
        order_id,
        event_id,
        user_id,
        ticket_type_id,
        check_in_time,
        check_in_method,
        status,
        checked_by
    )
VALUES
    (?, ?, ?, ?, NOW(), 'manual', 'checked_in', ?) ON DUPLICATE KEY
UPDATE check_in_time = NOW(),
check_in_method = 'manual',
status = 'checked_in',
checked_by =
VALUES
    (checked_by),
    updated_at = NOW();

-- ============================================
-- READ
-- ============================================
-- name: GetAttendanceByOrderID :one
SELECT
    a.id,
    a.order_id,
    a.event_id,
    a.user_id,
    a.ticket_type_id,
    a.check_in_time,
    a.check_in_method,
    a.status,
    a.checked_by,
    a.created_at,
    a.updated_at,
    u.name AS user_name,
    u.email AS user_email,
    e.title AS event_title,
    tt.name AS ticket_type_name,
    c.name AS checked_by_name
FROM
    attendances a
    JOIN users u ON a.user_id = u.id
    JOIN events e ON a.event_id = e.id
    JOIN ticket_types tt ON a.ticket_type_id = tt.id
    LEFT JOIN users c ON a.checked_by = c.id
WHERE
    a.order_id = ?
    AND a.deleted_at IS NULL;

-- name: GetAttendanceByEventID :many
SELECT
    a.id,
    a.order_id,
    a.user_id,
    u.name AS user_name,
    u.email AS user_email,
    a.check_in_time,
    a.check_in_method,
    a.status,
    a.checked_by,
    c.name AS checked_by_name
FROM
    attendances a
    JOIN users u ON a.user_id = u.id
    LEFT JOIN users c ON a.checked_by = c.id
WHERE
    a.event_id = ?
    AND a.status = 'checked_in'
    AND a.deleted_at IS NULL
ORDER BY
    a.check_in_time DESC;

-- name: GetAttendanceByUserID :many
SELECT
    a.id,
    a.order_id,
    a.event_id,
    a.user_id,
    a.ticket_type_id,
    a.check_in_time,
    a.check_in_method,
    a.status,
    a.checked_by,
    a.created_at,
    e.title AS event_title,
    e.start_time,
    e.end_time
FROM
    attendances a
    JOIN events e ON a.event_id = e.id
WHERE
    a.user_id = ?
    AND a.deleted_at IS NULL
ORDER BY
    a.check_in_time DESC;

-- ============================================
-- STATISTICS
-- ============================================
-- name: CountAttendeesByEvent :one
SELECT
    COUNT(DISTINCT user_id)
FROM
    attendances
WHERE
    event_id = ?
    AND status = 'checked_in'
    AND deleted_at IS NULL;

-- name: CountAttendeesByOrganizer :one
SELECT
    COUNT(DISTINCT a.user_id)
FROM
    attendances a
    JOIN events e ON a.event_id = e.id
WHERE
    e.user_id = ?
    AND a.status = 'checked_in'
    AND a.deleted_at IS NULL;

-- name: CountAttendeesByTicketType :one
SELECT
    COUNT(*)
FROM
    attendances
WHERE
    ticket_type_id = ?
    AND status = 'checked_in'
    AND deleted_at IS NULL;

-- name: GetEventAttendanceRate :many
SELECT
    e.id AS event_id,
    e.title AS event_title,
    COUNT(DISTINCT o.id) AS total_orders,
    COUNT(DISTINCT a.user_id) AS attended,
    ROUND(
        COALESCE(
            COUNT(DISTINCT a.user_id) / NULLIF(COUNT(DISTINCT o.id), 0) * 100,
            0
        ),
        2
    ) AS attendance_rate
FROM
    events e
    LEFT JOIN orders o ON o.event_id = e.id
    AND o.status = 'paid'
    AND o.deleted_at IS NULL
    LEFT JOIN attendances a ON a.event_id = e.id
    AND a.status = 'checked_in'
    AND a.deleted_at IS NULL
WHERE
    e.user_id = ?
    AND e.deleted_at IS NULL
GROUP BY
    e.id,
    e.title
ORDER BY
    attendance_rate DESC;

-- name: GetCheckInTimeline :many
SELECT
    HOUR (check_in_time) AS hour,
    COUNT(*) AS count
FROM
    attendances
WHERE
    event_id = ?
    AND status = 'checked_in'
    AND deleted_at IS NULL
GROUP BY
    HOUR (check_in_time)
ORDER BY
    hour ASC;

-- name: GetCheckInTimelineByDate :many
SELECT
    DATE (check_in_time) AS date,
    COUNT(*) AS count
FROM
    attendances
WHERE
    event_id = ?
    AND status = 'checked_in'
    AND deleted_at IS NULL
GROUP BY
    DATE (check_in_time)
ORDER BY
    date ASC;

-- ============================================
-- UPDATE / DELETE
-- ============================================
-- name: CancelAttendance :exec
UPDATE attendances
SET
    status = 'cancelled',
    updated_at = NOW()
WHERE
    order_id = ?
    AND event_id = ?
    AND deleted_at IS NULL;

-- name: SoftDeleteAttendance :exec
UPDATE attendances
SET
    deleted_at = NOW(),
    updated_at = NOW()
WHERE
    id = ?;

-- ============================================
-- VALIDATION
-- ============================================
-- name: CheckAlreadyCheckedIn :one
SELECT
    COUNT(*)
FROM
    attendances
WHERE
    order_id = ?
    AND event_id = ?
    AND status = 'checked_in'
    AND deleted_at IS NULL;

-- name: GetOrderStatus :one
SELECT
    status
FROM
    orders
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: GetEventTimeRange :one
SELECT
    start_time,
    end_time
FROM
    events
WHERE
    id = ?
    AND deleted_at IS NULL;

-- CHECK
-- name: CheckAlreadyCheckedInForUpdate :one
SELECT
    status = 'checked_in' AS already_checked_in
FROM
    attendances
WHERE
    order_id = ?
    AND event_id = ?
    AND deleted_at IS NULL FOR
UPDATE;

-- name: CheckInByQRInsertOnly :exec
INSERT INTO
    attendances (
        order_id,
        event_id,
        user_id,
        ticket_type_id,
        check_in_time,
        check_in_method,
        status,
        checked_by
    )
VALUES
    (?, ?, ?, ?, NOW(), 'qr_code', 'checked_in', ?);

-- name: GetOrderItemDetail :one
SELECT
    oi.order_id,
    oi.event_id,
    oi.ticket_type_id,
    o.user_id
FROM
    order_items oi
    JOIN orders o ON o.id = oi.order_id
WHERE
    oi.id = ?;