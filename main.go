package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/elgs/gosqljson"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const theCase string = "lower"
const TIME_TO_CACHE float64 = 60

func main() {
	initDB()
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
	resources, err := runquery(db, TIME_TO_CACHE, query)
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
	// resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	resources, err := runquery(db, TIME_TO_CACHE, query, node, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, resources)
}
func createResource(c *gin.Context) {
	node := c.Param("node")

	// mengambil profile table
	query := "SHOW COLUMNS FROM " + node

	// Field             | Type             | Null | Key | Default | Extra
	resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	var field []string
	var values []string
	for _, r := range resources {
		// r["field"]
		// r["type"]
		// r["default"]
		// r["extra"]
		// r["key"]
		// r["null"]

		if r["null"] == "NO" && r["default"] == "" && !strings.Contains(r["extra"], "auto_increment") {
			// wajib user menginput datanya!
			val := c.PostForm(r["field"])
			if val == "" {
				c.JSON(http.StatusNotFound, "Some of field is required")
				return
			}
			field = append(field, r["field"])
			values = append(values, val)
			// simpan val dan field
		} else {
			// tidak wajib, jika tidak kosong aja yg dinput
			val := c.PostForm(r["field"])
			if val != "" {
				field = append(field, r["field"])
				values = append(values, val)
			}
		}
	}

	// membuat query insert di sini
	strField := strings.Join(field[:], ",")
	strValues := strings.Join(values[:], ",")
	query = "insert into " + node + "(" + strField + ") values (" + strValues + ")"
	res, err := db.Exec(query)

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
func updateResource(c *gin.Context) {
	node := c.Param("node")
	id := c.Param("id")
	// mengambil profile table
	query := "SHOW COLUMNS FROM " + node

	// Field             | Type             | Null | Key | Default | Extra
	resources, err := gosqljson.QueryDbToMap(db, theCase, query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	var sets []string
	for _, r := range resources {
		// r["field"]
		// r["type"]
		// r["default"]
		// r["extra"]
		// r["key"]
		// r["null"]
		val := c.PostForm(r["field"])
		if val != "" {
			sets = append(sets, r["field"]+"="+val)
		}
	}

	// membuat query insert di sini
	// query = "insert into " + node + "(" + strField + ") values (" + strValues + ")"
	strSets := strings.Join(sets, ",")
	query = "update " + node + " set " + strSets + " where id = " + id
	res, err := db.Exec(query)

	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
func deleteResource(c *gin.Context) {
	node := c.Param("node")
	id := c.Param("id")

	query := "delete from " + node + " where id =" + id
	_, err := db.Exec(query)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, "Record have deleted")
	fmt.Println(node, id)
}
