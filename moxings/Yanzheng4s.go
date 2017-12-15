package moxings

type Yanzheng4s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
	Jiami      string `gorm:"not null"`
	Fangfa     string `gorm:"not null"`
}

func (Yanzheng4s) TableName() string {
	return "Yanzheng4s"
}
