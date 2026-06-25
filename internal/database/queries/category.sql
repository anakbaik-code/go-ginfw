-- name: CreateCategory :execresult
INSERT INTO
    categories (name, slug)
VALUES
    (?, ?);

-- name: ListCategories :many
SELECT
    id,
    name,
    slug,
    created_at,
    updated_at
FROM
    categories
WHERE
    deleted_at IS NULL
ORDER BY
    name ASC;

-- name: GetCategoryByID :one
SELECT
    id,
    name,
    slug,
    created_at,
    updated_at
FROM
    categories
WHERE
    id = ?
    AND deleted_at IS NULL
LIMIT
    1;

-- name: UpdateCategory :exec
UPDATE categories
SET
    name = ?,
    slug = ?
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: DeleteCategory :exec
-- Soft Delete: Cukup update timestamp deleted_at
UPDATE categories
SET
    deleted_at = NOW()
WHERE
    id = ?;

-- name: GetCategoryByName :one
SELECT
    id,
    name,
    slug,
    created_at,
    updated_at
FROM
    categories
WHERE
    name = ?
    AND deleted_at IS NULL
LIMIT
    1;