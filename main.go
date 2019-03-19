package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "dbname=pqgotest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	fmt.Println("Successfully connected to pqgotest")

	var tempLow int
	err = db.QueryRow("SELECT temp_lo FROM weather WHERE place = 'Snøhetta'").Scan(&tempLow)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The temperature at Snøhetta is %d\n", tempLow)
}
