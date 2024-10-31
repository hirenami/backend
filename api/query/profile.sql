-- name: CreateBiography :exec
UPDATE users
SET biography = ?
WHERE userId = ?;

-- name: CreateHeaderImage :exec
UPDATE users
SET header_image = ?
WHERE userId = ?;

-- name: CreateIconImage :exec
UPDATE users
SET icon_image = ?
WHERE userId = ?;

-- name: GetProfile :one
SELECT * FROM users WHERE userId = ?;