-- name: Timeline :many
SELECT * FROM tweets
WHERE userId IN (
	SELECT followerId FROM follows
	WHERE followingId = ?
) AND isDeleted = false AND isReply = false ORDER BY created_at DESC;