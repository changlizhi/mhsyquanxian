package quanju

import (
	"github.com/jinzhu/gorm"
	"mhsyquanxian/gongjus"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Db() *gorm.DB {
	Db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyquanxian?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("kaiqi shujuku shibai ---", err)
	}
	log.Println("Shujuku111-----", Db)
	return Db
}

func FlakeId() *gongjus.IdWorker {
	log.Println(111)

	Flakeid, err := gongjus.NewIdWorker()
	if err != nil {
		log.Println("flakeid shengcheng shibai", err)
	}
	log.Println("Flakeid-----", Flakeid)
	return Flakeid
}
