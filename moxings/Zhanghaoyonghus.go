package moxings

type Zhanghaoyonghus struct {
	Id             int64
	Zhanghaobianma string `gorm:"not null"`
	Yonghubianma   string `gorm:"not null"`
}

func (Zhanghaoyonghus) TableName() string {
	return "Zhanghaoyonghus"
}
