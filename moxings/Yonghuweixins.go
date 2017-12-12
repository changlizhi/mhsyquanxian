package moxings

type Yonghuweixins struct {
	Id           int64
	Weixinhao    string `gorm:"not null"` //微信的名称肯定是有的，但如果没有就默认为Openid
	Yonghubianma string `gorm:"not null"`
}

func (Yonghuweixins) TableName() string {
	return "Yonghuweixins"
}
