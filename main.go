package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Result struct {
	id      int64
	name    string
	link    string
	detials string
	Types   string
	tags    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env line 28 ", err)
	}

	// Database connection
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// Routes
	router := gin.Default()
	router.GET("/search", GetLinks)

	// Run the router
	router.Run()
}

func GetLinks(c *gin.Context) {
	query := "SELECT name FROM links WHERE id < 20"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetLinks) db.Query", err)
	}

	results := []Result{}
	for res.Next() {
		var link Result
		err := res.Scan(&link.Types)
		if err != nil {
			log.Fatal("(GetLinks) res.Scan", err)
		}
		results = append(results, link)
	}

	c.JSON(http.StatusOK, results)
}
