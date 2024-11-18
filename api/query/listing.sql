-- name: GetListing :many
SELECT * from listing
WHERE listingId = ?;

-- name: GetUserListings :many
SELECT * from listing
WHERE userId = ?;

-- name: GetListingByTweet :one
SELECT * from listing
WHERE tweetId = ?;