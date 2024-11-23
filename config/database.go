package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB()  {
	var err error
	DB, err = sql.Open("mysql", "root:strong_password@tcp(127.0.0.1:3307)/inventory_db")
	if err != nil {
		log.Fatal("Failed to connect Database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to Ping Database: ", err)
	}
	log.Println("Database connected successfully!")
}