-- name: SumRevenue :one
SELECT
    CAST(COALESCE(SUM(amount), 0)AS SIGNED) AS total
FROM
    payments
WHERE
    status = 'paid'
    AND deleted_at IS NULL;

-- name: RevenuePerWeek :many
SELECT
    YEAR(created_at) AS year,
    WEEK(created_at, 1) AS week,
    CAST(COALESCE(SUM(amount), 0)AS SIGNED) AS revenue
FROM
    payments
WHERE
    status = 'paid'
    AND YEAR(created_at) = ?
    AND deleted_at IS NULL
GROUP BY
    YEAR(created_at),
    WEEK(created_at, 1)
ORDER BY
    YEAR(created_at),
    WEEK(created_at, 1);

-- name: RevenuePerMonth :many
SELECT
    YEAR(created_at) AS year,
    MONTH(created_at) AS month,
    CAST(COALESCE(SUM(amount), 0)AS SIGNED) AS revenue
FROM
    payments
WHERE
    status = 'paid'
    AND YEAR(created_at) = ?
    AND deleted_at IS NULL
GROUP BY
    YEAR(created_at),
    MONTH(created_at)
ORDER BY
    YEAR(created_at),
    MONTH(created_at);

-- name: RevenuePerYear :many
SELECT
    YEAR(created_at) AS year,
    CAST(COALESCE(SUM(amount), 0)AS SIGNED) AS revenue
FROM
    payments
WHERE
    status = 'paid'
    AND deleted_at IS NULL
GROUP BY
    YEAR(created_at)
ORDER BY
    YEAR(created_at);

-- name: RevenuePerDateRange :many
SELECT
    DATE(created_at) AS date,
    CAST(COALESCE(SUM(amount), 0)AS SIGNED) AS revenue
FROM
    payments
WHERE
    status = 'paid'
    AND deleted_at IS NULL
    AND created_at BETWEEN ? AND ?
GROUP BY
    DATE(created_at)
ORDER BY
    DATE(created_at);