package drivers

import (
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbOpen = 5
const maxIdleLifetime = 5 * time.Minute

// ConnectSQL creates a new database connection
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		return nil, err
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbOpen)
	d.SetConnMaxLifetime(maxIdleLifetime)

	dbConn.SQL = d
	if err = testDB(dbConn.SQL); err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB checks the connection
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}

// NewDatabase creates a new database connection
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}