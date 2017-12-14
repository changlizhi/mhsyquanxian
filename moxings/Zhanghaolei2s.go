package moxings

type Zhanghaolei2s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
}

func (Zhanghaolei2s) TableName() string {
	return "Zhanghaolei2s"
}
