package entsqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"modernc.org/sqlite"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	if c, ok := conn.(driver.ExecerContext); ok {
		if _, err := c.ExecContext(context.Background(), "PRAGMA foreign_keys = on;", nil); err != nil {
			conn.Close()
			return nil, fmt.Errorf("failed to enable enable foreign keys: %w", err)
		}
	}

	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}
