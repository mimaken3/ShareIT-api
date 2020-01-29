package infrastructure

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func init() {
	var err error
	DBMS := "mysql"
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	DB, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal(err)
	}
}
