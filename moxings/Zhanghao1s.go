package moxings

// 表编码用来做表关联snowflake
// 编码用来做统一编码，可以用拼音，也可以用邮箱，策略比较随意，
// 但如果用户填了别的可用做登录的帐号信息则这个值是相同的
// 值就是跟表编码一一对应的
// 所以表编码是用来认证的，而编码使用来授权的，不需要专门的统一编码表
// 如果同一个编码下的多条记录都允许帐号统一授权则用需要循环每个令牌来进行匹配，而且这里默认是真.

type Zhanghao1s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Shouquan   int    `gorm:"not null"`
	Zhi        string `gorm:"not null"`
}

func (Zhanghao1s) TableName() string {
	return "Zhanghao1s"
}
