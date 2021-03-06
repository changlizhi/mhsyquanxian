package moxings

// 方案：当用户点击登录时，帐号获取了多种验证类型，用多种验证类型进行校验时，循环
// 循环多种校验时需要多种值，此时再循环帐号对应出来的验证值，进行匹配，如果匹配成功则访问有效，
// 此时校验时有N种校验类型，也可能有N个值，所以一个用户登录时会for循环嵌套进行值的校验
// 注册时帐号类型会得到提示：邮箱，电话，微信等，如果没有就提示错误
// 登录时如果不给用户选择哪种类型就要循环出来然后根据一种用户类型把验证策略查出来，其实一种用户类型只会匹配一种验证类型？
// 不一定，比如邮箱，如果后期要实时登录校验启用动态密码，那就要密码和验证码两种校验方式支持。
// 所以用户类型和校验类型要多对多，但帐号不允许和校验类型进行对应
// 然后每一种校验值进行反向匹配得到匹配成功的条件
// 也就是说会需要三个for循环来匹配校验---其实如果用户授权了允许自己的帐号权限互通，则还是两个for循环

type Gl1he8s struct {
	Id          int64
	Biaobianma1 string `gorm:"not null"`
	Biaobianma8 string `gorm:"not null"`
}
