package moxings

type Gl2he4s struct {
	Id          int64
	Biaobianma2 string `gorm:"not null"`
	Biaobianma4 string `gorm:"not null"`
}

func (Gl2he4s) TableName() string {
	return "Gl2he4s"
}
