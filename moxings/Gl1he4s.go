package moxings

type Gl1he4s struct {
	Id          int64
	Biaobianma1 string `gorm:"not null"`
	Biaobainam4 string `gorm:"not null"`
}

func (Gl1he4s) TableName() string {
	return "Gl1he4s"
}
