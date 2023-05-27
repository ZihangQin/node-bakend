package static

/*
用户个人信息
*/

type UserInfos struct {
	Id           uint   `gorm:"primaryKey" json:"id"` //用户uid
	UserName     string `json:"username"`             //用户昵称
	UserPassword string `json:"userpassword"`         //用户密码
	Emil         string `json:"emil"`                 //用户邮箱
	Phone        string `json:"phone"`                //手机号
	State        int    `json:"state"`                //用户账号权限，0代表普通用户，1代表特殊用户
	Calculus     int64  `json:"calculus"`             //用户积分
}
