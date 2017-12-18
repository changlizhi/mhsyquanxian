package qidongqian

import (
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
	"log"
)

func Chuangjianbiao() {
	dbs := quanju.Db().Set("gorm:table_options", "ENGINE=InnoDB")
	defer dbs.Close()
	dbs.CreateTable(
		&moxings.Biaoji10s{},
		&moxings.Gl1he2s{},
		&moxings.Gl1he5s{},
		&moxings.Gl1he8s{},
		&moxings.Gl2he4s{},
		&moxings.Gl5he6s{},
		&moxings.Gl8he9s{},
		&moxings.Juese5s{},
		&moxings.Lingpai8s{},
		&moxings.Lingpailei9s{},
		&moxings.Yanzheng4s{},
		&moxings.Zhanghao1s{},
		&moxings.Zhanghaolei2s{},
		&moxings.Zhanghaoxinxi3s{},
		&moxings.Ziduan7s{},
		&moxings.Ziyuan6s{},
	)
	if dbs.Error != nil{
		log.Println("Chuangjianbiao---dbs.Error---",dbs.Error)
	}
}
