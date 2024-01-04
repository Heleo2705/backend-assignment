-- name: CreateNote :one
INSERT INTO "Notes"(user_id,content)
VALUES ($1,$2)
RETURNING *;

-- name: UpdateNote :one
UPDATE "Notes"
SET content=$2,last_updated=$3
WHERE id=$1
RETURNING *;

-- name: GetNotesForUser :many
SELECT * FROM "Notes"
WHERE user_id=$1;