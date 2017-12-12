package moxings

type Jueseziyuans struct {
	Id           int64
	Juesebianma  string `gorm:"not null"`
	Ziyuanbianma string `gorm:"not null"`
}

func (Jueseziyuans) TableName() string {
	return "Jueseziyuans"
}
