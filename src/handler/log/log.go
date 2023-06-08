package log

import (
	"bk/src/constant"
	"bk/src/utils"
)

func GetLog(token string) {
	userInfo,err := utils.VerifyToken(token,constant.SECRET)
	if err!=nil && userInfo != nil{
		return
	}


}