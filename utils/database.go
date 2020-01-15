package utils

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/kshvakov/clickhouse"
	"log"
	"os"
)

var db *sqlx.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbUser := os.Getenv("db_user")
	dbPassword := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbDriver := os.Getenv("db_driver")
	dbType := os.Getenv("db_type")
	dbPort := os.Getenv("db_port")

	dbUri := dbType + "://" + dbHost + ":" + dbPort + "?" + "username=" + dbUser + "&password=" + dbPassword + "&database=" + dbName + "&debug=true"

	log.Println(dbUri)

	conn, err := sqlx.Open(dbDriver, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *sqlx.DB {
	return db
}
