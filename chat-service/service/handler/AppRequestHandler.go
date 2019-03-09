package handler

import (
	actions "github.com/domsu/stranger-chat/model/external/googleactions"
	"strings"
)
import "github.com/domsu/stranger-chat/middleware"
import "github.com/domsu/stranger-chat/model"
import "github.com/domsu/stranger-chat/model/transformer"

type AppRequestHandler struct {
	config                     model.Config
	messages                   map[string]model.Messages
	middleware                 []middleware.IMiddleware
	partialRequestTransformer  transformer.IPartialRequestTransformer
	partialResponseTransformer transformer.IPartialResponseTransformer
}

func NewAppRequestHandler(
	config model.Config,
	messages map[string]model.Messages,
	partialRequestTransformer transformer.IPartialRequestTransformer,
	partialResponseTransformer transformer.IPartialResponseTransformer,
	middleware []middleware.IMiddleware) AppRequestHandler {

	return AppRequestHandler{
		config:                     config,
		messages:                   messages,
		middleware:                 middleware,
		partialRequestTransformer:  partialRequestTransformer,
		partialResponseTransformer: partialResponseTransformer,
	}
}

func (handler AppRequestHandler) HandleRequest(request actions.AppRequest) actions.Response {
	context := model.Context{}
	context.PartialRequest = handler.partialRequestTransformer.Transform(request)
	context.Config = handler.config
	context.Messages = handler.getMessagesForLocale(request.User.Locale)

	for _, v := range handler.middleware {
		if v.Process(&context) == false {
			break
		}
	}

	return handler.partialResponseTransformer.Transform(context.PartialResponse, context.UserStorage, context.Messages, context.Config)
}

func (handler AppRequestHandler) getMessagesForLocale(locale string) model.Messages {
	countryRegion := strings.Split(locale, "-")
	messages, ok := handler.messages[countryRegion[0]]
	if ok {
		return messages
	} else {
		return handler.messages["en"]
	}
}
