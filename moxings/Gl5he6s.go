package moxings

type Gl5he6s struct {
	Id          int64
	Biaobianma5 string `gorm:"not null"`
	Biaobianma6 string `gorm:"not null"`
}

func (Gl5he6s) TableName() string {
	return "Gl5he6s"
}
