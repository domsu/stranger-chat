package handler

import (
	"github.com/domsu/stranger-chat/model"
	. "github.com/domsu/stranger-chat/repository"
)

type StatsRequestHandler struct {
	conversationRepository IConversationRepository
	userRepository         IUserRepository
}

func NewStatsRequestHandler(conversationRepository IConversationRepository, userRepository IUserRepository) StatsRequestHandler {
	return StatsRequestHandler{conversationRepository: conversationRepository, userRepository: userRepository}
}

func (handler StatsRequestHandler) HandleRequest() model.StatsResponse {
	return model.StatsResponse{
		UsersCount:         len(handler.userRepository.GetUsers()),
		ActiveUsersCount:   len(handler.userRepository.GetActiveUsers()),
		ConversationsCount: len(handler.conversationRepository.GetConversations()),
	}
}
