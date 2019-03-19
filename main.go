package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type weather struct {
	Place         string
	LowTemp       int
	HighTemp      int
	Precipitation int
	Date          time.Time
}

type weatherReports struct {
	WeatherReports []weather
}

func main() {
	connStr := "dbname=pqgotest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	fmt.Println("Successfully connected to pqgotest")

	rows, err := db.Query(`
		SELECT *
		FROM weather`)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	weatherReports := weatherReports{}
	for rows.Next() {
		report := weather{}
		err = rows.Scan(
			&report.Place,
			&report.LowTemp,
			&report.HighTemp,
			&report.Precipitation,
			&report.Date,
		)
		weatherReports.WeatherReports = append(weatherReports.WeatherReports, report)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, report := range weatherReports.WeatherReports {
		fmt.Printf("The weather data for %s was read from the db\n", report.Place)
	}
}
