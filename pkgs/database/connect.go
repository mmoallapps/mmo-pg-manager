package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// create the connection string to connect to the database
var connStr = "host=mmoallapps.jhacorp.com port=5432 user=postgres password=mmoallapps dbname=mmoallapps sslmode=disable"

// connect
var Db *sql.DB
var err error

func init() {
	// connect to the database
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// check if the connection is alive
	if err = Db.Ping(); err != nil {
		panic(err)
	}
}

// close the connection when the program exits
func Close() {
	if err := Db.Close(); err != nil {
		panic(err)
	}
}

// getter for the db
func DB() *sql.DB {
	return Db
}
