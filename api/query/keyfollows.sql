-- name: CreateKeyFollow :exec
INSERT keyfollows (
	followerId, followingId
) VALUES (
	?, ?
);

-- name: DeleteKeyFollow :exec
DELETE FROM keyfollows
WHERE followerId = ? AND followingId = ?;

-- name: GetFollowRequest :many
SELECT followingId FROM keyfollows WHERE followerId = ?;

-- name: IsFollowRequest :one
SELECT EXISTS (
	SELECT 1
	FROM keyfollows
	WHERE followerId = ? AND followingId = ?
);

-- name: DeleteKeyFollows :exec
DELETE FROM keyfollows
WHERE followerId = ?;