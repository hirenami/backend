-- name: CreateNotification :exec
INSERT INTO notifications (
	senderId, replyId, type, contentId
) VALUES (
	?, ?, ?, ?
);

-- name: GetNotifications :many
SELECT * FROM notifications WHERE replyId = ? ORDER BY createdAt DESC;

-- name: GetNotification :one
SELECT * FROM notifications WHERE notificationsId = ?;

-- name: UpdateNotificationStatus :exec
UPDATE notifications
SET status = 'read'
WHERE notificationsId = ?;

-- name: IsNotificationExist :one
SELECT EXISTS (
    SELECT 1 
    FROM notifications 
	WHERE notificationsId = ?
);