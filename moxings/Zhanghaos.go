package moxings

type Zhanghaos struct {
	Id             int64
	Zhanghaobianma int64  `gorm:"not null"`
	Shoujihao      int64  `gorm:"not null;DEFAULT:0"`
	Yonghuming     string `gorm:"not null;DEFAULT:0"`
	Youxiang       string `gorm:"not null;DEFAULT:0"`
}

func (Zhanghaos) TableName() string {
	return "Zhanghaos"
}
