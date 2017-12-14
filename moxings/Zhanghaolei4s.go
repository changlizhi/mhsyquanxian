package moxings

type Zhanghaolei4s struct {
	Id         int64
	Biaobianma string
	Bianma     string
	Mingcheng  string
}

func (Zhanghaolei4s) TableName() string {
	return "Zhanghaolei4s"
}
