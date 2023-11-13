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
	Id      int64
	Name    string
	Link    string
	Details string
	Types   string
	Tags    string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
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
	query := "SELECT * FROM links WHERE id < 20"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetLinks) db.Query", err)
	}

	results := []Result{}
	for res.Next() {
		var link Result
		err := res.Scan(&link.Id, &link.Name, &link.Link, &link.Details, &link.Types, &link.Tags)
		if err != nil {
			log.Fatal("(GetLinks) res.Scan", err)
		}
		results = append(results, link)
	}

	c.JSON(http.StatusOK, results)
}
