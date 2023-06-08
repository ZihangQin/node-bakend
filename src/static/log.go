package static

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	Uid       uint   `json:"uid"`        //用户id
	Operate   string `json:"operate"`    //操作
	HashValue string `json:"hash_value"` //hash值
	IsTrue    bool   `json:"is_true"`    //真实性
}
