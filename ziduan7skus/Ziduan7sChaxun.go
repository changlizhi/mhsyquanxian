package ziduan7skus

import (
	"log"
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func ChaxunZiduan7s(mx moxings.Ziduan7s) *moxings.Ziduan7s {
	find := quanju.Db().Find(&moxings.Ziduan7s{}, mx)
	if find.Error != nil {
		log.Println("ChaxunZiduan7s---find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Ziduan7s)
	return ret
}