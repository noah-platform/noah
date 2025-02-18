package core

type TestEntry struct {
	ID          string        `json:"testId"`
	Module      string        `json:"module"`
	QuestionSet []interface{} `json:"questionSet"`
	Audio       string        `json:"audio"`
	AudioScript string        `json:"audioScript"`
}

type TestInfo struct {
	ID     string `json:"testId"`
	Module string `json:"module"`
}
