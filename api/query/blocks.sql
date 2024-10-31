-- name: CreateBlock :exec
INSERT INTO blocks (
	blockerId, blockedId
) VALUES (
	?, ?
);

-- name: DeleteBlock :exec
DELETE FROM blocks WHERE blockerId = ? AND blockedId = ?;

-- name: IsBlocked :one
SELECT EXISTS (
	SELECT 1 
	FROM blocks 
	WHERE blockerId = ? AND blockedId = ?
);

-- name: GetBlocks :many
SELECT blockedId FROM blocks WHERE blockerId = ?;