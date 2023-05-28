package TestList

import (
	"bk/src/db"
	"bk/src/static"
	"math"
)

//根据页数获取试题
func GetTestLists(pages int) ([]static.TestQuestions, int, error) {
	var tests []static.TestQuestions
	// 计算记录总数和总页数
	var total int64
	if err := db.DB.Model(&static.TestQuestions{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(10)))

	// 分页查询试题数据
	err := db.DB.Offset((pages - 1) * 10).Limit(10).Find(&tests).Error
	if err != nil {
		return nil, 0, err
	}
	return tests, totalPages, nil
}


//向数据库添加试题数据
func SetTest(title string, class string, score int, titleType string,
	difficulty string, questionsSetter string, answer string) error {
	test := static.TestQuestions{
		Title:      title,
		Class:      class,
		Score:      score,
		TitleType:  titleType,
		Difficulty: difficulty,
		QuestionsSetter:   questionsSetter,
		Answer:     answer,
	}
	return db.DB.Create(&test).Error
}
