package main

import (
	"encoding/json"
	"github.com/domsu/stranger-chat/appconfig"
	. "github.com/domsu/stranger-chat/handler"
	"github.com/domsu/stranger-chat/middleware"
	"github.com/domsu/stranger-chat/model"
	actions "github.com/domsu/stranger-chat/model/external/googleactions"
	"github.com/domsu/stranger-chat/model/transformer"
	. "github.com/domsu/stranger-chat/repository"
	"github.com/mholt/certmagic"
	"gopkg.in/yaml.v2"
	"net/http"
)

func main() {
	config := parseConfig(appconfig.Config)

	conversationRepository := NewConversationRepository(config)
	userRepository := NewUserRepository(config)

	appRequestMiddleware := []middleware.IMiddleware{
		middleware.NewSetUserStorage(),
		middleware.NewEnsureNotHealthCheck(),
		middleware.NewSetUser(userRepository),
		middleware.NewSetResponse(userRepository, conversationRepository),
		middleware.NewSetupUserLastActive(),
	}

	appRequestHandler := NewAppRequestHandler(
		config, getMessages(),
		&transformer.PartialRequestTransformer{}, &transformer.PartialResponseTransformer{},
		appRequestMiddleware,
	)
	statsRequestHandler := NewStatsRequestHandler(conversationRepository, userRepository)

	mux := getHttpHandler(&statsRequestHandler, &appRequestHandler)

	if err := certmagic.HTTPS([]string{config.ServiceDomain, "www." + config.ServiceDomain}, mux); err != nil {
		panic(err)
	}
}

func getHttpHandler(statsRequestHandler IStatsRequestHandler, appRequestHandler IAppRequestHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		response := statsRequestHandler.HandleRequest()
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		appRequest := parseAppRequest(r)
		response := appRequestHandler.HandleRequest(appRequest)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			panic(err)
		}
	})
	return mux
}

func parseAppRequest(r *http.Request) actions.AppRequest {
	decoder := json.NewDecoder(r.Body)
	var request actions.AppRequest
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}
	return request
}

func parseConfig(data string) model.Config {
	config := model.Config{}
	err := yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		panic("Unable to parse config")
	}

	return config
}

func getMessages() map[string]model.Messages {
	messages := map[string]model.Messages{
		"en": parseMessages(appconfig.MessagesEn),
		"pl": parseMessages(appconfig.MessagesPl),
	}
	return messages
}

func parseMessages(data string) model.Messages {
	messages := model.Messages{}
	err := yaml.Unmarshal([]byte(data), &messages)
	if err != nil {
		panic("Unable to parse messages")
	}

	return messages
}
