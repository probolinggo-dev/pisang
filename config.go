package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

// const host, user, password, database string = "127.0.0.1", "mbp", "mbp#123", "test"

// Config :
type Config struct {
	Database Database `json:"database"`
}

// Database :
type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func loadSettings() (db Database, err error) {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "/config.json")
	jsonFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// close file
	defer jsonFile.Close()

	// convert jsonFile ke byteArray
	byteVal, _ := ioutil.ReadAll(jsonFile)

	// convert byteArray ke variable Config
	var config Config
	err = json.Unmarshal(byteVal, &config)
	db = config.Database
	return
}
