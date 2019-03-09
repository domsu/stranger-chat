package transformer

import (
	"encoding/json"
	"github.com/domsu/stranger-chat/model"
	. "github.com/domsu/stranger-chat/model/external/googleactions"
)

type PartialResponseTransformer struct {
	partialResponse model.PartialResponse
	userStorage     model.UserStorage
	messages        model.Messages
	config          model.Config
}

func (transformer PartialResponseTransformer) Transform(partialResponse model.PartialResponse, userStorage model.UserStorage, messages model.Messages, config model.Config) Response {
	transformer.partialResponse = partialResponse
	transformer.userStorage = userStorage
	transformer.messages = messages
	transformer.config = config

	response := Response{}
	response.ConversationToken = ""
	response.UserStorage = transformer.getUserStorageString()
	response.ExpectUserResponse = partialResponse.ExpectUserResponse

	if len(partialResponse.FinalResponse) != 0 {
		response.FinalResponse = &FinalResponse{RichResponse: transformer.getRichResponse(partialResponse.FinalResponse)}
	} else {
		response.ExpectedInputs = []ExpectedInput{transformer.getExpectedInput()}
	}

	return response
}

func (transformer PartialResponseTransformer) getExpectedInput() ExpectedInput {
	RichResponse := transformer.getRichResponse(transformer.partialResponse.TextResponse)

	input := ExpectedInput{
		InputPrompt: InputPrompt{
			RichInitialPrompt: RichResponse,
		},
		PossibleIntents: []ExpectedIntent{
			{"actions.intent.TEXT"},
		},
		SpeechBiasingHints: nil,
	}
	return input
}

func (transformer PartialResponseTransformer) getRichResponse(textResponse string) RichResponse {
	var suggestion *[]Suggestion = nil
	items := []Item{
		{
			SimpleResponse: &SimpleResponse{
				TextToSpeech: "",
				Ssml:         "<speak>" + textResponse + "</speak>",
				DisplayText:  "",
			},
			MediaResponse: nil,
		},
	}
	if transformer.partialResponse.MultiMediaResponse == model.WaitingForParticipantSound {
		items = append(items, Item{
			SimpleResponse: nil,
			MediaResponse: &MediaResponse{
				MediaType: "AUDIO",
				MediaObjects: []MediaObject{
					{
						Name:        transformer.messages.MusicTitle,
						Description: transformer.messages.MusicDescription,
						ContentUrl:  transformer.config.MusicOnHoldUrl,
					},
				},
			},
		})

		suggestion = &[]Suggestion{
			{transformer.messages.CheckStatus},
		}
	}
	RichResponse := RichResponse{
		Items:       items,
		Suggestions: suggestion,
	}
	return RichResponse
}

func (transformer PartialResponseTransformer) getUserStorageString() string {
	userStorageJson, err := json.Marshal(transformer.userStorage)

	if err != nil {
		panic("Error while parsing user response")
	}

	return string(userStorageJson)
}
