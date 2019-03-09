package middleware

import (
	. "github.com/domsu/stranger-chat/model"
)

//go:generate mockgen -destination mocks/MockMiddleware.go -package=mocks github.com/domsu/stranger-chat/middleware IMiddleware

type IMiddleware interface {
	Process(ctx *Context) bool
}
