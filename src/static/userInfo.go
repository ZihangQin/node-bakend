package static

/*
用户个人信息
*/

type UserInfos struct {
	Id           uint   `gorm:"primaryKey" json:":id"`
	UserName     string `json:"username"`
	UserPassword string `json:"userpassword"`
}
