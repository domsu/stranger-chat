package handler

import "github.com/domsu/stranger-chat/model/external/googleactions"

//go:generate mockgen -destination mocks/MockAppRequestHandler.go -package=mocks github.com/domsu/stranger-chat/handler IAppRequestHandler

type IAppRequestHandler interface {
	HandleRequest(request googleactions.AppRequest) googleactions.Response
}
