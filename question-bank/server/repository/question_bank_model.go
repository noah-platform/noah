package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type TestDocument struct {
	ID          primitive.ObjectID `bson:"_id"`
	Scenario    string             `bson:"scenario"`
	AudioScript string             `bson:"audioscript"`
	QuestionSet []interface{}      `bson:"questionset"`
	Audio       string             `bson:"audio"`
	Module      string             `bson:"module"`
}
