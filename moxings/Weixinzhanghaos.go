package moxings

type Weixinzhanghaos struct {
	Id           int64
	Openid       string
	Weixinhao    string
	Yonghubianma string
}

func (Weixinzhanghaos) TableName() string {
	return "Weixinzhanghaos"
}
