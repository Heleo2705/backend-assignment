-- name: CreateUser :one
INSERT INTO "User"(name,password)
VALUES ($1,$2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM "User"
WHERE name=$1;