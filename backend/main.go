package main

import (
	"database/sql"
	"fmt"
	"time"
    "github.com/tren03/artechles-ReactTS/backend/global"
    
    "github.com/tren03/artechles-ReactTS/backend/databaseOpr"
    

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
		panic(err)
	}
	defer db.Close()

	fmt.Println("connected yes !")
	p := global.Post{
		ID:       54,
		Title:    "Go By example",
		Url:      "https://gobyexample.com/",
		Body:     "A site that is a good primer for all of the main concepts covered in go language with code snippets for each.",
		Date:     time.Now(),
		Category: "Go Programming",
	}
    databaseopr.AddPost(db,p)	

	sqlGet := `SELECT * from test;`
	row := db.QueryRow(sqlGet)
	var name string
	err = row.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(name)
}
