package databaseopr

import (
	"database/sql"
	"fmt"
	"log"

    "github.com/tren03/artechles-ReactTS/backend/global"
	_ "github.com/lib/pq"
)


func AddPost(db *sql.DB, p global.Post) {
	sqlStatement := `
	INSERT INTO posts (title, url, body, date, category)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, p.Title, p.Url, p.Body, p.Date,p.Category)
	if err != nil {
		log.Fatalln("insert error", err)
	} else {
		fmt.Println("Record inserted successfully")
	}

}
 
