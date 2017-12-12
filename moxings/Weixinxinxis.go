package moxings

type Weixinxinxis struct {
	Id           int
	OpenId       string `gorm:"not null;column:openid"` // 用户的唯一标识
	Nickname     string `gorm:"column:nickname"`        // 用户昵称
	Weixinhao    string `gorm:"column:weixinhao"`       // 用户昵称
	Sex          int    `gorm:"column:sex"`             // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	City         string `gorm:"column:city"`            // 普通用户个人资料填写的城市
	Province     string `gorm:"column:province"`        // 用户个人资料填写的省份
	Country      string `gorm:"column:country"`         // 国家, 如中国为CN
	HeadImageURL string `gorm:"column:headimgurl"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像）， 用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Privilege    string `gorm:"column:privilege"`       // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionId      string `gorm:"column:unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

func (Weixinxinxis) TableName() string {
	return "Weixinxinxis"
}
