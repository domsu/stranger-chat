package transformer

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
)

//go:generate mockgen -destination mocks/MockPartialRequestTransformer.go -package=mocks github.com/domsu/stranger-chat/model/transformer IPartialRequestTransformer

type IPartialRequestTransformer interface {
	Transform(request googleactions.AppRequest) model.PartialRequest
}
