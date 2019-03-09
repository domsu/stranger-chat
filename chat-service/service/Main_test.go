package main

import (
	"github.com/domsu/stranger-chat/handler/mocks"
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet_ShouldUseAppRequestHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockStatsRequestHandler := mocks.NewMockIStatsRequestHandler(mockCtrl)
	mockAppRequestHandler := mocks.NewMockIAppRequestHandler(mockCtrl)
	mockAppRequestHandler.EXPECT().HandleRequest(gomock.Any()).Return(googleactions.Response{ConversationToken: "token"})

	handler := getHttpHandler(mockStatsRequestHandler, mockAppRequestHandler)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	res, err := http.Post(testServer.URL, "application/json", strings.NewReader(`{"user":{"userId":"ABwppHEQcl1aKVhDONeol1P3ObNQWI5m4V4CnaJn5Os9rDdWDOez85Kxy1iSdHpQ54Y6YXtHxmxtUmZ9Ssc","locale":"en-US","lastSeen":"2019-02-09T22:27:51Z","userStorage":"{\"userId\":\"aa9714f7-01f9-4f6b-97b9-749f1f5ff132\",\"onboardingFinished\":true}"},"conversation":{"conversationId":"ABwppHESZRr8R2ri5S8n4NWfNAK2xwlCVxKaaDkQKL20YpkL35tjILrrCtZWQ_3AR9cBWxdWGoPkJB8WPTA","type":"NEW"},"inputs":[{"intent":"actions.intent.TEXT","rawInputs":[{"inputType":"VOICE","query":"Check status"}],"arguments":[{"name":"text","rawText":"Check status","textValue":"Check status"}]}],"surface":{"capabilities":[{"name":"actions.capability.SCREEN_OUTPUT"},{"name":"actions.capability.MEDIA_RESPONSE_AUDIO"},{"name":"actions.capability.WEB_BROWSER"},{"name":"actions.capability.AUDIO_OUTPUT"}]},"isInSandbox":true,"requestType":"SIMULATOR"}`))
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	println(string(data))
	if strings.Trim(string(data), "\n") != `{"conversationToken":"token","expectUserResponse":false,"expectedInputs":null}` {
		t.Fail()
	}
}

func TestGetStatsRequest_ShouldUseStatsRequestHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockStatsRequestHandler := mocks.NewMockIStatsRequestHandler(mockCtrl)
	mockAppRequestHandler := mocks.NewMockIAppRequestHandler(mockCtrl)
	mockStatsRequestHandler.EXPECT().HandleRequest().Return(model.StatsResponse{UsersCount: 1, ActiveUsersCount: 2, ConversationsCount: 3})

	handler := getHttpHandler(mockStatsRequestHandler, mockAppRequestHandler)
	testServer := httptest.NewServer(handler)
	defer testServer.Close()

	res, err := http.Get(testServer.URL + "/stats")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Trim(string(data), "\n") != `{"usersCount":1,"activeUsersCount":2,"conversationsCount":3}` {
		t.Fail()
	}
}

func TestParseAppRequest_ShouldPanic_WhenInvalidFormat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	request, e := http.NewRequest("POST", "localhost", strings.NewReader("invalid format"))
	if e != nil {
		log.Fatal(e)
	}

	parseAppRequest(request)
}

func TestParseConfig_ShouldParseConfig(t *testing.T) {
	config := parseConfig("userActiveWhenLastSeenMinutesAgo: 1")

	if config.UserActiveWhenLastSeenMinutesAgo != 1 {
		t.Fail()
	}
}

func TestParseConfig_ShouldPanic_WhenInvalidFormat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	parseConfig("invalid format")
}

func TestGetMessages_ShouldReturnArrayWithEnglishAndPolishMessages(t *testing.T) {
	messages := getMessages()

	if _, ok := messages["en"]; !ok {
		t.Fail()
	}

	if _, ok := messages["pl"]; !ok {
		t.Fail()
	}
}

func TestParseMessages_ShouldPanic_WhenInvalidFormat(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()

	parseMessages("invalid format")
}
