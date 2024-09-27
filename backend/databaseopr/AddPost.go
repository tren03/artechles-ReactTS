package databaseopr
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/tren03/artechles-ReactTS/backend/global"
)

func AddPost(db *sql.DB, p global.Post) {
	sqlStatement := `
	INSERT INTO posts (title, url, body, date, category)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, p.Title, p.Url, p.Body, p.Date, p.Category)
	if err != nil {
		log.Println("insert error", err)
	} else {
		fmt.Println("Record inserted successfully")
	}

}

func AddMultiplePosts(db *sql.DB, p []global.Post) {
	sqlStatement := `
	INSERT INTO posts (title, url, body, date, category)
	VALUES ($1, $2, $3, $4, $5)`

	for index, post := range p {
		_, err := db.Exec(sqlStatement, post.Title, post.Url, post.Body, post.Date, post.Category)
		if err != nil {
			log.Println("insert error multiple", err)
		} else {
			log.Println("inserted record ",index )
		}

	}

}
