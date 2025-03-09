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
	Passage       string             `json:"passage"`
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
	Body                  string        `json:"body,omitempty"`
	ImageUrls             []string      `json:"imageUrls,omitempty"`
	Title                 string        `json:"title,omitempty"`
	SelectCount           int32         `json:"selectCount,omitempty"`
	Choices               []ChoiceEntry `json:"choices,omitempty"`
	AudioUrl              string        `json:"audioUrl,omitempty"`
	MaximumQuestionRepeat int32         `json:"maximumQuestionRepeat,omitempty"`
	PreparationDuration   int32         `json:"preparationDuration,omitempty"`
	MaximumAnswerDuration int32         `json:"maximumAnswerDuration,omitempty"`
	TaskCard              string        `json:"taskCard,omitempty"`
}

type ChoiceEntry struct {
	Text string `json:"text"`
}
type TestInfo struct {
	ID     string `json:"testId"`
	Module string `json:"module"`
}
