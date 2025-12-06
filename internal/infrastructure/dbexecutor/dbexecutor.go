package dbexecutor

import "database/sql"

// This encapsulates *sql.DB and *SQL.tx common properties I use to be able
// to work with Unit of Work pattern
type DbExecutor interface {
	Query(query string, args ...any) (*sql.Rows, error)
	Exec(query string, args ...any) (sql.Result, error)
}
