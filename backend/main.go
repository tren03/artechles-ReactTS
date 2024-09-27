package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "articles"
)

// the format i should get from frontend ->

/*
	type postData = {
	    Title: string;
	    Url: string;
	    Body: string;
	    Date: string;
	    Category: string;
	};
*/

func main() {
	fmt.Println("testing postgress con")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Not connected to database")
	}
	defer db.Close()
	fmt.Println("connected yes !")

}
