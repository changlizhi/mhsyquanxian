package moxings

import "time"

//令牌类型：比如有的校验需要多个值，当值的策略发生改变时就需要这个标记
//比如验证码校验，如果验证码发送之后3分钟，就失效，也要先用类型标记做判断,否则时间值就是简单的显示值生成时间

type Lingpai8s struct {
	Id          int64
	Biaobianma  string    `gorm:"not null"`
	Bianma      string    `gorm:"not null"`
	Zhi         string    `gorm:"not null"`
	Biaobianma9 string    `gorm:"not null"` // 令牌类型：必须令牌类型表中去选
	Shijian     time.Time `gorm:"not null;DEFAULT:'1970-01-01 10:00:00'"`
}
