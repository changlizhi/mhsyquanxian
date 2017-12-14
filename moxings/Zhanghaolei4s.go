package moxings

type Zhanghaolei4s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
}

func (Zhanghaolei4s) TableName() string {
	return "Zhanghaolei4s"
}
