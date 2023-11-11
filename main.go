package main

import (
	"database/sql"
	getreq "db/api/get"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	err := godotenv.Load()

	//Creates a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	router := gin.Default()
	router.GET("/links", getreq.GetLinks)

	//Runs the router
	router.Run()
}
