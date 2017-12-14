package moxings

type Ziyuan6s struct {
	Id           int64
	Mingcheng    string `gorm:"not null"`
	Lujing       string `gorm:"not null"`
	Ziyuanbianma string `gorm:"not null"`
	Ziyuanshuzi  int64  `gorm:"not null"`
	Miaoshu      string `gorm:"not null;DEFAULT:0"`
}

func (Ziyuan6s) TableName() string {
	return "Ziyuan6s"
}
