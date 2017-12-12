package moxings

type Jueseyonghus struct {
	Id           int64
	Juesebianma  string `gorm:"not null"`
	Yonghubianma string `gorm:"not null"`
}

func (Jueseyonghus) TableName() string {
	return "Jueseyonghus"
}
