package ceshiyong

import "mhsyquanxian/moxings"

func Ziduan7scharu() *moxings.Ziduan7s {
	ret := &moxings.Ziduan7s{
		Id:         1000,
		Biaobianma: "zd1000",
		Bianma:     "csziduan",
		Mingcheng:  "测试字段",
		Leixing:    "int",
		Changdu:    0,
	}
	return ret
}

func Ziduan7schaxun() moxings.Ziduan7s {
	ret := moxings.Ziduan7s{
		Biaobianma: "zd1000",
	}
	return ret
}
