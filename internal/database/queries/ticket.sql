-- name: BookTicket :execresult
-- Dipakai saat warga klik "Booking" (Status default: 'booked')
INSERT INTO
    tickets (event_id, user_id, ticket_code, status)
VALUES
    (?, ?, ?, ?);

-- name: GetTicketByCode :one
-- Dipakai pas panitia nge-scan QR Code tiket di gerbang masuk
SELECT
    t.id,
    t.event_id,
    t.user_id,
    t.ticket_code,
    t.status,
    t.created_at,
    e.title AS event_title,
    e.start_time AS event_start_time,
    u.name AS attendee_name
FROM
    tickets t
    JOIN events e ON t.event_id = e.id
    JOIN users u ON t.user_id = u.id
WHERE
    t.ticket_code = ?
    AND t.deleted_at IS NULL
LIMIT
    1;

-- name: ListMyTickets :many
-- Dipakai di aplikasi mobile/web warga buat lihat daftar tiket yang mereka beli
SELECT
    t.id,
    t.ticket_code,
    t.status,
    t.created_at,
    e.title AS event_title,
    e.location AS event_location,
    e.start_time AS event_start_time
FROM
    tickets t
    JOIN events e ON t.event_id = e.id
WHERE
    t.user_id = ?
    AND t.deleted_at IS NULL
ORDER BY
    t.created_at DESC;

-- name: UpdateTicketStatus :exec
-- Dipakai pas pembayaran lunas ('paid') atau pas tiket di-scan masuk ('used')
UPDATE tickets
SET
    status = ?,
    updated_at = NOW()
WHERE
    id = ?
    AND deleted_at IS NULL;

-- name: SoftDeleteTicket :exec
-- Membatalkan tiket tanpa menghapus history finansial
UPDATE tickets
SET
    deleted_at = NOW(),
    status = 'cancelled'
WHERE
    id = ?;

-- name: HardDeleteTicket :exec
-- Hapus fisik dari disk (Khusus clean up data testing)
DELETE FROM tickets
WHERE
    id = ?;