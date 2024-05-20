package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DBConnect()
}

func DBConnect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error verifying connection arguments:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging DB:", err)
	} else {
		fmt.Println("DB Connection Successful!")
	}
}

func ReturnDB() *sql.DB {
	if db == nil {
		log.Fatal("Database connection is not initialized")
	}
	return db
}
