// This file is automatically generated.
package data

import (
  "errors"

  "github.com/jackc/pgx"
)

const VERSION = "0.0.2"

var ErrNotFound = errors.New("not found")

type Queryer interface {
  Query(sql string, args ...interface{}) (*pgx.Rows, error)
  QueryRow(sql string, args ...interface{}) *pgx.Row
  Exec(sql string, arguments ...interface{}) (pgx.CommandTag, error)
}
