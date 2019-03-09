package transformer

import (
	. "encoding/json"
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
	"reflect"
	"testing"
)

func TestTransform_ShouldReturnPartialRequest_WhenVoiceInput(t *testing.T) {
	transformer := PartialRequestTransformer{}
	partialRequest := transformer.Transform(getTestAppRequest(jsonAppRequestVoiceInput))

	userStorage := `{"userId":"b1e4b509-6298-4f1d-b7ff-42540bfb82d5","onboardingFinished":true}`
	expected := model.PartialRequest{
		Type:        model.Text,
		Text:        "Meet people online",
		Storage:     &userStorage,
		InputType:   model.Voice,
		HealthCheck: false,
	}

	if !reflect.DeepEqual(expected, partialRequest) {
		t.Fail()
	}
}

func TestTransform_ShouldReturnPartialRequest_WhenKeyboardInput(t *testing.T) {
	transformer := PartialRequestTransformer{}
	partialRequest := transformer.Transform(getTestAppRequest(jsonAppRequestKeyboardInput))

	userStorage := `{"userId":"b1e4b509-6298-4f1d-b7ff-42540bfb82d5","onboardingFinished":true}`
	expected := model.PartialRequest{
		Type:        model.Text,
		Text:        "Meet people online",
		Storage:     &userStorage,
		InputType:   model.Keyboard,
		HealthCheck: false,
	}

	if !reflect.DeepEqual(expected, partialRequest) {
		t.Fail()
	}
}

func TestTransform_ShouldReturnPartialRequest_WhenHealthCheck(t *testing.T) {
	transformer := PartialRequestTransformer{}
	partialRequest := transformer.Transform(getTestAppRequest(jsonAppRequestHealthCheck))

	userStorage := `{"userId":"b1e4b509-6298-4f1d-b7ff-42540bfb82d5","onboardingFinished":true}`
	expected := model.PartialRequest{
		Type:        model.Text,
		Text:        "Meet people online",
		Storage:     &userStorage,
		InputType:   model.Keyboard,
		HealthCheck: true,
	}

	if !reflect.DeepEqual(expected, partialRequest) {
		t.Fail()
	}
}

func getTestAppRequest(json string) googleactions.AppRequest {
	var request googleactions.AppRequest
	err := Unmarshal([]byte(json), &request)
	if err != nil {
		panic(err)
	}
	return request
}

const jsonAppRequestVoiceInput = `{
   "user":{
      "userId":"ABwppHEG9HgmIqUzvDnE2ahsWKKAPtUz-S86b43AxvIM5uxxjbiW0J9wkfV_QZPlFIgKRR1FvczXXLarodY",
      "locale":"pl-PL",
      "lastSeen":"2019-02-05T09:44:46Z",
      "userStorage":"{\"userId\":\"b1e4b509-6298-4f1d-b7ff-42540bfb82d5\",\"onboardingFinished\":true}"
   },
   "conversation":{
      "conversationId":"ABwppHE3FfG7rrf-iM1U5VBzEIZisgV_9C5UJVFFx3315Oi6CQjStYOuhg8Nqy3scHivw55YS3oJWzXm_vg",
      "type":"NEW"
   },
   "inputs":[
      {
         "intent":"actions.intent.TEXT",
         "rawInputs":[
            {
               "inputType":"VOICE",
               "query":"Meet people online"
            }
         ],
         "arguments":[
            {
               "name":"text",
               "rawText":"Meet people online",
               "textValue":"Meet people online"
            }
         ]
      }
   ],
   "surface":{
      "capabilities":[
         {
            "name":"actions.capability.MEDIA_RESPONSE_AUDIO"
         },
         {
            "name":"actions.capability.WEB_BROWSER"
         },
         {
            "name":"actions.capability.AUDIO_OUTPUT"
         },
         {
            "name":"actions.capability.SCREEN_OUTPUT"
         }
      ]
   },
   "isInSandbox":true,
   "availableSurfaces":[
      {
         "capabilities":[
            {
               "name":"actions.capability.WEB_BROWSER"
            },
            {
               "name":"actions.capability.AUDIO_OUTPUT"
            },
            {
               "name":"actions.capability.SCREEN_OUTPUT"
            }
         ]
      }
   ]
}`

