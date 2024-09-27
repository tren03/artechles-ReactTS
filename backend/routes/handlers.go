package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tren03/artechles-ReactTS/backend/databaseopr"
	"github.com/tren03/artechles-ReactTS/backend/global"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HELLO WEB")
}

func HandleGetAllPosts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var posts []global.Post
		posts, err := databaseopr.GetAllPosts(db)
		if err != nil {
			log.Println("error in getting posts from db in handler")
			http.Error(w, "error in getting posts from db in handler", http.StatusInternalServerError)
            return

		}
		json_posts, err := json.Marshal(posts)
		if err != nil {
			log.Println("error while marshalling json")
			http.Error(w, "error while marshalling json", http.StatusInternalServerError)
            return 
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(json_posts)
		if err != nil {
			log.Println("err while writing data to response")
			http.Error(w, "err while writing data to response", http.StatusInternalServerError)
            return
		}

	}
}
