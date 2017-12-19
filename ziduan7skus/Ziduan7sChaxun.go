package ziduan7skus

import (
	"log"
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)
//返回mx中条件的数据，若mx的数据都为""或0的话就返回Id最大的一条记录
func ChaxunZiduan7s(mx moxings.Ziduan7s) *moxings.Ziduan7s {
	find := quanju.Db().Find(&moxings.Ziduan7s{}, mx)
	if find.Error != nil {
		log.Println("ChaxunZiduan7s---find.Error---", find.Error)
		return nil
	}
	ret := find.Value.(*moxings.Ziduan7s)
	return ret
}
func SuoyouZiduan7s() *[]moxings.Ziduan7s {
	mxs := []moxings.Ziduan7s{}
	find := quanju.Db().Find(&mxs)
	return find.Value.(*[]moxings.Ziduan7s)
}