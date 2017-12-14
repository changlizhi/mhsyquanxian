package moxings

type Yanzheng2s struct {
	Id         int64
	Biaobianma string
	Bianma     string
	Mingcheng  string
	Celue      string
}

func (Yanzheng2s) TableName() string {
	return "Yanzheng2s"
}
