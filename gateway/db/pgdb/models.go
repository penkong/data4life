// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package pgdb

import (
	"database/sql"
)

type Sth struct {
	ID   int64          `db:"id" json:"id"`
	Name string         `db:"name" json:"name"`
	Bio  sql.NullString `db:"bio" json:"bio"`
}

type Todo struct {
	ID   int64          `db:"id" json:"id"`
	Name string         `db:"name" json:"name"`
	Bio  sql.NullString `db:"bio" json:"bio"`
}
