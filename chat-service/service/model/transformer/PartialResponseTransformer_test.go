package transformer

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestTransform_ShouldTransformNonFinalPartialResponse(t *testing.T) {
	partialResponse := model.PartialResponse{TextResponse: "text response", MultiMediaResponse: model.WaitingForParticipantSound, ExpectUserResponse: true}
	userId := uuid.MustParse("39bec698-d6c3-4b57-89db-077156df261f")
	userStorage := model.UserStorage{UserId: &userId}
	messages := model.Messages{MusicTitle: "music title", MusicDescription: "music desc", CheckStatus: "check"}
	config := model.Config{MusicOnHoldUrl: "http://www.google.com/music.mp3"}

	response := PartialResponseTransformer{}.Transform(partialResponse, userStorage, messages, config)
	expected := googleactions.Response{
		ConversationToken:  "",
		UserStorage:        `{"userId":"39bec698-d6c3-4b57-89db-077156df261f","onboardingFinished":false}`,
		ExpectUserResponse: true,
		ExpectedInputs: []googleactions.ExpectedInput{{
			InputPrompt: googleactions.InputPrompt{
				RichInitialPrompt: googleactions.RichResponse{
					Items: []googleactions.Item{
						{
							SimpleResponse: &googleactions.SimpleResponse{
								TextToSpeech: "",
								Ssml:         "<speak>text response</speak>",
								DisplayText:  "",
							},
							MediaResponse: nil,
						},
						{
							SimpleResponse: nil,
							MediaResponse: &googleactions.MediaResponse{
								MediaType: "AUDIO",
								MediaObjects: []googleactions.MediaObject{
									{
										Name:        "music title",
										Description: "music desc",
										ContentUrl:  "http://www.google.com/music.mp3",
									},
								},
							},
						},
					},
					Suggestions: &[]googleactions.Suggestion{
						{"check"},
					},
				},
			},
			PossibleIntents: []googleactions.ExpectedIntent{
				{"actions.intent.TEXT"},
			},
			SpeechBiasingHints: nil,
		}},
		FinalResponse: nil,
	}

	if !reflect.DeepEqual(expected, response) {
		t.Fatal()
	}
}

func TestTransform_ShouldTransformFinalPartialResponse(t *testing.T) {
	partialResponse := model.PartialResponse{FinalResponse: "final response", MultiMediaResponse: model.Empty}
	userId := uuid.MustParse("39bec698-d6c3-4b57-89db-077156df261f")
	userStorage := model.UserStorage{UserId: &userId}
	messages := model.Messages{}
	config := model.Config{}

	response := PartialResponseTransformer{}.Transform(partialResponse, userStorage, messages, config)
	expected := googleactions.Response{
		ConversationToken:  "",
		UserStorage:        `{"userId":"39bec698-d6c3-4b57-89db-077156df261f","onboardingFinished":false}`,
		ExpectUserResponse: false,
		FinalResponse: &googleactions.FinalResponse{RichResponse: googleactions.RichResponse{
			Items: []googleactions.Item{
				{
					SimpleResponse: &googleactions.SimpleResponse{
						TextToSpeech: "",
						Ssml:         "<speak>final response</speak>",
						DisplayText:  "",
					},
					MediaResponse: nil,
				},
			},
			Suggestions: nil,
		},
		}}

	if !reflect.DeepEqual(expected, response) {
		t.Fatal()
	}
}
