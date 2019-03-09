package middleware

import "github.com/domsu/stranger-chat/model"

type EnsureNotHealthCheck struct{}

func NewEnsureNotHealthCheck() EnsureNotHealthCheck {
	return EnsureNotHealthCheck{}
}

func (EnsureNotHealthCheck) Process(ctx *model.Context) bool {
	if ctx.PartialRequest.HealthCheck {
		ctx.PartialResponse.TextResponse = ctx.Messages.WaitingForParticipant
		ctx.PartialResponse.MultiMediaResponse = model.WaitingForParticipantSound
		ctx.PartialResponse.ExpectUserResponse = true
		ctx.UserStorage.OnboardingFinished = true
		return false
	} else {
		return true
	}
}
