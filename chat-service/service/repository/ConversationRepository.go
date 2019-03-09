package repository

import (
	. "github.com/domsu/stranger-chat/model"
	"github.com/google/uuid"
)

type ConversationRepository struct {
	conversations []*Conversation
	config        Config
}

func NewConversationRepository(config Config) *ConversationRepository {
	return &ConversationRepository{config: config}
}

func (repo *ConversationRepository) GetConversationForUser(userId uuid.UUID) *Conversation {
	for i, v := range repo.conversations {
		if v.Participants[0] == userId || v.Participants[1] == userId {
			return repo.conversations[i]
		}
	}
	return nil
}

func (repo *ConversationRepository) RemoveConversation(conversationId uuid.UUID) {
	var conversations []*Conversation
	for i := range repo.conversations {
		if repo.conversations[i].Id != conversationId {
			conversations = append(conversations, repo.conversations[i])
		}
	}
	repo.conversations = conversations
}

func (repo *ConversationRepository) GetConversations() []*Conversation {
	return repo.conversations
}

func (repo *ConversationRepository) CreateConversation(initiator uuid.UUID, participant uuid.UUID) *Conversation {
	conversation := Conversation{Id: uuid.New(), Participants: []uuid.UUID{initiator, participant}, ConversationState: UserOneStartsConversation, RecentReply: nil}
	repo.conversations = append(repo.conversations, &conversation)
	return &conversation
}
