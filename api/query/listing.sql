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
INSERT INTO listing (userId, tweetId, listingname, listingdescription, listingprice, type, stock)
VALUES (?, ?, ?, ?, ?, ?, ?);
