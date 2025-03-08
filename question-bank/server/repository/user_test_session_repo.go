package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserTestSessionRepository struct {
	userTestSession *mongo.Collection
}

type UserTestSessionRepoDependencies struct {
	UserTestSession *mongo.Collection
}

func NewUserTestSessionRepository(deps UserTestSessionRepoDependencies) *UserTestSessionRepository {
	return &UserTestSessionRepository{
		userTestSession: deps.UserTestSession,
	}
}
