package driver

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"time"
)

//DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

//dbConn reference to DB
var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectSQL Create connection pool for PG
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	//Check for error and die
	if err != nil {
		panic(err)
	}

	//Set Conn param
	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewDatabase creates a new database for the application
func NewDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	//Check for error
	if err != nil {
		return nil, err
	}

	//Test connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}