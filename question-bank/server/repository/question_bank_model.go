package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestDocument struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Duration    int32              `bson:"duration"`
	Instruction string             `bson:"instruction"`
	Module      string             `bson:"module"`
	Sections    []SectionDocument  `bson:"sections"`
}

type SectionDocument struct {
	ID            primitive.ObjectID    `bson:"_id"`
	Title         string                `bson:"title"`
	Instruction   string                `bson:"instruction"`
	AudioUrl      string                `bson:"audioUrl"`
	QuestionCount int32                 `bson:"questionCount"`
	QuestionSet   []QuestionSetDocument `bson:"questionset"`
}

type QuestionSetDocument struct {
	Instruction  string             `bson:"instruction"`
	ResponseType string             `bson:"responseType"`
	Questions    []QuestionDocument `bson:"questions"`
}

type QuestionDocument struct {
	Body                  string           `bson:"body,omitempty"`
	ImageUrls             []string         `bson:"imageUrls,omitempty"`
	Title                 string           `bson:"title,omitempty"`
	SelectCount           int32            `bson:"selectCount,omitempty"`
	Choices               []ChoiceDocument `bson:"choices,omitempty"`
	AudioUrl              string           `bson:"audioUrl,omitempty"`
	MaximumQuestionRepeat int32            `bson:"maximumQuestionRepeat,omitempty"`
	PreparationDuration   int32            `bson:"preparationDuration,omitempty"`
	MaximumAnswerDuration int32            `bson:"maximumAnswerDuration,omitempty"`
	TaskCard              string           `bson:"taskCard,omitempty"`
}

type ChoiceDocument struct {
	Text string `bson:"text"`
}
