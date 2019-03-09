package middleware

import (
	"github.com/domsu/stranger-chat/model"
	"testing"
)

func TestProcess_ShouldReturnWaitingMessage_WhenHealthCheck(t *testing.T) {
	ctx := model.Context{PartialRequest: model.PartialRequest{HealthCheck: true}}
	middleware := NewEnsureNotHealthCheck()

	continueProcessing := middleware.Process(&ctx)

	if ctx.PartialResponse.TextResponse != ctx.Messages.WaitingForParticipant ||
		ctx.PartialResponse.MultiMediaResponse != model.WaitingForParticipantSound ||
		ctx.PartialResponse.ExpectUserResponse != true ||
		ctx.UserStorage.OnboardingFinished != true ||
		continueProcessing {
		t.Fail()
	}
}

func TestProcess_ShouldNotAffectResponse_WhenNotHealthCheck(t *testing.T) {
	ctx := model.Context{}
	middleware := NewEnsureNotHealthCheck()

	continueProcessing := middleware.Process(&ctx)

	if ctx.PartialResponse.TextResponse != "" ||
		ctx.PartialResponse.MultiMediaResponse != model.Empty ||
		!continueProcessing {
		t.Fail()
	}
}
