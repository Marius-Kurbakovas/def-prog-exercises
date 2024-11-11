package safesql

import (
	"context"
	"database/sql"
)

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	return db.db.QueryContext(ctx, query.s, args...)
}

type Rows = sql.Rows

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) ExecContext(ctx context.Context, query TrustedSQL, args ...any) (Result, error) {
	return db.db.ExecContext(ctx, query.s, args...)
}

type Result = sql.Result
