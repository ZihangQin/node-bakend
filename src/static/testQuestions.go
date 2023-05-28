package static

import "gorm.io/gorm"

type TestQuestions struct {
	gorm.Model
	Title           string
	Class           string
	Score           int
	TitleType       string
	Difficulty      string
	QuestionsSetter string
	Answer          string
}
