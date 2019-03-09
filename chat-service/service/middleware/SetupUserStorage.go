package middleware

import (
	"encoding/json"
	. "github.com/domsu/stranger-chat/model"
)

type SetupUserStorage struct {
}

func NewSetUserStorage() SetupUserStorage {
	return SetupUserStorage{}
}

func (action SetupUserStorage) Process(ctx *Context) bool {
	userStorage := UserStorage{}

	if ctx.PartialRequest.Storage != nil && len(*ctx.PartialRequest.Storage) != 0 {
		if err := json.Unmarshal([]byte(*ctx.PartialRequest.Storage), &userStorage); err != nil {
			panic("Invalid user storage format in request")
		}
	}
	ctx.UserStorage = userStorage

	return true
}
