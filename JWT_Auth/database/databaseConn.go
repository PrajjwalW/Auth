package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

//var db *sql.DB = databaseConn()

func DatabaseConn() *sql.DB {
	connString := "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error while connecting to db", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		log.Fatal("Error while pinging db", err)
	}

	return db

	//defer db.Close()

}

func CloseDbConn(db *sql.DB) {
	db.Close()
}
