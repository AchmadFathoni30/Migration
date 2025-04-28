package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/denisenkom/go-mssqldb"
)

var DB *gorm.DB
var MSSQL *sql.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitPostgres()
	InitSQLServer()
}

func InitPostgres() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB"),
		os.Getenv("PG_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}
}

func InitSQLServer() {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		os.Getenv("MSSQL_HOST"),
		os.Getenv("MSSQL_USER"),
		os.Getenv("MSSQL_PASSWORD"),
		os.Getenv("MSSQL_DB"),
	)

	var err error
	MSSQL, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("SQL Server connection failed: %v", err)
	}

	if err = MSSQL.Ping(); err != nil {
		log.Fatalf("SQL Server ping failed: %v", err)
	}
}
