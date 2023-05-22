package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// establishing connection to the database
	fmt.Println("Connecting to db.... ")
	connStr := "user=postgres dbname=testing password=''sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	CheckError(err)
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nSuccessfully connected to the database!")
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
