-- name: CreateUser :execresult
INSERT INTO
    users (
        name,
        email,
        password_hash,
        phone,
        address,
        role,
        is_active,
        created_at,
        updated_at
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, (NOW()), (NOW()));

-- name: GetUserByID :one
SELECT
    id,
    name,
    email,
    phone,
    address,
    role,
    is_active,
    created_at,
    updated_at
FROM
    users
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: GetUserByEmail :one
SELECT
    id,
    name,
    email,
    password_hash,
    phone,
    address,
    role,
    is_active,
    created_at,
    updated_at
FROM
    users
WHERE
    email = ?
    AND deleted_at IS NULL;

-- name: ListUsers :many
SELECT
    id,
    name,
    email,
    phone,
    address,
    role,
    is_active,
    created_at,
    updated_at
FROM
    users
WHERE
    deleted_at IS NULL
ORDER BY
    id DESC
LIMIT
    ?
OFFSET
    ?;

-- name: UpdateUser :exec
UPDATE users
SET
    name = ?,
    phone = ?,
    address = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateUserEmail :exec
UPDATE users
SET
    email = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateUserPassword :exec
UPDATE users
SET
    password_hash = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateUserRole :exec
UPDATE users
SET
    role = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: UpdateUserStatus :exec
UPDATE users
SET
    is_active = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteUser :exec
UPDATE users
SET
    deleted_at = (NOW()),
    updated_at = (NOW())
WHERE
    id = ?;

-- name: HardDeleteUser :exec
DELETE FROM users
WHERE
    id = ?;

-- name: SearchUsers :many
SELECT
    id,
    name,
    email,
    phone,
    address,
    role,
    is_active,
    created_at,
    updated_at
FROM
    users
WHERE
    deleted_at IS NULL
    AND (
        name LIKE ?
        OR email LIKE ?
    )
ORDER BY
    id DESC
LIMIT
    ?
OFFSET
    ?;

-- name: GetUserByRefreshToken :one
SELECT
    id,
    name,
    email,
    role,
    is_active
FROM
    users
WHERE
    refresh_token = ?
    AND deleted_at IS NULL;

-- name: UpdateUserRefreshToken :exec
UPDATE users
SET
    refresh_token = ?,
    updated_at = (NOW())
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: ListActiveUsers :many
SELECT
    id,
    name,
    email,
    role,
    created_at
FROM
    users
WHERE
    is_active = true
    AND deleted_at IS NULL
ORDER BY
    created_at DESC;

-- name: CountOrganizerAll :one
SELECT
    COUNT(*) as total
FROM
    users
WHERE
    role = 'organizer'
    AND deleted_at IS NULL;

-- name: CountOrganizerActive :one
SELECT
    COUNT(*) as total
FROM
    users
WHERE
    role = 'organizer'
    AND deleted_at IS NULL
    AND is_active = true;

-- name: CountUsersActive :one
SELECT
    COUNT(*) as total
FROM
    users
WHERE
    role = 'user'
    AND deleted_at IS NULL
    AND is_active = true;

-- name: ListOrganizers :many
SELECT
    id,
    name,
    email,
    phone,
    address,
    role,
    is_active,
    created_at,
    updated_at,
    deleted_at
FROM
    users
WHERE
    role = 'organizer'
    AND deleted_at IS NULL
ORDER BY
    created_at DESC
LIMIT
    ?
OFFSET
    ?;
-- name: CountUsers :one
SELECT
    COUNT(*) as total
FROM
    users
WHERE
    role = 'user'
    AND deleted_at IS NULL;