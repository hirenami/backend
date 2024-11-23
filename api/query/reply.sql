-- name: CreateReply :exec
INSERT INTO tweets (
	userId, isReply, content, media_url, review
) VALUES (
	?, true, ?, ?, ?
);

-- name: GetLastInsertID :one
SELECT LAST_INSERT_ID() AS tweetId;

-- name: GetUsersReplies :many
SELECT * FROM tweets 
WHERE userId = ? and isReply = true ORDER BY createdAt DESC;

-- name: RelateReplyToTweet :exec
INSERT INTO relations (
	tweetId, replyId
) VALUES (
	?, ?
);

-- name: UnrelateReplyToTweet :exec
DELETE FROM relations
WHERE tweetId = ? AND replyId = ?;

-- name: GetRepliesToTweet :many
SELECT replyId FROM relations
WHERE tweetId = ? ORDER BY created_at DESC;

-- name: GetTweetRepliedTo :one
SELECT tweetId FROM relations
WHERE replyId = ?;

-- name: PlusOneReply :exec
UPDATE tweets SET replies = replies + 1
WHERE tweetId = ?;

-- name: MinusOneReply :exec
UPDATE tweets SET replies = replies - 1
WHERE tweetId = ?;

-- name: CountReplies :one
SELECT replies FROM tweets
WHERE tweetId = ?;

-- name: IsReplyExists :one
SELECT EXISTS (
	SELECT 1 
	FROM relations
	WHERE replyId = ?
);