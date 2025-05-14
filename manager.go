package main

import (
	"fmt"
	"mmoallapps/mmo-pg-manager/pkgs/database"
	"time"
)

func main() {
	// clear the screen
	startTime := time.Now()
	defer database.Close()
	// clear the screen
	fmt.Print("\033[H\033[2J")
	// connect to the database
	// print the title
	fmt.Println("MMO PG Manager")
	fmt.Println("MMO PG Manager is a tool to manage the MMO PostgreSQL database")

	database.SeedCases()
	fmt.Println("Seeding cases took :", time.Since(startTime))

	// run database.updateDBCases() every minute
	func() {
		for {
			updateStartTime := time.Now()
			fmt.Println("Updating cases...")
			database.UpdateDBCases()
			fmt.Println("Updating cases took :", time.Since(updateStartTime))
			time.Sleep(30 * time.Second)
		}
	}()

	fmt.Println("Time taken:", time.Since(startTime))
}
