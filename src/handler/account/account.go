package account

import (
	"bk/src/db"
	"bk/src/static"
	"bk/src/utils"
	"errors"
	"fmt"
	"strconv"
)
//登录方法
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

//注册方法
func RegisterAccount(name string, phone string, password string, email string) (bool,error) {
	//合法性校验
	if name == "" || password == "" || email == ""{
		return false, errors.New("参数非法")
	}

	//昵称合法性
	isSameName ,err := utils.IsNameSame(name)
	if err != nil {
		return false,errors.New("名称查询错误")
	}
	if !isSameName {
		return false, errors.New("该用户名已存在请更换")
	}

	//邮箱合法性验证
	_, err = utils.IsEmil(email)
	if err != nil {
		return false,err
	}

	//手机号合法性
	if phone != ""{
		isPhone ,err := utils.IsMobile(phone)
		if err!=nil {
			return  false,  errors.New("手机号不合法")
		}
		if !isPhone {
			return false,  errors.New("手机号格式错误")
		}
	}

	//将正确的注册数据保存进数据库
	users :=  static.UserInfos{
		UserName:     name,
		UserPassword: password,
		Emil:         email,
		Phone:        phone,
		State:        0,
	}
	err = db.DB.Model(&static.UserInfos{}).Create(&users).Error
	if err != nil {
		return false, err
	}


	return true, nil
}
