package moxings

type Zhanghao1s struct {
	Id         int64
	Biaobianma string `gorm:"not null"` //用来做表关联snowflake
	Bianma     string `gorm:"not null"` //用来用户识别用户输入或者用拼音替代
}

func (Zhanghao1s) TableName() string {
	return "Zhanghao1s"
}
