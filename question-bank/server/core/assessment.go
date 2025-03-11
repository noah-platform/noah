package core

type WritingTask string

const (
	WritingTaskOne WritingTask = "1"
	WritingTaskTwo WritingTask = "2"
)

type WritingAssessmentInput struct {
	Task         WritingTask `json:"writingTask"`
	QuestionBody string      `json:"questionBody"`
	Answer       string      `json:"answer"`
	ImageUrls    []string    `json:"imageUrls"`
}

type WritingAssessmentOutput struct {
	TaskAchievementScore               string `json:"taskAchievementScore"`
	TaskAchievementComment             string `json:"taskAchievementComment"`
	CoherenceAndCohesionScore          string `json:"coherenceAndCohesionScore"`
	CoherenceAndCohesionComment        string `json:"coherenceAndCohesionComment"`
	LexicalResourceScore               string `json:"lexicalResourceScore"`
	LexicalResourceComment             string `json:"lexicalResourceComment"`
	GrammaticalRangeAndAccuracyScore   string `json:"grammaticalRangeAndAccuracyScore"`
	GrammaticalRangeAndAccuracyComment string `json:"grammaticalRangeAndAccuracyComment"`
	OverallScore                       string `json:"overallScore"`
	OverallComment                     string `json:"overallComment"`
}
