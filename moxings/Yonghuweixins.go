package moxings

type Yonghuweixins struct {
	Id           int64
	Weixinhao    string
	Yonghubianma string
}

func (Yonghuweixins) TableName() string {
	return "Yonghuweixins"
}
