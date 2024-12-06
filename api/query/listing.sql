-- name: GetListing :one
SELECT * from listing
WHERE listingId = ?;

-- name: GetUserListings :many
SELECT * from listing
WHERE userId = ?;

-- name: GetListingByTweet :one
SELECT * from listing
WHERE tweetId = ?;

-- name: CreateListing :exec
INSERT INTO listing (listingId, userId, tweetId, listingname, listingdescription, listingprice, type, stock, `condition`)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: DeleteStock :exec
UPDATE listing
SET stock = stock - 1
WHERE listingId = ?;

-- name: UpdateListing :exec
UPDATE users
SET listingnum = listingnum + 1
WHERE userId = ?;

-- name: GetRandomListings :many
SELECT listingId 
FROM listing
ORDER BY RAND()
LIMIT 5; -- ランダムに10件取得