package moxings

type Weixinzhanghaos struct {
	Id        int64
	Openid    string `gorm:"not null"`
	Weixinhao string `gorm:"not null;DEFAULT:0"` //如果微信号不存在则，
}

func (Weixinzhanghaos) TableName() string {
	return "Weixinzhanghaos"
}
