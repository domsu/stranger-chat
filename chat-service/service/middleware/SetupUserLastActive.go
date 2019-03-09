package middleware

import (
	. "github.com/domsu/stranger-chat/model"
	"time"
)

type SetupUserLastActive struct {
}

func NewSetupUserLastActive() SetupUserLastActive {
	return SetupUserLastActive{}
}

func (action SetupUserLastActive) Process(ctx *Context) bool {
	ctx.User.LastActive = time.Now()

	return true
}
