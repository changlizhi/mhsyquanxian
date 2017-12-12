package moxings

type Jueseyonghus struct {
	Id           int64
	Juesebianma  string
	Yonghubianma string
}

func (Jueseyonghus) TableName() string {
	return "Jueseyonghus"
}
