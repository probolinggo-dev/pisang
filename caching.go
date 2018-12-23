package main

import (
	"database/sql"
	"time"

	"github.com/elgs/gosqljson"
)

type kuki struct {
	interval  float64
	updatedAt time.Time
	value     []map[string]string
}

var kukis map[string]kuki

func runquery(con *sql.DB, interval float64, query string, params ...interface{}) ([]map[string]string, error) {
	if val, ok := kukis[query]; ok {
		elapsed := time.Since(val.updatedAt)
		if elapsed.Seconds() < val.interval {
			return val.value, nil
		}
	}
	result, err := gosqljson.QueryDbToMap(db, theCase, query, params)
	if err != nil {
		return nil, err
	}
	go simpanresult(query, result, interval)
	return result, nil
}

func simpanresult(key string, value []map[string]string, interval float64) {
	if kukis == nil {
		kukis = make(map[string]kuki)
	}
	kukibaru := kuki{10, time.Now(), value}
	kukis[key] = kukibaru
}
