package databaseopr

import (
	"database/sql"
	"log"

	"github.com/tren03/artechles-ReactTS/backend/global"
)

func GetAllPosts(db *sql.DB) ([]global.Post,error){
	rows, err := db.Query("SELECT * from posts")
	if err != nil {
		log.Println("error while getting posts from the database")
	}

	var posts []global.Post
	for rows.Next() {
		var p global.Post
		err := rows.Scan(&p.ID, &p.Title, &p.Url, &p.Body, &p.Date, &p.Category)
		if err != nil {
			log.Println("error in reading rows to variable")
            return posts,err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		log.Println("error in reading rows to variable => 2")
        return posts,err
	}
    return posts,nil
}
