package model

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Id                uuid.UUID
	Participants      []uuid.UUID
	ConversationState ConversationState
	RecentReply       *string
}

type ConversationState int

const (
	UserOneStartsConversation ConversationState = 0
	UserOneReplies            ConversationState = 1
	UserTwoReplies            ConversationState = 2
)

type Question int

type User struct {
	Id                      uuid.UUID
	LastActive              time.Time
	BlockedUsersIds         []uuid.UUID
	LastKnownConversationId *uuid.UUID
}

func (user User) IsActive(config Config) bool {
	activeDuration := time.Duration(config.UserActiveWhenLastSeenMinutesAgo) * time.Minute
	return user.LastActive.Add(activeDuration).After(time.Now())
}

func (conversation *Conversation) GetParticipantId(userId uuid.UUID) uuid.UUID {
	participantId := conversation.Participants[0]
	if participantId == userId {
		participantId = conversation.Participants[1]
	}
	return participantId
}
