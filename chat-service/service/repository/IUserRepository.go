package repository

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/google/uuid"
)

//go:generate mockgen -destination mocks/MockUserRepository.go -package=mocks github.com/domsu/stranger-chat/repository IUserRepository

type IUserRepository interface {
	CreateNewUser() *model.User
	GetUser(userId uuid.UUID) *model.User
	GetUsers() []*model.User
	GetActiveUsers() []*model.User
}
