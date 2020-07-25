package main

import (
	"database/sql"

	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{Addr: "redis:6379"})
	if err := client.Ping().Err(); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/performance_schema")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal()
	}

	// アプリケーションでの稼働待ち
	// for {
	// 	err = db.Ping()
	// 	if err == nil {
	// 		break
	// 	}
	// 	time.Sleep(3 * time.Second)
	// }

	log.Println("success")
}
