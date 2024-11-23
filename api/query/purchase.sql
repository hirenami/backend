-- name: GetPurchase :one
SELECT * from purchase
WHERE purchaseId = ?;

-- name: GetUserPurchases :many
SELECT * from purchase
WHERE userId = ?;

-- name: CreatePurchase :exec
INSERT INTO purchase (userId, listingId)
VALUES (?, ?);