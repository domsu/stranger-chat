package handler

import "github.com/domsu/stranger-chat/model"

//go:generate mockgen -destination mocks/MockStatsRequestHandler.go -package=mocks github.com/domsu/stranger-chat/handler IStatsRequestHandler

type IStatsRequestHandler interface {
	HandleRequest() model.StatsResponse
}
