// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: notes.sql

package db

import (
	"context"
	"time"
)

const createNote = `-- name: CreateNote :one
INSERT INTO "Notes"(user_id,content)
VALUES ($1,$2)
RETURNING id, user_id, content, created_at, last_updated
`

type CreateNoteParams struct {
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}

func (q *Queries) CreateNote(ctx context.Context, arg CreateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, createNote, arg.UserID, arg.Content)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const deleteNote = `-- name: DeleteNote :one
DELETE FROM "Notes"
WHERE id=$1
RETURNING id, user_id, content, created_at, last_updated
`

func (q *Queries) DeleteNote(ctx context.Context, id int64) (Note, error) {
	row := q.db.QueryRowContext(ctx, deleteNote, id)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const getNote = `-- name: GetNote :one
SELECT id, user_id, content, created_at, last_updated FROM "Notes"
WHERE user_id=$1 AND id=$2
`

type GetNoteParams struct {
	UserID int64 `json:"user_id"`
	ID     int64 `json:"id"`
}

func (q *Queries) GetNote(ctx context.Context, arg GetNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, getNote, arg.UserID, arg.ID)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const getNotesForUser = `-- name: GetNotesForUser :many
SELECT id, user_id, content, created_at, last_updated FROM "Notes"
WHERE user_id=$1
`

func (q *Queries) GetNotesForUser(ctx context.Context, userID int64) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getNotesForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Note
	for rows.Next() {
		var i Note
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.CreatedAt,
			&i.LastUpdated,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateNote = `-- name: UpdateNote :one
UPDATE "Notes"
SET content=$2,last_updated=$3
WHERE id=$1
RETURNING id, user_id, content, created_at, last_updated
`

type UpdateNoteParams struct {
	ID          int64     `json:"id"`
	Content     string    `json:"content"`
	LastUpdated time.Time `json:"last_updated"`
}

func (q *Queries) UpdateNote(ctx context.Context, arg UpdateNoteParams) (Note, error) {
	row := q.db.QueryRowContext(ctx, updateNote, arg.ID, arg.Content, arg.LastUpdated)
	var i Note
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}
