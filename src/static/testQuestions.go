package static

import "gorm.io/gorm"

type TestQuestions struct {
	gorm.Model
	Title           string `gorm:"NOT NULL" ` //题目
	Class           string `gorm:"NOT NULL" ` //班级
	Score           int    `gorm:"NOT NULL" ` //分值
	TitleType       string `gorm:"NOT NULL" ` //题目类型
	Difficulty      string `gorm:"NOT NULL" ` //难度系数
	QuestionsSetter string `gorm:"NOT NULL" ` // 出题人
	Answer          string `gorm:"NOT NULL" ` //答案
}
