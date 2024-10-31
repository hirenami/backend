-- name: CreateHashtag :exec
INSERT INTO hashtags (
	hashtag, tweetId
) VALUES (
	?, ?
);

-- name: UpdateHashtag :exec
UPDATE hashtags
SET hashtag = ?
WHERE tweetId = ?;

-- name: DeleteHashtag :exec
DELETE FROM hashtags
WHERE tweetId = ?;