package moxings

type Yanzheng2s struct {
	Id          int64
	Biaobianma  string `gorm:"not null"`
	Bianma      string `gorm:"not null"`
	Mingcheng   string `gorm:"not null"`
	Celue       string `gorm:"not null"`
	Celuefangfa string `gorm:"not null"`
}

func (Yanzheng2s) TableName() string {
	return "Yanzheng2s"
}
