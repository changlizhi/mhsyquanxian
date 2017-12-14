package moxings

type Juese5s struct {
	Id          int64
	Juesebianma string `gorm:"not null"`
	Mingcheng   string `gorm:"not null"`
	Miaoshu     string `gorm:"not null;DEFAULT:0"`
}

func (Juese5s) TableName() string {
	return "Juese5s"
}
