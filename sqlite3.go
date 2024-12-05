package entsqlite

import (
	"database/sql"
	"database/sql/driver"
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
	return conn, nil
}

func init() {
	sql.Register("sqlite3", sqliteDriver{Driver: &sqlite.Driver{}})
}
