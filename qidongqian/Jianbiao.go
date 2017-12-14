package qidongqian

import (
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func Chuangjianbiao() {
	quanju.Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(
		&moxings.Gl1he4s{},
		&moxings.Gl1he5s{},
		&moxings.Gl2he4s{},
		&moxings.Gl5he6s{},
		&moxings.Juese5s{},
		&moxings.Yanzheng2s{},
		&moxings.Zhanghao1s{},
		&moxings.Zhanghaolei4s{},
		&moxings.Zhanghaoxinxi3s{},
		&moxings.Ziduan7s{},
		&moxings.Ziyuan6s{},
	)
}

