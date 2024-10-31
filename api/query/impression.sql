-- name: PlusImpression :exec
UPDATE tweets
SET impressions = impressions + 1
WHERE tweetId = ?;

-- name: GetImpression :one
SELECT impressions
FROM tweets
WHERE tweetId = ?;