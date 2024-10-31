-- name: SearchByHashtag :many
SELECT * FROM tweets
WHERE tweetId IN (
	SELECT tweetId FROM hashtags
	WHERE hashtag = ?
) AND isDeleted = false ORDER BY createdAt DESC;

-- name: SearchByKeyword :many
SELECT * FROM tweets
WHERE content LIKE CONCAT('%', ? , '%') AND isDeleted = false ORDER BY createdAt DESC;

-- name: SearchUser :many
SELECT * FROM users
WHERE username LIKE CONCAT('%', ? , '%') ORDER BY createdAt DESC;
