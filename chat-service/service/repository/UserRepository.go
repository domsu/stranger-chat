package repository

import (
	. "github.com/domsu/stranger-chat/model"
	"github.com/google/uuid"
	"time"
)

type UserRepository struct {
	users  []*User
	config Config
}

func NewUserRepository(config Config) *UserRepository {
	return &UserRepository{config: config}
}

func (repo *UserRepository) CreateNewUser() *User {
	user := User{Id: uuid.New(), LastActive: time.Now(), BlockedUsersIds: []uuid.UUID{}, LastKnownConversationId: nil}
	repo.users = append(repo.users, &user)
	return &user
}

func (repo *UserRepository) GetUser(userId uuid.UUID) *User {
	for i := range repo.users {
		if repo.users[i].Id == userId {
			return repo.users[i]
		}
	}

	return nil
}

func (repo *UserRepository) GetUsers() []*User {
	return repo.users
}

func (repo *UserRepository) GetActiveUsers() []*User {
	var activeUsers []*User
	var activeDuration = time.Duration(repo.config.UserActiveWhenLastSeenMinutesAgo) * time.Minute
	for i, v := range repo.users {
		if v.LastActive.Add(activeDuration).After(time.Now()) {
			activeUsers = append(activeUsers, repo.users[i])
		}
	}
	return activeUsers
}
