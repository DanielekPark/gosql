package getreq

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Result struct {
	id      int64
	name    string
	link    string
	detials string
	Type    string
	tags    string
}

func GetLinks(c *gin.Context) {
	query := "SELECT * FROM links WHERE ID < 10"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetProducts) db.Query", err)
	}

	results := []Result{}
	for res.Next() {
		var list Result
		err := res.Scan(&list.id, &list.name)
		if err != nil {
			log.Fatal("There was a problem with ", err)
		}
		results = append(results, list)
	}

	c.JSON(http.StatusOK, results)
}
