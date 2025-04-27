package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

func DeleteCase(case_id int) {
	// delete the case from the database
	query := fmt.Sprintf("DELETE FROM \"jCase\" where case_id=%d;", case_id)
	_, err := Db.Exec(query)
	if err != nil {
		log.Println("Error deleting case:", err)
	}

	fmt.Println("Case deleted successfully")
}

func DeleteCases(case_ids []int) {

	// delete the cases from the database
	ids := make([]string, len(case_ids))
	for i, id := range case_ids {
		ids[i] = strconv.Itoa(id)
	}
	query := fmt.Sprintf("DELETE FROM \"jCase\" where case_id IN (%s);", strings.Join(ids, ", "))
	_, err := Db.Exec(query)
	if err != nil {
		log.Println("Error deleting cases:", err)
	}
	fmt.Println("Cases deleted successfully")
}

func ClearTable(table string) {
	// connect to the database

	// delete all the cases from the database
	query := fmt.Sprintf("TRUNCATE TABLE \"%s\" CASCADE;", table)
	_, err := Db.Exec(query)
	if err != nil {
		log.Println("Error clearing table:", err)
	}
	// reset the sequence for the table
	// query = fmt.Sprintf("ALTER SEQUENCE \"%s_case_id_seq\" RESTART WITH 1;", table)
	// _, err = db.Exec(query)
	// if err != nil {
	// 	log.Println("Error resetting sequence:", err)
	// }
	fmt.Println("All rows deleted successfully from", table)
}
