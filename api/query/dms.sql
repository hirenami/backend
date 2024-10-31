-- name: CreateDm :exec
INSERT INTO dms (
	senderId, receiverId, content , media_url
) VALUES (
	?, ?, ?, ?
);

-- name: DeleteDm :exec
DELETE FROM dms WHERE dmsId = ?;

-- name: GetDms :many
SELECT * FROM dms 
WHERE (senderId = ? AND receiverId = ?) OR (senderId = ? AND receiverId = ?) ORDER BY createdAt DESC;

-- name: GetLastMessages :many
SELECT * FROM dms 
WHERE (senderId = ? AND receiverId = ?) OR (senderId = ? AND receiverId = ?) ORDER BY createdAt DESC
LIMIT 1;

-- name: GetDm :one
SELECT * FROM dms WHERE dmsId = ?;

-- name: GetDmsUsers :many
SELECT DISTINCT 
    CASE 
        WHEN senderId = ? THEN receiverId 
        ELSE senderId 
    END AS otherUserId
FROM dms
WHERE senderId = ? OR receiverId = ?;

-- name: SetDmStatus :exec
UPDATE dms
SET status = 'read'
WHERE (senderId = ? AND receiverId = ?) OR (senderId = ? AND receiverId = ?);