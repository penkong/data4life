// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package pgdb

import ()

type Token struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Occur int32  `db:"occur" json:"occur"`
}