package TestList

import (
	"bk/src/constant"
	"bk/src/db"
	"bk/src/static"
	"bk/src/utils"
	"errors"
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
func GetTestLists(pages int, token string) (interface{}, int, error) {
	var tests []Test
	var databasesTest []static.TestQuestions

	//进行token验证
	us, err := utils.VerifyToken(token, constant.SECRET)
	if err != nil {
		return nil, 0, err
	}
	if us == nil {
		return nil, 0, errors.New("token验证失败")
	}
	// 计算记录总数和总页数
	var total int64
	if err := db.DB.Model(&static.TestQuestions{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(12)))

	// 分页查询试题数据
	err = db.DB.Offset((pages - 1) * 12).Limit(12).Find(&databasesTest).Error
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
func SetTest(title string, class string, score string, titleType string,
	difficulty string, answer string, token string) error {
	if title == "" || class == "" || score == "" || titleType == "" || difficulty == "" || answer == "" || token == "" {
		return errors.New("参数不可为空")
	}
	//解析token
	u, err := utils.VerifyToken(token, constant.SECRET)
	if err != nil {
		return err
	}
	scoreInt, err := utils.StringToInt(score)
	if err != nil {
		return err
	}
	test := static.TestQuestions{
		Title:           title,
		Class:           class,
		Score:           scoreInt,
		TitleType:       titleType,
		Difficulty:      difficulty,
		QuestionsSetter: u.Username,
		Answer:          answer,
	}
	return db.DB.Create(&test).Error
}

//删除试题
func DeleteTests(idLists map[string]string) error {
	var idList []int
	if idLists == nil {
		return errors.New("参数不能为空")
	}
	for k, v := range idLists {
		if k == "myCheckbox" {
			continue
		}
		IntV, err := utils.StringToInt(v)
		if err != nil {
			return err
		}
		idList = append(idList, IntV)
	}
	sort.Ints(idList)
	u := make([]static.TestQuestions, len(idList))
	for i := 0; i <= len(idList)-1; i++ {
		u[i].ID = uint(idList[i])
	}
	err := db.DB.Model(&static.TestQuestions{}).Delete(&u).Error
	if err != nil {
		return err
	}

	return nil
}

//搜索试题
func SearchTests(data string) ([]Test, int, error) {
	if data == "" {
		return nil, 0, errors.New("搜索内容不能为空！")
	}

	var tests []static.TestQuestions
	err := db.DB.Model(&static.TestQuestions{}).Where("id LIKE ?", data).
		Or("title LIKE ?", data).Or("title_type LIKE ?", data).
		Or("difficulty LIKE ?", data).Or("questions_setter LIKE ?", data).Limit(12).Find(&tests).Error

	if err != nil {
		return nil, 0, err
	}

	// 计算记录总数和总页数
	var total int64
	if err := db.DB.Model(&static.TestQuestions{}).Where("id LIKE ?", data).
		Or("title LIKE ?", data).Or("title_type LIKE ?", data).
		Or("difficulty LIKE ?", data).Or("questions_setter LIKE ?", data).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(10)))

	var t []Test
	for i := 0; i < len(tests); i++ {
		times := utils.TimeFormat(tests[i].UpdatedAt)
		test := Test{
			Id:              tests[i].ID,
			UpdateAt:        times,
			Title:           tests[i].Title,
			Class:           tests[i].Class,
			Score:           tests[i].Score,
			TitleType:       tests[i].TitleType,
			Difficulty:      tests[i].Difficulty,
			QuestionsSetter: tests[i].QuestionsSetter,
			Answer:          tests[i].Answer,
		}
		t = append(t, test)
	}

	fmt.Println(t, totalPages)

	return t, totalPages, nil
}

func UpdateTests(id string, types string, grade string, content string, difficulty string,
	score string, answer string, token string) error {
	_, err := utils.VerifyToken(token, constant.SECRET)
	if err != nil {
		return err
	}
	if id == "" {
		return errors.New("请指定要修改的试题序号")
	}

	scoreInt, err := utils.StringToInt(score)
	if err != nil {
		return err
	}

	test := static.TestQuestions{
		Title:      content,
		Class:      grade,
		Score:      scoreInt,
		TitleType:  types,
		Difficulty: difficulty,
		Answer:     answer,
	}
	idInt, _ := utils.StringToInt(id)
	db.DB.Model(&static.TestQuestions{}).Where("id = ?", idInt).Updates(&test)
	return nil
}
