package qidongqian

import (
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func Chuangjianbiao() {
	dbs := quanju.Db().Set("gorm:table_options", "ENGINE=InnoDB")
	defer dbs.Close()
	dbs.CreateTable(
		&moxings.Gl1he2s{},
		&moxings.Gl2he4s{},
		&moxings.Gl1he5s{},
		&moxings.Gl5he6s{},
		&moxings.Juese5s{},
		&moxings.Yanzheng4s{},
		&moxings.Zhanghao1s{},
		&moxings.Zhanghaolei2s{},
		&moxings.Zhanghaoxinxi3s{},
		&moxings.Ziduan7s{},
		&moxings.Ziyuan6s{},
	)
}
