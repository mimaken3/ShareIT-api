package infrastructure

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	var err error
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "shareit?parseTime=true"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	DB, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Fatal(err)
	}
}
