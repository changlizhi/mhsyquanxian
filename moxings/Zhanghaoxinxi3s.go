package moxings

type Zhanghaoxinxi3s struct {
	Id          int64
	Biaobianma1 string `gorm:"not null"`
	Biaobianma7 string `gorm:"not null"`
	Zhi         string `gorm:"not null"`
}

func (Zhanghaoxinxi3s) TableName() string {
	return "Zhanghaoxinxi3s"
}
