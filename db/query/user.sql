-- name: CreateUser :one
INSERT INTO "User"(uid)
VALUES ($1)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "User"
WHERE uid=$1;