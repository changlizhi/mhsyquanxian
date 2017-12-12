package moxings

// 密码、验证码、短信、微信，从类型表来？
// 不应该，因为每种类型都会对应一种验证逻辑，所以必须有开发工作，做成常量更合适
// 密码值、验证码策略、短信策略、openid策略
// 手机号、用户名、邮箱、微信号
type Yanzhengs struct {
	Id       int64
	Leixing  string `gorm:"not null"`
	Zhi      string `gorm:"not null"`
	Zhanghao string `gorm:"not null"`
}

func (Yanzhengs) TableName() string {
	return "Yanzhengs"
}
