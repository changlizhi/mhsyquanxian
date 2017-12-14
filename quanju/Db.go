package quanju

import (
	"github.com/jinzhu/gorm"
	"mhsyquanxian/gongjus"
)

var Db, _ = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mhsyquanxian?charset=utf8&parseTime=True&loc=Local")
var Idworker, _ = gongjus.NewIdWorker()

