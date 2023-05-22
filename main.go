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

	// inserting row into database
	ins_sql_statement := `INSERT INTO students VALUES ($1, $2, $3,$4)`
	_, err = db.Exec(ins_sql_statement, 23, "Rabbit", "B.Tech", "English")
	CheckError(err)
	fmt.Println("\nRow inserted successfully!")

	// updating row in database
	upd_sql_statement := `UPDATE students SET std_lang=$1 WHERE std_id=$2;`
	res, err := db.Exec(upd_sql_statement, "Hindi", 15)
	CheckError(err)
	count, err := res.RowsAffected()

	CheckError(err)
	fmt.Printf("rows updated: %v\n", count)

	// delete row/rows in a table
	del_sql_statement := `DELETE FROM students WHERE std_id=$1;`
	res, err = db.Exec(del_sql_statement, 17)
	CheckError(err)
	count, err = res.RowsAffected()

	CheckError(err)
	fmt.Printf("rows deleted: %v\n", count)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
