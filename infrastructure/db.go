package infrastructure

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/appengine"
)

var DB *gorm.DB

func init() {
	var err error
	DBMS := "mysql"
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}

	cloudSQLConnection := os.Getenv("CLOUD_SQL_CONNECTION")
	localSQLConnection := os.Getenv("LOCAL_SQL_CONNECTION")
	if appengine.IsAppEngine() {
		// GAE
		DB, err = gorm.Open(DBMS, cloudSQLConnection)
	} else {
		// Local
		DB, err = gorm.Open(DBMS, localSQLConnection)
	}

	if err != nil {
		log.Fatal(err)
	}
}
