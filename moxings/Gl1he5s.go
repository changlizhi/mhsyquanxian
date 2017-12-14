package moxings

type Gl1he5s struct {
	Id          int64
	Biaobianma1 string `gorm:"not null"`
	Biaobianma5 string `gorm:"not null"`
}
func (Gl1he5s) TableName() string {
	return "Gl1he5s"
}
