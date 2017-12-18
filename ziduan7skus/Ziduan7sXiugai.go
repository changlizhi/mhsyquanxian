package ziduan7skus

import (
	"mhsyquanxian/quanju"
	"mhsyquanxian/moxings"
	"log"
)

func XiugaiZiduan7s(mx moxings.Ziduan7s) *moxings.Ziduan7s {
	if mx.Biaobianma == "" {
		log.Println("ShanchuZiduan7s---表编码不能为空")
		return nil
	}
	mx.Id = 0// 不允许带id传过来，但如果传过来了也没关系，只会根据表编码修改
	log.Println("mx111-----", mx)
	mxcx := &moxings.Ziduan7s{
		Biaobianma : mx.Biaobianma,
	}
	find := quanju.Db().Find(&moxings.Ziduan7s{}, mxcx)
	if find.Error != nil {
		log.Println("ShanchuZiduan7s---find.Error---", find.Error)
		return nil
	}
	cxmx := find.Value.(*moxings.Ziduan7s)
	mx.Id = cxmx.Id
	log.Println("mx1-----", mx)
	gengxin := find.Update(&mx)
	ret := gengxin.Value.(*moxings.Ziduan7s)
	log.Println(gengxin.Value)
	return ret
}
