package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB

// init 為特別的函式，只會在程式初始化時執行一次
func init() {
	var err error
	db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"),
		os.Getenv("DBNAME")))
	fmt.Println(os.Getenv("DBUSER"), ":", os.Getenv("DBPASSWORD"), "@tcp(", os.Getenv("DBHOST")+":"+os.Getenv("DBPORT"), ")/", os.Getenv("DBNAME"))

	if err != nil {
		log.Println(err)
	}
}

func CloseDB() {
	defer db.Close()
}

func Get() *gorm.DB {
	return db
}
