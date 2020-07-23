package main

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:password@tcp(rdb:3306)/performance_schema")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("success")
	time.Sleep(1000000000 * time.Second)
}
