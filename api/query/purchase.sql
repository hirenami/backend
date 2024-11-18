-- name: GetPurchase :one
SELECT * from purchase
WHERE purchaseId = ?;

-- name: GetUserPurchases :many
SELECT * from purchase
WHERE userId = ?;