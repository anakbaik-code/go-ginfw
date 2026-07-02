-- name: CreateNotification :execresult
INSERT INTO
    notifications (
        user_id,
        title,
        message,
        type,
        is_read,
        reference_type,
        reference_id
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?);

-- name: GetNotificationByID :one
SELECT
    id,
    user_id,
    title,
    message,
    type,
    is_read,
    reference_type,
    reference_id,
    created_at,
    updated_at,
    deleted_at
FROM
    notifications
WHERE
    id = ?
    AND deleted_at IS NULL
LIMIT
    1;

-- name: ListNotificationsByUserID :many
SELECT
    id,
    user_id,
    title,
    message,
    type,
    is_read,
    reference_type,
    reference_id,
    created_at,
    updated_at,
    deleted_at
FROM
    notifications
WHERE
    user_id = ?
    AND deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT
    ?
OFFSET
    ?;

-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET
    is_read = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: MarkAllNotificationsAsRead :exec
UPDATE notifications
SET
    is_read = TRUE,
    updated_at = CURRENT_TIMESTAMP
WHERE
    user_id = ?
    AND deleted_at IS NULL;

-- name: CountUnreadNotifications :one
SELECT
    COUNT(*)
FROM
    notifications
WHERE
    user_id = ?
    AND is_read = FALSE
    AND deleted_at IS NULL;

-- name: DeleteNotification :exec
UPDATE notifications
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteNotificationsByUserID :exec
UPDATE notifications
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    user_id = ?
    AND deleted_at IS NULL;