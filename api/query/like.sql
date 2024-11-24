-- name: PlusLike :exec
UPDATE tweets 
SET likes = likes + 1
WHERE tweetId = ?;

-- name: MinusLike :exec
UPDATE tweets
SET likes = likes - 1
WHERE tweetId = ?;

-- name: GetLikes :one
SELECT likes FROM tweets
WHERE tweetId = ?;

-- name: CreateLike :exec
INSERT INTO likes (
	userId, tweetId
) VALUES (
	?, ?
);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE userId = ? AND tweetId = ?;

-- name: CheckLike :many
SELECT userId FROM likes
WHERE tweetId = ?;

-- name: GetUsersLikes :many
SELECT l.tweetId
FROM likes l
JOIN tweets t ON l.tweetId = t.tweetId
WHERE l.userId = ? -- likesテーブルのuserIdを明示
  AND t.isDeleted = false
ORDER BY l.createdAt DESC;

-- name: IsLiked :one
SELECT EXISTS (
	SELECT 1
	FROM likes
	WHERE userId = ? AND tweetId = ?
);