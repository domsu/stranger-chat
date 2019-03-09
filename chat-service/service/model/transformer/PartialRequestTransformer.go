package transformer

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
)

type PartialRequestTransformer struct{}

func (transformer PartialRequestTransformer) Transform(request googleactions.AppRequest) model.PartialRequest {
	requestType := model.Action
	if request.Input[0].Intent == "actions.intent.TEXT" {
		requestType = model.Text
	}

	text := request.Input[0].RawInputs[0].Query

	inputType := model.Voice
	if request.Input[0].RawInputs[0].InputType == "KEYBOARD" {
		inputType = model.Keyboard
	}

	healthCheck := false
	for _, input := range request.Input {
		for _, arg := range input.Arguments {
			if arg.Name == "is_health_check" {
				healthCheck = true
			}
		}
	}

	return model.PartialRequest{Type: requestType, Text: text, Storage: request.User.UserStorage, InputType: inputType, HealthCheck: healthCheck}
}
