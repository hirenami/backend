-- name: CreateFollow :exec
INSERT follows (
	followerId, followingId
) VALUES (
	?, ?
);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE followerId = ? AND followingId = ?;

-- name: CountFollowing :one
SELECT COUNT(followerId) FROM follows WHERE followingId = ?;

-- name: CountFollower :one
SELECT COUNT(followingId) FROM follows WHERE followerId = ?;

-- name: GetFollowing :many
SELECT followerId FROM follows WHERE followingId = ?;

-- name: GetFollower :many
SELECT followingId FROM follows WHERE followerId = ?;

-- name: IsFollowing :one
SELECT EXISTS (
	SELECT 1
	FROM follows
	WHERE followerId = ? AND followingId = ?
);