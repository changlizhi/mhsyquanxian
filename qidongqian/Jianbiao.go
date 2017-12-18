package qidongqian

import (
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func Chuangjianbiao() {
	dbs := quanju.Db().Set("gorm:table_options", "ENGINE=InnoDB")
	defer dbs.Close()
	dbs.CreateTable(
		&moxings.Gl1he8s{},
		&moxings.Lingpai8s{},
		&moxings.Lingpailei9s{},
	)
}
