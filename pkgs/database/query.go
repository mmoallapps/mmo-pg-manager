package database

import (
	"strings"

	_ "github.com/lib/pq"
)

// make a connection to the database

// get all the tables
func GetTables() ([]string, error) {
	rows, err := Db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// create a slice to hold the table names
	var tables []string
	// iterate over the rows and append the table names to the slice
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}
		// print out the table name
		tables = append(tables, table)
	}
	// check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// return the slice of table names
	return tables, nil
}

// run a raw query against the database
func RunQuery(query string) (string, error) {
	// run the query
	rows, err := Db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	// create a slice to hold the results
	var results []string
	// iterate over the rows and append the results to the slice
	for rows.Next() {
		var result string
		if err := rows.Scan(&result); err != nil {
			return "", err
		}
		results = append(results, result)
	}
	// check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return "", err
	}
	// return the results as a string
	return strings.Join(results, "\n"), nil
}
