package transformer

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
)

//go:generate mockgen -destination mocks/MockPartialResponseTransformer.go -package=mocks github.com/domsu/stranger-chat/model/transformer IPartialResponseTransformer

type IPartialResponseTransformer interface {
	Transform(partialResponse model.PartialResponse, userStorage model.UserStorage, messages model.Messages, config model.Config) googleactions.Response
}
