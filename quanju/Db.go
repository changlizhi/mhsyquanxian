package quanju

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *gorm.DB

func Db() *gorm.DB {
	if db == nil {
		log.Println(111)
		db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyquanxian?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Println("kaiqi shujuku shibai ---", err)
		}
		return db
	}
	log.Println("Shujuku111-----", db)
	return db
}
