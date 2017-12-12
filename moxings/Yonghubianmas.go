package moxings

type Yonghubianmas struct {
	Id           int64
	Yonghubianma string
}

func (Yonghubianmas) TableName() string {
	return "Yonghubianmas"
}
