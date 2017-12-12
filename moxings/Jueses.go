package moxings

type Jueses struct {
	Id          int64
	Juesebianma string `gorm:"not null"`
	Mingcheng   string `gorm:"not null"`
	Miaoshu     string `gorm:"not null;DEFAULT:0"`
}

func (Jueses) TableName() string {
	return "Jueses"
}