const jsonAppRequestKeyboardInput = `{
   "user":{
      "userId":"ABwppHEG9HgmIqUzvDnE2ahsWKKAPtUz-S86b43AxvIM5uxxjbiW0J9wkfV_QZPlFIgKRR1FvczXXLarodY",
      "locale":"pl-PL",
      "lastSeen":"2019-02-05T09:44:46Z",
      "userStorage":"{\"userId\":\"b1e4b509-6298-4f1d-b7ff-42540bfb82d5\",\"onboardingFinished\":true}"
   },
   "conversation":{
      "conversationId":"ABwppHE3FfG7rrf-iM1U5VBzEIZisgV_9C5UJVFFx3315Oi6CQjStYOuhg8Nqy3scHivw55YS3oJWzXm_vg",
      "type":"NEW"
   },
   "inputs":[
      {
         "intent":"actions.intent.TEXT",
         "rawInputs":[
            {
               "inputType":"KEYBOARD",
               "query":"Meet people online"
            }
         ],
         "arguments":[
            {
               "name":"text",
               "rawText":"Meet people online",
               "textValue":"Meet people online"
            }
         ]
      }
   ],
   "surface":{
      "capabilities":[
         {
            "name":"actions.capability.MEDIA_RESPONSE_AUDIO"
         },
         {
            "name":"actions.capability.WEB_BROWSER"
         },
         {
            "name":"actions.capability.AUDIO_OUTPUT"
         },
         {
            "name":"actions.capability.SCREEN_OUTPUT"
         }
      ]
   },
   "isInSandbox":true,
   "availableSurfaces":[
      {
         "capabilities":[
            {
               "name":"actions.capability.WEB_BROWSER"
            },
            {
               "name":"actions.capability.AUDIO_OUTPUT"
            },
            {
               "name":"actions.capability.SCREEN_OUTPUT"
            }
         ]
      }
   ]
}`

const jsonAppRequestHealthCheck = `{
   "user":{
      "userId":"ABwppHEG9HgmIqUzvDnE2ahsWKKAPtUz-S86b43AxvIM5uxxjbiW0J9wkfV_QZPlFIgKRR1FvczXXLarodY",
      "locale":"pl-PL",
      "lastSeen":"2019-02-05T09:44:46Z",
      "userStorage":"{\"userId\":\"b1e4b509-6298-4f1d-b7ff-42540bfb82d5\",\"onboardingFinished\":true}"
   },
   "conversation":{
      "conversationId":"ABwppHE3FfG7rrf-iM1U5VBzEIZisgV_9C5UJVFFx3315Oi6CQjStYOuhg8Nqy3scHivw55YS3oJWzXm_vg",
      "type":"NEW"
   },
   "inputs":[
      {
         "intent":"actions.intent.TEXT",
         "rawInputs":[
            {
               "inputType":"KEYBOARD",
               "query":"Meet people online"
            }
         ],
         "arguments":[
            {
               "name":"text",
               "rawText":"Meet people online",
               "textValue":"Meet people online"
            },
			{
				"name":"is_health_check"
			}
         ]
      }
   ],
   "surface":{
      "capabilities":[
         {
            "name":"actions.capability.MEDIA_RESPONSE_AUDIO"
         },
         {
            "name":"actions.capability.WEB_BROWSER"
         },
         {
            "name":"actions.capability.AUDIO_OUTPUT"
         },
         {
            "name":"actions.capability.SCREEN_OUTPUT"
         }
      ]
   },
   "isInSandbox":true,
   "availableSurfaces":[
      {
         "capabilities":[
            {
               "name":"actions.capability.WEB_BROWSER"
            },
            {
               "name":"actions.capability.AUDIO_OUTPUT"
            },
            {
               "name":"actions.capability.SCREEN_OUTPUT"
            }
         ]
      }
   ]
}`
