package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDBCOnn = 10
const maxIdleDBConn = 5
const maxLifeDBTime = 5 * time.Minute

func ConnectSQL(database string, dsn string) (*DB, error) {
	d, err := NewDatabase(database, dsn)
	if err != nil {
		panic(err)
	}
	d.SetMaxOpenConns(maxOpenDBCOnn)
	d.SetMaxIdleConns(maxIdleDBConn)
	d.SetConnMaxLifetime(maxLifeDBTime)
	dbConn.SQL = d

	if err := testDB(dbConn.SQL); err != nil {
		return nil, err
	}
	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	return err
}

func NewDatabase(database string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(database, dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
