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
	Tags   []string          `json:"tags"`
	Status TestSessionStatus `json:"status"`
}

type UserTestSession struct {
	UserID       string            `bson:"userId"`
	Test         *TestEntry        `bson:"test"`
	Answers      map[string]string `bson:"answers"`
	StartedAt    time.Time         `bson:"startedAt"`
	LastActiveAt time.Time         `bson:"lastActiveAt"`
	ElapsedTime  int               `bson:"elapsedTime"`
	CompletedAt  *time.Time        `bson:"completedAt"`
}
