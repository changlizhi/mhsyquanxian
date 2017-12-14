package moxings

type Gl1he2s struct {
	Id          int64
	Biaobianma1 string `gorm:"not null"`
	Biaobianma2 string `gorm:"not null"`
}

func (Gl1he2s) TableName() string {
	return "Gl1he2s"
}
