package moxings

// 此表的数据讲清楚每种令牌的组合规则和校验方式
// 密码，就只要一个就可以了，当然也可以邮箱是邮箱的密码，
// 手机是手机的密码，如果帐号匹配了就可以了。
// 如果手机密码和邮箱密码要用一个就关联到一个去。
// 校验的时候有区分，校验完了还是要用一个统一帐号去做权限分配
// 所以还要对这个区分进行说明，而这个说明是与帐号在一起的。所以要用一个帐号唯一表来统领这个帐号表，
// 如果在这个统领表内有关联的就还要把关联的拿出来查询，如果匹配了任何一个就通过，
// 因为很可能会把手机的密码记成手机的密码，但这个规则要用户在功能里设置才能过，
// 如果是敏感帐号则不允许用户这样设置，必须用两种密码。
// 虽然最后登录之后看到的信息和权限很可能是需要互通的（也可以设置不同的帐号是否互通），
// 这要在角色权限去设置了，设置的时候要把用户信息进行查询转换
// 这个表里最重要的就是对令牌进行说明，建立所有的规则

type Lingpailei9s struct {
	Id         int64
	Biaobianma string `gorm:"not null"`
	Bianma     string `gorm:"not null"`
	Mingcheng  string `gorm:"not null"`
	Guize      string `gorm:"not null"`
	Miaoshu    string `gorm:"not null"`
}
