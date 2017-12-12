package moxings

type Ziyuans struct {
	Id           int64
	Mingcheng    string
	Lujing       string
	Ziyuanbianma string
	Ziyuanshuzi  int64
}

func (Ziyuans) TableName() string {
	return "Ziyuans"
}
