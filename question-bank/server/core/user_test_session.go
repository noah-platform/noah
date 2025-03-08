package core

import "time"

type UserTestSession struct {
	Test         *TestEntry        `json:"test" bson:"test"`
	Answers      map[string]string `json:"answers" bson:"answers"`
	StartedAt    time.Time         `json:"startedAt" bson:"startedAt"`
	LastActiveAt time.Time         `json:"lastActiveAt" bson:"lastActiveAt"`
	ElapsedTime  int               `json:"elapsedTime" bson:"elapsedTime"`
}
