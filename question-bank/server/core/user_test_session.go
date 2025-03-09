package core

import "time"

type TestSessionStatus string

const (
	TestSessionStatusNotStarted TestSessionStatus = "NOT_STARTED"
	TestSessionStatusInProgress TestSessionStatus = "IN_PROGRESS"
	TestSessionStatusCompleted  TestSessionStatus = "COMPLETED"
)

type UserTestInfo struct {
	ID     string            `json:"testId"`
	Module string            `json:"module"`
	Status TestSessionStatus `json:"status"`
}

type UserTestSession struct {
	Test         *TestEntry        `json:"test" bson:"test"`
	Answers      map[string]string `json:"answers" bson:"answers"`
	StartedAt    time.Time         `json:"startedAt" bson:"startedAt"`
	LastActiveAt time.Time         `json:"lastActiveAt" bson:"lastActiveAt"`
	ElapsedTime  int               `json:"elapsedTime" bson:"elapsedTime"`
}
