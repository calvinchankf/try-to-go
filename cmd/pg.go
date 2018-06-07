package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "calvin"
	DB_PASSWORD = "handsome"
	DB_NAME     = "bauhaus"
)

func main() {
	fmt.Println("try PQ")

	// connStr := "postgres://calvin:handsome@localhost:5432/bauhaus?sslmode=disable"
	// db, err := sql.Open("postgres", connStr)

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	/*
		CREATE TABLE IF NOT EXISTS "user" (
			"id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			"username" TEXT,
			"password" TEXT,
			"created_timestamp" TIMESTAMP WITH TIME ZONE DEFAULT now(),
			"last_login_timestamp" TIMESTAMP WITH TIME ZONE DEFAULT now()
		);
	*/

	fmt.Println("select fields")
	rows, err := db.Query(`SELECT id, username, password FROM public.user`)
	checkErr(err)

	for rows.Next() {
		var id string
		var username string
		var password string
		err = rows.Scan(&id, &username, &password)
		checkErr(err)
		fmt.Printf("%v %v %v \n", id, username, password)
	}

	fmt.Println("select all")
	rows, err = db.Query(`SELECT * FROM public.user`)
	checkErr(err)

	var id string
	for rows.Next() {
		// var id string
		var username string
		var password string
		var createdTimestamp string                                                   // can be string
		var lastTimestamp time.Time                                                   // or can be Date
		err = rows.Scan(&id, &username, &password, &createdTimestamp, &lastTimestamp) // all fields must be filled in order to be Scan() or an error
		checkErr(err)
		fmt.Printf("%v %v %v %v %v \n", id, username, password, createdTimestamp, lastTimestamp)
	}

	fmt.Printf("update row %s \n", id)
	// stmt, err := db.Prepare("update public.user set username=$1 where id=$2 returning *")
	// res, err := stmt.Exec("cal cal cal", id)
	rows, err = db.Query("update public.user set username=$1 where id=$2 returning *", "cal", id)
	checkErr(err)
	for rows.Next() {
		// var id string
		var username string
		var password string
		var createdTimestamp string                                                   // can be string
		var lastTimestamp time.Time                                                   // or can be Date
		err = rows.Scan(&id, &username, &password, &createdTimestamp, &lastTimestamp) // all fields must be filled in order to be Scan() or an error
		checkErr(err)
		fmt.Printf("%v %v %v %v %v \n", id, username, password, createdTimestamp, lastTimestamp)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
