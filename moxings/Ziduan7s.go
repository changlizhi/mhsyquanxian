package moxings

type Ziduan7s struct {
	Id         int64
	Biaobianma string
	Bianma     string
	Mingcheng  string
	Leixing    string
	Changdu    int
}

func (Ziduan7s) TableName() string {
	return "Ziduan7s"
}
