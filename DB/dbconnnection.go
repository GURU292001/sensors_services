package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Gdb *sql.DB

// docker run -d   --name mysql-container   -p 3308:3306   -e MYSQL_ROOT_PASSWORD=root   -e MYSQL_DATABASE=mydb   -e MYSQL_USER=myuser   -e MYSQL_PASSWORD=mypassword   mysql:8.0

func Db_connection() error {

	// load the env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}
	log.Println("os.Getenv(DB_NAME):", os.Getenv("DB_NAME"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	Gdb, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return err

	}
	// Connection pool settings (important for production)
	Gdb.SetMaxOpenConns(25)                 // max active connections
	Gdb.SetMaxIdleConns(25)                 // max idle connections
	Gdb.SetConnMaxLifetime(5 * time.Minute) // recycle connections

	err = Gdb.Ping()
	if err != nil {
		log.Printf("Failed to ping DB: %v", err)
		return err

	}
	log.Println("Database connection established")
	return nil
}
