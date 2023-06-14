package logger

import (
	"bk/src/db"
	"bk/src/static"
	"math"
)
/*将操作日志保存在区块链中*/
func SetLogger(data string, time string, opName string, username string, hash string) error {
	//将日志操作保存在数据库中
	logger := static.Log{
		UserName:  username,
		Operate:   opName,
		HashValue: hash,
		Date:      data,
		Time:      time,
	}
	err := db.DB.Model(&static.Log{}).Create(&logger).Error
	if err != nil {
		return err
	}

	return nil
}


/*将数据库中的操作日志获取出来*/
func GetLogger(pages int) ([]static.Log,int,error) {
	var log []static.Log
	//计算总页数
	var total int64
	if err := db.DB.Model(&static.Log{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(12)))
	err := db.DB.Offset((pages - 1) * 12).Limit(12).Model(&static.Log{}).Find(&log).Error
	if err != nil {
		return nil,0,err
	}
	return log,totalPages, nil
}