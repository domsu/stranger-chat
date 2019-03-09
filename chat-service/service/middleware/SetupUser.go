package middleware

import (
	. "github.com/domsu/stranger-chat/model"
	. "github.com/domsu/stranger-chat/repository"
)

type SetupUser struct {
	userRepository IUserRepository
}

func NewSetUser(userRepository IUserRepository) SetupUser {
	return SetupUser{userRepository: userRepository}
}

func (action SetupUser) Process(ctx *Context) bool {
	if ctx.UserStorage.UserId == nil || action.userRepository.GetUser(*ctx.UserStorage.UserId) == nil {
		ctx.User = action.userRepository.CreateNewUser()
	} else {
		ctx.User = action.userRepository.GetUser(*ctx.UserStorage.UserId)
	}
	ctx.UserStorage.UserId = &ctx.User.Id

	return true
}
