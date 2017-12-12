package moxings

type Yonghuzhanghaos struct {
	Id         int64
	Shoujihao  int64
	Yonghuming string
	Youxiang   string
	Yonghubianma string
}

func (Yonghuzhanghaos) TableName() string {
	return "Yonghuzhanghaos"
}
