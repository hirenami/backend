-- name: CreateAccount :exec
INSERT INTO users (firebaseUid ,userId, username) VALUES (?, ?, ?);

-- name: CreateIsPrivate :exec
UPDATE users
SET isPrivate = ?
WHERE userId = ?;

-- name: CreateIsDeleted :exec
UPDATE users
SET isDeleted = ?
WHERE userId = ?;

-- name: CreateIsAdmin :exec
UPDATE users
SET isAdmin = ?
WHERE userId = ?;

-- name: GetIsPrivate :one
SELECT isPrivate FROM users WHERE userId = ?;

-- name: GetIsDeleted :one
SELECT isDeleted FROM users WHERE userId = ?;

-- name: GetIsAdmin :one
SELECT isAdmin FROM users WHERE userId = ?;

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