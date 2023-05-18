package account

import (
	"bk/src/db"
	"bk/src/static"
	"bk/src/utils"
	"errors"
	"fmt"
	"strconv"
)

func LoginAccount(name string, password string) (bool, string, error) {
	var account static.UserInfos

	//合法性校验
	if name == "" || password == "" {
		return false, "", errors.New("账号密码非法")
	}

	err := db.DB.Find(&account, "user_name = ?", name).Error
	if err != nil {
		fmt.Println(err)
		return false, "", err
	}

	if account.UserPassword != password {
		return false, "", errors.New("账号或密码错误")
	}

	token, err := utils.GenerateToken("123446", strconv.FormatUint(uint64(account.Id), 10), account.UserName)
	if err != nil {
		fmt.Println(err)
		return false, "", err
	}

	return true, token, nil
}

func RegisterAccount(name string, phone string, password string, nickName string)  {

}
