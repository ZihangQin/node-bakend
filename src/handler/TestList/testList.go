package TestList

import (
	"bk/src/db"
	"bk/src/static"
	"bk/src/utils"
	"fmt"
	"math"
	"sort"
)

type Test struct {
	Id              uint   `json:"id"`
	UpdateAt        string `json:"updateAt"`
	Title           string `json:"title"`
	Class           string `json:"class"`
	Score           int    `json:"score"`
	TitleType       string `json:"title_type"`
	Difficulty      string `json:"difficulty"`
	QuestionsSetter string `json:"questionsSetter"`
	Answer          string `json:"answer"`
}

//根据页数获取试题
func GetTestLists(pages int) (interface{}, int, error) {
	var tests []Test
	var databasesTest []static.TestQuestions
	// 计算记录总数和总页数
	var total int64
	if err := db.DB.Model(&static.TestQuestions{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(10)))

	// 分页查询试题数据
	err := db.DB.Offset((pages - 1) * 12).Limit(12).Find(&databasesTest).Error
	if err != nil {
		return nil, 0, err
	}
	for i := 0; i < len(databasesTest); i++ {
		times := utils.TimeFormat(databasesTest[i].UpdatedAt)
		test := Test{
			Id:              databasesTest[i].ID,
			UpdateAt:        times,
			Title:           databasesTest[i].Title,
			Class:           databasesTest[i].Class,
			Score:           databasesTest[i].Score,
			TitleType:       databasesTest[i].TitleType,
			Difficulty:      databasesTest[i].Difficulty,
			QuestionsSetter: databasesTest[i].QuestionsSetter,
			Answer:          databasesTest[i].Answer,
		}
		tests = append(tests, test)
	}

	return tests, totalPages, nil
}

//向数据库添加试题数据
func SetTest(title string, class string, score int, titleType string,
	difficulty string, questionsSetter string, answer string) error {
	test := static.TestQuestions{
		Title:           title,
		Class:           class,
		Score:           score,
		TitleType:       titleType,
		Difficulty:      difficulty,
		QuestionsSetter: questionsSetter,
		Answer:          answer,
	}
	return db.DB.Create(&test).Error
}

func DeleteTests(idLists map[string]string) {
	var idList []int

	for k, v := range idLists {
		if k == "myCheckbox" {
			continue
		}
		IntV, err := utils.StringToInt(v)
		if err != nil {
			return
		}
		idList = append(idList, IntV)
	}
	sort.Ints(idList)
	fmt.Println(idList)
	u := make([]static.TestQuestions, len(idList))
	for i := 0; i <= len(idList)-1; i++ {
		u[i].ID = uint(idList[i])
	}
	fmt.Println(u)
	db.DB.Model(&static.TestQuestions{}).Delete(&u)
}
