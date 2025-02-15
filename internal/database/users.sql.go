// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createdUsser = `-- name: CreatedUsser :one
INSERT INTO users(id, created_at, updated_at, name)
VALUES(
    $1,
    $2,
    $3,
    $4
)
RETURNING id, created_at, updated_at, name
`

type CreatedUsserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func (q *Queries) CreatedUsser(ctx context.Context, arg CreatedUsserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createdUsser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
	)
	return i, err
}
