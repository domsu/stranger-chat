package middleware

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestShouldCreateNewUser_WhenUserIdInStorageIsNull(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	context := model.Context{}
	user := model.User{Id: uuid.New()}
	mockUserRepository := mocks.NewMockIUserRepository(mockCtrl)
	mockUserRepository.EXPECT().
		CreateNewUser().
		Return(&user)
	middleware := NewSetUser(mockUserRepository)

	middleware.Process(&context)

	if !reflect.DeepEqual(user, *context.User) {
		t.Fail()
	}
}

func TestShouldCreateNewUser_WhenUserIdStorageNotInRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	nonExistingUserUuid := uuid.New()
	context := model.Context{UserStorage: model.UserStorage{UserId: &nonExistingUserUuid}}
	user := model.User{Id: uuid.New()}
	mockUserRepository := mocks.NewMockIUserRepository(mockCtrl)
	mockUserRepository.EXPECT().
		GetUser(nonExistingUserUuid).
		Return(nil)
	mockUserRepository.EXPECT().
		CreateNewUser().
		Return(&user)
	middleware := NewSetUser(mockUserRepository)

	middleware.Process(&context)

	if !reflect.DeepEqual(user, *context.User) {
		t.Fail()
	}
}

func TestShouldGetUserFromRepository_WhenUserIdStorageInRepository(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	existingUserId := uuid.New()
	context := model.Context{UserStorage: model.UserStorage{UserId: &existingUserId}}
	user := model.User{Id: existingUserId}
	mockUserRepository := mocks.NewMockIUserRepository(mockCtrl)
	mockUserRepository.EXPECT().
		GetUser(existingUserId).
		Return(&user).
		AnyTimes()
	middleware := NewSetUser(mockUserRepository)

	middleware.Process(&context)

	if !reflect.DeepEqual(user, *context.User) {
		t.Fail()
	}
}
