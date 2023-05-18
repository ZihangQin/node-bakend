package utils

import (
	"bk/src/db"
	"bk/src/static"
	"errors"
	"regexp"
)

/**
判断是否为正确的手机号
*/
func IsMobile(mobile string) (bool, string) {
	result, _ := regexp.MatchString(`^(1[3-9][0-9]\d{4,8})$`, mobile)
	if result {
		return true, ""
	}
	return false, "手机号格式不正确"
}


/**
判断昵称是否相同
*/
func IsNameSame(name string) (bool, error) {
	var us []static.UserInfos
	db.DB.Where("user_name = ?", name).Find(&us)
	if len(us) <= 0 {
		return true, nil
	}
	return false, errors.New( "昵称已存在")
}