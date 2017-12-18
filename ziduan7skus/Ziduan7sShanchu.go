package ziduan7skus

import (
	"log"
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func ShanchuZiduan7s(mx moxings.Ziduan7s) {
	find := quanju.Db().Find(&moxings.Ziduan7s{}, mx)
	if find.Error != nil {
		log.Println("ShanchuZiduan7s---find.Error---", find.Error)
	}
	cxmx := find.Value.(*moxings.Ziduan7s)
	delete := find.Delete(&cxmx)
	log.Println(delete)
}
