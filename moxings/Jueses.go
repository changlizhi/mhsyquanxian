package moxings

type Jueses struct {
	Id          int64
	Juesebianma string
	Mingcheng   string
	Miaoshu     string
}

func (Jueses) TableName() string {
	return "Jueses"
}
