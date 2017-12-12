package moxings

type Jueseziyuans struct {
	Id           int64
	Juesebianma  string
	Ziyuanbianma string
}

func (Jueseziyuans) TableName() string {
	return "Jueseziyuans"
}
