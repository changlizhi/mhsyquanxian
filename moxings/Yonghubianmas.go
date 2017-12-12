package moxings

type Yonghubianmas struct {
	Id           int64
	Yonghubianma string `gorm:"not null"`
}

func (Yonghubianmas) TableName() string {
	return "Yonghubianmas"
}
