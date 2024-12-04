-- name: Timeline :many
SELECT * FROM tweets
WHERE userId IN (
	SELECT followerId FROM follows
	WHERE follows.followingId = ?
) AND isDeleted = false AND isReply = false
UNION ALL
SELECT * FROM tweets 
WHERE 
    isDeleted = FALSE
	AND isReply = FALSE
    AND created_at >= NOW() - INTERVAL 5 DAY
    AND userId NOT IN (
        SELECT followerId FROM follows WHERE follows.followingId = ?
    )
ORDER BY
	created_at DESC,
    likes + 2 * retweets DESC
LIMIT 100;