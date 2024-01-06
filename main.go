package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func main () {

	// initialize router
	r := gin.Default()
	// connect to db
	db, err := sql.Open("mysql", "todoadmin:test1234@tcp(127.0.0.1:3306)/gotodo")
	if err != nil {
		log.Fatal(err)
	}	
	defer db.Close()
	// start 

	r.Run()


	
}