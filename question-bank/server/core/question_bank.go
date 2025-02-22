package core

type TestEntry struct {
	ID          string         `json:"testId"`
	Title       string         `json:"title"`
	Duration    int32          `json:"duration"`
	Instruction string         `json:"instruction"`
	Module      string         `json:"module"`
	Sections    []SectionEntry `json:"sections"`
}

type SectionEntry struct {
	ID            string             `json:"sectionId"`
	Title         string             `json:"title"`
	Instruction   string             `json:"instruction"`
	AudioUrl      string             `json:"audioUrl"`
	QuestionCount int32              `json:"questionCount"`
	QuestionSet   []QuestionSetEntry `json:"questionSet"`
}

type QuestionSetEntry struct {
	Instruction  string          `json:"instruction"`
	ResponseType string          `json:"responseType"`
	Questions    []QuestionEntry `json:"questions"`
}

type QuestionEntry struct {
	Body      string   `json:"body"`
	ImageUrls []string `json:"imageUrls"`
}
type TestInfo struct {
	ID     string `json:"testId"`
	Module string `json:"module"`
}
