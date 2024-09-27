package databaseopr

import (
	"database/sql"
	"fmt"
	"log"
    "github.com/tren03/artechles-ReactTS/backend/global"
)

func Dbconnect() (*sql.DB,error){

	fmt.Println("connecting to postgress")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		global.Host, global.Port, global.User, global.Password, global.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Not connected to database")
        return nil,err
	}
	fmt.Println("connected yes !")
    return db,nil
    
}
