package moxings

type Ziyuan6s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Biaoshuzi  int64  `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
	Lujing     string `gorm:"not null"`
	Fangfa     string `gorm:"not null"` //提交、删除
	Miaoshu    string `gorm:"not null;DEFAULT:0"`
}

func (Ziyuan6s) TableName() string {
	return "Ziyuan6s"
}
