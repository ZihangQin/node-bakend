package browse

import (
	"bk/src/constant"
	"bk/src/db"
	"bk/src/static"
	"bk/src/utils"
)

//首页获取用户信息方法
func GetUserInfo(token string) (string, int64, error) {
	userInfo, err := utils.VerifyToken(token, constant.SECRET)
	if err != nil {
		return "", 0, err
	}
	//fmt.Println(userInfo.Id)
	//fmt.Println(userInfo.Username,userInfo.UserID)
	//通过id获取积分
	var user static.UserInfos
	err = db.DB.Model(&static.UserInfos{}).Where("id = ?", userInfo.UserID).Find(&user).Error
	if err != nil {
		return "", 0, err
	}
	//fmt.Println(user)
	return user.UserName, user.Calculus, nil
}
