package handler

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/domsu/stranger-chat/repository/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestHandleRequest_ShouldHandleRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepository := mocks.NewMockIUserRepository(ctrl)
	mockConversationRepository := mocks.NewMockIConversationRepository(ctrl)
	mockUserRepository.EXPECT().GetUsers().Return([]*model.User{{}, {}})
	mockUserRepository.EXPECT().GetActiveUsers().Return([]*model.User{{}})
	mockConversationRepository.EXPECT().GetConversations().Return([]*model.Conversation{{}})

	handler := NewStatsRequestHandler(mockConversationRepository, mockUserRepository)
	response := handler.HandleRequest()

	if response.UsersCount != 2 ||
		response.ActiveUsersCount != 1 ||
		response.ConversationsCount != 1 {
		t.Fail()
	}
}
