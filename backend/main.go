package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"github.com/tren03/artechles-ReactTS/backend/databaseopr"
	"github.com/tren03/artechles-ReactTS/backend/routes"
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
    db,err := databaseopr.Dbconnect()
    if err!=nil{
        log.Println("Database connection failed")
    }
    defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/getallposts", routes.HandleGetAllPosts(db))
	r.HandleFunc("/", routes.HandleRoot)

	handler := cors.Default().Handler(r)
	log.Println("starting server at port 3000")
	log.Fatal(http.ListenAndServe(":3000", handler))
}
