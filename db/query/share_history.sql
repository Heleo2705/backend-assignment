-- name: ShareNote :one
INSERT INTO "ShareHistory"(owner_id,shared_id)
VALUES ($1,$2)
RETURNING *;

-- name: GetSharedNotes :many
SELECT * FROM "ShareHistory"
WHERE owner_id=$1;