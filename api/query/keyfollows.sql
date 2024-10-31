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