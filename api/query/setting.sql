-- name: CreateAccount :exec
INSERT INTO users (firebaseUid ,userId, username) VALUES (?, ?, ?);

-- name: CreateIsPrivate :exec
UPDATE users
SET isPrivate = ?
WHERE userId = ?;

-- name: CreateIsAdmin :exec
UPDATE users
SET isAdmin = ?
WHERE userId = ?;

-- name: CreateIsPremium :exec
UPDATE users
SET isPremium = true
WHERE userId = ?;

-- name: UpdateUsername :exec
UPDATE users
SET username = ?
WHERE userId = ?;

-- name: IsUserExists :one
SELECT EXISTS (
    SELECT 1 
    FROM users 
    WHERE userId = ?
);

-- name: GetIdByUID :one
SELECT userId FROM users WHERE firebaseUid = ?;