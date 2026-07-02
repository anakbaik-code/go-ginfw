-- name: CreateEventMedia :execresult
INSERT INTO
    event_media (event_id, image_path, is_primary)
VALUES
    (?, ?, ?);

-- name: GetEventMediaByID :one
SELECT
    id,
    event_id,
    image_path,
    is_primary,
    created_at
FROM event_media
WHERE
    id = ?
LIMIT 1;

-- name: GetEventPrimaryMedia :one
SELECT
    id,
    event_id,
    image_path,
    is_primary,
    created_at
FROM
    event_media
WHERE
    event_id = ?
    AND is_primary = TRUE
LIMIT
    1;

-- name: ListEventMedia :many
SELECT
    id,
    event_id,
    image_path,
    is_primary,
    created_at
FROM
    event_media
WHERE
    event_id = ?
ORDER BY
    is_primary DESC,
    created_at ASC;

-- name: ResetEventPrimaryMedia :exec
UPDATE event_media
SET
    is_primary = FALSE
WHERE
    event_id = ?;

-- name: SetEventPrimaryMedia :exec
UPDATE event_media
SET
    is_primary = TRUE
WHERE
    id = ?
    AND event_id = ?;

-- name: DeleteEventMedia :exec
DELETE FROM event_media
WHERE
    id = ?;