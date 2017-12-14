package moxings

type Zhanghaoxinxi3s struct {
	Id          int64
	Biaobianma1 string
	Biaobianma7 string
	Zhi         string
}

func (Zhanghaoxinxi3s) TableName() string {
	return "Zhanghaoxinxi3s"
}
