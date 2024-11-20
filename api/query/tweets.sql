-- name: CreateTweet :exec
INSERT INTO tweets (
	userId, content, media_url
) VALUES (
	?, ?, ?
);

-- name: GetUsersTweets :many
SELECT * FROM tweets 
WHERE userId = ? and isReply = false and isDeleted = false ORDER BY created_at DESC;

-- name: DeleteTweet :exec
UPDATE tweets
SET isDeleted = true
WHERE tweetId = ?;

-- name: EditTweet :exec
UPDATE tweets
SET content = ?, media_url = ?
WHERE tweetId = ?;

-- name: CreateRetweet :exec
INSERT INTO tweets (
	userId, retweetId, content
) VALUES (
	?, ? , ''
);

-- name: CreateQuote :exec
INSERT INTO tweets (
	userId, retweetId, content, media_url, isQuote
) VALUES (
	?, ?, ?, ?, true
);

-- name: GetRetweets :many
SELECT userId FROM tweets
WHERE retweetId = ? and isQuote = false and isDeleted = false ORDER BY created_at DESC;

-- name: GetQuotes :many
SELECT * FROM tweets
WHERE retweetId = ? and isQuote = true and isDeleted = false ORDER BY created_at DESC;

-- name: PlusRetweet :exec
UPDATE tweets
SET retweets = retweets + 1
WHERE tweetId = ?;

-- name: MinusRetweet :exec
UPDATE tweets
SET retweets = retweets - 1
WHERE tweetId = ?;

-- name: GetRetweetsCount :one
SELECT retweets FROM tweets
WHERE tweetId = ? and isDeleted=false;

-- name: IsTweetExists :one
SELECT EXISTS (
    SELECT 1 
    FROM tweets 
	WHERE tweetId = ?
);

-- name: GetUserId :one
SELECT userId FROM tweets
WHERE tweetId = ?;

-- name: GetRetweetId :one
SELECT retweetId FROM tweets
WHERE tweetId = ?;

-- name: GetTweet :one
SELECT * FROM tweets
WHERE tweetId = ?;                

-- name: IsRetweet :one
SELECT EXISTS (
	SELECT 1 
	FROM tweets 
	WHERE retweetId = ? and isDeleted = false and userId = ? and isQuote = false
);

-- name: GetTweetId :one
SELECT tweetId FROM tweets
WHERE retweetId = ? and userId = ? and isDeleted = false and isQuote = false;