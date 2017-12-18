package ziduan7skus

import (
	"log"
	"mhsyquanxian/moxings"
	"mhsyquanxian/quanju"
)

func ZengjiaZiduan7s(mx *moxings.Ziduan7s) bool {
	cr := quanju.Db().Create(mx)
	if cr.Error != nil {
		log.Println("cr.Error---", cr.Error)
		return false
	}
	return true
}
