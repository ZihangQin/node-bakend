package static

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	UserName  string `json:"userName"`                 //用户名
	Operate   string `json:"operate"`                  //操作
	HashValue string `gorm:"size:64" json:"hashValue"` //hash值
	Date      string `json:"date"`                     //操作日期
	Time      string `json:"time"`                     //操作时间
}
