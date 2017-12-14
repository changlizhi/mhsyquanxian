package moxings

type Ziduan7s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
	Leixing    string `gorm:"not null"`
	Changdu    int `gorm:"not null"`
}

func (Ziduan7s) TableName() string {
	return "Ziduan7s"
}
