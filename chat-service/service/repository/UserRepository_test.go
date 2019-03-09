package repository

import (
	. "github.com/domsu/stranger-chat/model"
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)

func TestNewUserRepository_ShouldCreateNewRepository(t *testing.T) {
	repository := NewUserRepository(Config{UserActiveWhenLastSeenMinutesAgo: 1})

	if repository.config.UserActiveWhenLastSeenMinutesAgo != 1 {
		t.Error("Wrong appconfig")
	}
	if repository.users != nil {
		t.Error("Wrong conversations")
	}
}

func TestCreateNewUser_ShouldCreateNewUserAndToRepository(t *testing.T) {
	repository := NewUserRepository(Config{})

	user := repository.CreateNewUser()

	if !reflect.DeepEqual(
		*user,
		User{Id: user.Id, LastActive: user.LastActive, BlockedUsersIds: []uuid.UUID{}, LastKnownConversationId: nil}) {
		t.Error("Wrong user")
	}
	if len(repository.GetUsers()) != 1 {
		t.Error("User not added to repository")
	}
}

func TestGetUser_ShouldReturnCorrectUser(t *testing.T) {
	repository := NewUserRepository(Config{})
	user1 := repository.CreateNewUser()
	repository.CreateNewUser()

	if repository.GetUser(user1.Id).Id != user1.Id {
		t.Error("Wrong user")
	}
}

func TestGetUser_ShouldReturnNilWhenNoUser(t *testing.T) {
	repository := NewUserRepository(Config{})
	repository.CreateNewUser()

	if repository.GetUser(uuid.New()) != nil {
		t.Error("Wrong result")
	}
}

func TestGetUsers_ShouldReturnUsers(t *testing.T) {
	repository := NewUserRepository(Config{})
	user1 = *repository.CreateNewUser()
	user2 = *repository.CreateNewUser()

	if !reflect.DeepEqual(
		repository.GetUsers(),
		[]*User{&user1, &user2}) {
		t.Error("Wrong users")
	}
}

func TestGetActiveUsers_ShouldReturnActiveUsers(t *testing.T) {
	repository := NewUserRepository(Config{UserActiveWhenLastSeenMinutesAgo: 1})
	user1 := repository.CreateNewUser()
	user2 := repository.CreateNewUser()

	user1.LastActive = time.Now().Add(time.Duration(-2000000) * time.Minute)

	activeUsers := repository.GetActiveUsers()
	if len(activeUsers) != 1 || activeUsers[0].Id != user2.Id {
		t.Error("Wrong users")
	}
}
