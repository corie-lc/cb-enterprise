package main

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func getDatabaseItem() *sql.DB {
	var db *sql.DB

	// Capture connection properties.
	cfg := mysql.Config{
		User:   "doadmin",
		Passwd: os.Getenv("password"),
		Net:    "tcp",
		Addr:   os.Getenv("server_name") + ":25060",
		DBName: "collec",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		fmt.Println("Failure")
		log.Fatal(pingErr)
	}

	return db

}
