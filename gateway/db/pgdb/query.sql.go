// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package pgdb

import (
	"context"
)

const countNonUnique = `-- name: CountNonUnique :one
SELECT COUNT(*) FROM token WHERE occur > 1
`

func (q *Queries) CountNonUnique(ctx context.Context) (int64, error) {
	row := q.queryRow(ctx, q.countNonUniqueStmt, countNonUnique)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const readNonUnique = `-- name: ReadNonUnique :many


SELECT name, occur FROM token WHERE occur > 1 ORDER BY occur DESC LIMIT $1 OFFSET $2
`

type ReadNonUniqueParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

type ReadNonUniqueRow struct {
	Name  string `db:"name" json:"name"`
	Occur int32  `db:"occur" json:"occur"`
}

// -- name: BashInsert :exec
// INSERT INTO token(name, occur)(select * from unnest($1::int[], $2::int[]));
func (q *Queries) ReadNonUnique(ctx context.Context, arg ReadNonUniqueParams) ([]ReadNonUniqueRow, error) {
	rows, err := q.query(ctx, q.readNonUniqueStmt, readNonUnique, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadNonUniqueRow
	for rows.Next() {
		var i ReadNonUniqueRow
		if err := rows.Scan(&i.Name, &i.Occur); err != nil {
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

const writeToken = `-- name: WriteToken :exec
INSERT INTO token (name) VALUES ($1) 
ON CONFLICT (name) 
DO UPDATE SET occur = token.occur + 1
`

func (q *Queries) WriteToken(ctx context.Context, name string) error {
	_, err := q.exec(ctx, q.writeTokenStmt, writeToken, name)
	return err
}
