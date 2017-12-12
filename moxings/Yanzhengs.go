package moxings

type Yanzhengs struct {
	Id       int64
	Leixing  string // 密码、验证码、短信、微信，从类型表来？不应该，因为每种类型都会对应一种验证逻辑，所以必须有开发工作，做成常量更合适
	Zhi      string // 密码值、验证码策略、短信策略、openid策略
	Zhanghao string // 手机号、用户名、邮箱、微信号
}
