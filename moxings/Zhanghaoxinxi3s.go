package moxings

type Zhanghaoxinxi3s struct {
	Id       int64
	Xinxilei string
	Zhi      string
}

func (Zhanghaoxinxi3s) TableName() string {
	return "Zhanghaoxinxi3s"
}
