package middleware

import (
	"github.com/domsu/stranger-chat/model"
	"testing"
)

func TestProcess_ShouldSetLastActiveTime(t *testing.T) {
	ctx := model.Context{User: &model.User{}}
	middleware := NewSetupUserLastActive()

	continueProcessing := middleware.Process(&ctx)

	if ctx.User.LastActive.IsZero() ||
		!continueProcessing {
		t.Fail()
	}
}
