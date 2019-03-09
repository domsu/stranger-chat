package middleware

import (
	"github.com/domsu/stranger-chat/model"
	"testing"
)

func TestProcess_ShouldSetupUserStorage_WhenStorageDataInRequest(t *testing.T) {
	storageString := `{"userId": "14e8ef54-d0f8-4ccd-9bc6-f51fecb78e2d", "onboardingFinished": true}`
	ctx := model.Context{User: &model.User{}, PartialRequest: model.PartialRequest{Storage: &storageString}}
	middleware := NewSetUserStorage()

	continueProcessing := middleware.Process(&ctx)

	if ctx.UserStorage.UserId.String() != "14e8ef54-d0f8-4ccd-9bc6-f51fecb78e2d" ||
		ctx.UserStorage.OnboardingFinished != true ||
		!continueProcessing {
		t.Fail()
	}
}

func TestProcess_ShouldNotPanic_WhenStorageNotDataInRequest(t *testing.T) {
	ctx := model.Context{User: &model.User{}, PartialRequest: model.PartialRequest{}}
	middleware := NewSetUserStorage()

	continueProcessing := middleware.Process(&ctx)

	if !continueProcessing {
		t.Fail()
	}
}

func TestProcess_ShouldPanic_WhenStorageDataInRequestWithInvalidFormat(t *testing.T) {
	storageString := `\\`
	ctx := model.Context{User: &model.User{}, PartialRequest: model.PartialRequest{Storage: &storageString}}
	middleware := NewSetUserStorage()

	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	middleware.Process(&ctx)
}
