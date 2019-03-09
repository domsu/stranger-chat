package repository

import (
	"github.com/domsu/stranger-chat/model"
	"github.com/google/uuid"
)

//go:generate mockgen -destination mocks/MockConversationRepository.go -package=mocks github.com/domsu/stranger-chat/repository IConversationRepository

type IConversationRepository interface {
	GetConversationForUser(userId uuid.UUID) *model.Conversation
	RemoveConversation(conversationId uuid.UUID)
	GetConversations() []*model.Conversation
	CreateConversation(initiator uuid.UUID, participant uuid.UUID) *model.Conversation
}
