-- 1. Total Event
-- name: CountEventsByOrganizer :one
SELECT COUNT(*) FROM events 
WHERE user_id = ? AND deleted_at IS NULL;

-- 2. Total Event Active
-- name: CountEventsActiveByOrganizer :one
SELECT COUNT(*) FROM events 
WHERE user_id = ? AND status = 'active' AND deleted_at IS NULL;

-- 3. Total Orders dari Event Organizer
-- name: CountOrdersByOrganizer :one
SELECT COUNT(*) FROM orders o
JOIN events e ON o.event_id = e.id
WHERE e.user_id = ? AND o.deleted_at IS NULL;

-- 4. Total Revenue Organizer
-- name: TotalRevenueByOrganizer :one
SELECT COALESCE(SUM(o.total_amount), 0) FROM orders o
JOIN events e ON o.event_id = e.id
WHERE e.user_id = ? AND o.status = 'paid' AND o.deleted_at IS NULL;

-- 5. Revenue Per Month (Organizer)
-- name: RevenuePerMonthByOrganizer :many
SELECT
    YEAR(o.created_at) AS year,
    MONTH(o.created_at) AS month,
    COALESCE(SUM(o.total_amount), 0) AS revenue
FROM orders o
JOIN events e ON o.event_id = e.id
WHERE e.user_id = ? 
    AND o.status = 'paid' 
    AND YEAR(o.created_at) = ?
    AND o.deleted_at IS NULL
GROUP BY YEAR(o.created_at), MONTH(o.created_at)
ORDER BY YEAR(o.created_at), MONTH(o.created_at);

-- 6. Revenue Per Event (Organizer)
-- name: RevenuePerEventByOrganizer :many
SELECT
    e.id AS event_id,
    e.title AS event_title,
    COUNT(o.id) AS total_orders,
    COALESCE(SUM(o.total_amount), 0) AS revenue
FROM events e
LEFT JOIN orders o ON o.event_id = e.id AND o.status = 'paid' AND o.deleted_at IS NULL
WHERE e.user_id = ? AND e.deleted_at IS NULL
GROUP BY e.id, e.title
ORDER BY revenue DESC;

-- 7. Recent Orders (Organizer)
-- name: RecentOrdersByOrganizer :many
SELECT
    o.id AS order_id,
    e.id AS event_id,
    e.title AS event_title,
    u.id AS user_id,
    u.name AS user_name,
    u.email AS user_email,
    o.total_amount,
    o.status,
    o.created_at
FROM orders o
JOIN events e ON o.event_id = e.id
JOIN users u ON o.user_id = u.id
WHERE e.user_id = ? AND o.deleted_at IS NULL
ORDER BY o.created_at DESC
LIMIT ?;

-- 8. Top Events (Organizer)
-- name: TopEventsByOrganizer :many
SELECT
    e.id AS event_id,
    e.title,
    COUNT(o.id) AS total_orders,
    COALESCE(SUM(o.total_amount), 0) AS revenue,
    COUNT(DISTINCT o.user_id) AS attendees
FROM events e
LEFT JOIN orders o ON o.event_id = e.id AND o.status = 'paid' AND o.deleted_at IS NULL
WHERE e.user_id = ? AND e.deleted_at IS NULL
GROUP BY e.id, e.title
ORDER BY revenue DESC
LIMIT ?;
