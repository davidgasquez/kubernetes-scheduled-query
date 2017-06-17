package main

import (
	"os"
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
    "io/ioutil"
    "log"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    
    if len(os.Args) != 2 {
        fmt.Println("No query file given")
        return
    }
    
    var filePath string = os.Args[1]
    
	var user string = os.Getenv("REDSHIFT_USER")
	var password string = os.Getenv("REDSHIFT_PASSWORD")
    var endpoing string = os.Getenv("REDSHIFT_ENDPOINT")
    var database string = os.Getenv("REDSHIFT_DB_NAME")
    var port string = os.Getenv("REDSHIFT_DB_PORT")
    
    fileData, err := ioutil.ReadFile(filePath)
    query := string(fileData)    

    var uri string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, endpoing, port, database)

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