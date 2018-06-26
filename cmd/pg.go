package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

const (
	DB_USER     = "calvin"
	DB_PASSWORD = "handsome"
	DB_NAME     = "bauhaus"
)

var pgCmd = &cobra.Command{
	Use:   "pg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pg called")
		// pg()
		upsert()
	},
}

func init() {
	rootCmd.AddCommand(pgCmd)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func pg() {
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

func upsert() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var hashtable = make(map[string]string)
	hashtable["a"] = "calvi"
	hashtable["b"] = "calvin"
	hashtable["c"] = "calvin_c"

	// marshal
	jsonObj, _ := json.Marshal(hashtable)

	var id string
	var actionFlag string

	// insert or update every time
	// err = db.QueryRow(`
	// 	INSERT INTO public.orders (code, raw_data) VALUES ($1, $2)
	// 	ON CONFLICT (code)
	// 	DO UPDATE SET raw_data = $2
	// 	RETURNING id
	// 	`,
	// 	"calvin_code",
	// 	jsonObj,
	// ).Scan(&userid)

	// inserted, updated, none
	err = db.QueryRow(`
		WITH inserted AS (
			INSERT INTO public.orders (code, what, raw_data) VALUES ($1, $2, $3)
			ON CONFLICT (code) DO NOTHING
			RETURNING id
		),
		updated AS (
			UPDATE public.orders t
			SET
				what = $2,
				raw_data = $3
			WHERE code = $1
			AND (
				MD5(CAST((
					$2,
					$3::jsonb
				) AS TEXT))
			) IS DISTINCT FROM (
				SELECT MD5(CAST((
					what,
					raw_data
				) AS TEXT))
				FROM public.orders
				WHERE code = $1
			)
			RETURNING t.id
		),
		none AS (
			SELECT id
			FROM public.orders
			WHERE code = $1
			AND what = $2
			AND raw_data = $3
		)
		SELECT id, 'inserted' AS action_flag FROM inserted
		UNION
		SELECT id, 'updated' AS action_flag FROM updated
		UNION
		SELECT id, 'none' AS action_flag FROM none
		`,
		"calvin_code",
		"whateve",
		jsonObj,
	).Scan(&id, &actionFlag)

	checkErr(err)
	fmt.Println(id, actionFlag)
}
