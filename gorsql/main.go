package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("No query file given")
		return
	}

	user := os.Getenv("REDSHIFT_USER")
	password := os.Getenv("REDSHIFT_PASSWORD")
	endpoing := os.Getenv("REDSHIFT_ENDPOINT")
	database := os.Getenv("REDSHIFT_DB_NAME")
	port := os.Getenv("REDSHIFT_DB_PORT")

	fileData, err := ioutil.ReadFile(os.Args[1])
	query := string(fileData)

	if err != nil {
		log.Fatal(err)
	}

	uri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, endpoing, port, database)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	result, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
