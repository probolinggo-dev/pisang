package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/elgs/gosqljson"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

const theCase string = "lower"

func main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/:node", getResources)
		api.GET("/:node/:id", getResourceDetils)
		api.POST("/:node", createResource)
		api.PUT("/:node/:id", updateResource)
		api.DELETE("/:node/:id", deleteResource)
	}

	router.Run()
}

func initDB() {
	conf, err := loadSettings()
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("mysql", conf.User+":"+conf.Password+"@tcp("+conf.Host+":3306)/"+conf.DBName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Koneksi ke database sukses")
}
func getResources(c *gin.Context) {
	node := c.Param("node")
	query := fmt.Sprintf("select * from %s", node)
	resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, resources)
	fmt.Println(node)
}

func getResourceDetils(c *gin.Context) {
	node := c.Param("node")
	id := c.Param("id")
	query := fmt.Sprintf("select * from %s where id = '%s'", node, id)
	resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, resources)
}
func createResource(c *gin.Context) {
	node := c.Param("node")
	query := "SHOW COLUMNS FROM " + node
	resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, resources)
	fmt.Println(node)
}
func updateResource(c *gin.Context) {
	node := c.Param("node")
	id := c.Param("id")

	fmt.Println(node, id)
}
func deleteResource(c *gin.Context) {
	node := c.Param("node")
	id := c.Param("id")

	fmt.Println(node, id)
}
