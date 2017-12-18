package moxings

// 标记表，对所有有标记意义的值进行罗列,Zhi是用一个以冒号分割值和名称的字符串，用的字段名必须从Ziduan7s里取，
// 父编码的意义就是一个label，编码的意义就是一个父编码下有很多值，进行删除的时候方便点。
type Biaoji10s struct {
	Id          int64
	Biaobianma  string
	Bianma      string
	Fubianma    string
	Biaobianma7 string
	Zhi         string
	Miaoshu     string
}
