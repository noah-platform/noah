package core

type UserTestSession struct {
	Test    *TestEntry        `json:"test"`
	Answers map[string]string `json:"answers"`
}
