package handler

import (
	"github.com/domsu/stranger-chat/middleware"
	mocks2 "github.com/domsu/stranger-chat/middleware/mocks"
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/model/external/googleactions"
	"github.com/domsu/stranger-chat/model/transformer/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestHandleRequest_ShouldStopProcessing_WhenMiddlewareRequestReturnsFalse(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRequestTransformer := mocks.NewMockIPartialRequestTransformer(ctrl)
	mockResponseTransformer := mocks.NewMockIPartialResponseTransformer(ctrl)
	mockMiddleware1 := mocks2.NewMockIMiddleware(ctrl)
	mockMiddleware2 := mocks2.NewMockIMiddleware(ctrl)
	mockMiddleware3 := mocks2.NewMockIMiddleware(ctrl)

	mockRequestTransformer.EXPECT().Transform(gomock.Any()).Return(model.PartialRequest{})
	mockResponseTransformer.EXPECT().Transform(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(googleactions.Response{})
	mockMiddleware1.EXPECT().Process(gomock.Any()).Return(true)
	mockMiddleware2.EXPECT().Process(gomock.Any()).Return(false)
	mockMiddleware3.EXPECT().Process(gomock.Any()).Times(0)

	handler := NewAppRequestHandler(
		model.Config{}, make(map[string]model.Messages),
		mockRequestTransformer, mockResponseTransformer,
		[]middleware.IMiddleware{mockMiddleware1, mockMiddleware2, mockMiddleware3},
	)

	handler.HandleRequest(googleactions.AppRequest{})
}

func TestGetMessagesForLocale_ShouldReturnCorrectMessages(t *testing.T) {
	handler := AppRequestHandler{messages: map[string]model.Messages{"pl": {Welcome: "Siema"}, "en": {Welcome: "Hi"}}}

	if !reflect.DeepEqual(handler.getMessagesForLocale("en"), model.Messages{Welcome: "Hi"}) {
		t.Fail()
	}

	if !reflect.DeepEqual(handler.getMessagesForLocale("fr"), model.Messages{Welcome: "Hi"}) {
		t.Fail()
	}
}
