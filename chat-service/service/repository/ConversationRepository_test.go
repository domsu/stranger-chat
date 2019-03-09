package repository

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
	"time"
)
import . "github.com/domsu/stranger-chat/model"

var user1 = User{Id: uuid.New(), LastActive: time.Now()}
var user2 = User{Id: uuid.New(), LastActive: time.Now()}
var user3 = User{Id: uuid.New(), LastActive: time.Now()}
var user4 = User{Id: uuid.New(), LastActive: time.Now()}

var conversation1 = Conversation{Id: uuid.New(), Participants: []uuid.UUID{user1.Id, user2.Id}, ConversationState: UserOneStartsConversation, RecentReply: nil}
var conversation2 = Conversation{Id: uuid.New(), Participants: []uuid.UUID{user3.Id, user4.Id}, ConversationState: UserOneStartsConversation, RecentReply: nil}

func TestNewConversationRepository_ShouldReturnCorrectRepository(t *testing.T) {
	repository := NewConversationRepository(Config{UserActiveWhenLastSeenMinutesAgo: 1})

	if repository.config.UserActiveWhenLastSeenMinutesAgo != 1 {
		t.Error("Wrong appconfig")
	}
	if repository.conversations != nil {
		t.Error("Wrong conversations")
	}
}

func TestGetConversationForUser_ShouldReturnCorrectConversation(t *testing.T) {
	repository := ConversationRepository{conversations: []*Conversation{&conversation1, &conversation2}}

	if repository.GetConversationForUser(user1.Id).Id != conversation1.Id {
		t.Error("Wrong conversation returned")
	}
	if repository.GetConversationForUser(user2.Id).Id != conversation1.Id {
		t.Error("Wrong conversation returned")
	}
}

func TestGetConversationForUser_ShouldReturnNil_WhenNoConversation(t *testing.T) {
	repository := ConversationRepository{conversations: []*Conversation{&conversation1}}

	if repository.GetConversationForUser(user3.Id) != nil {
		t.Error("Expected nil")
	}
}

func TestRemoveConversation_ShouldRemoveConversation(t *testing.T) {

	repository := ConversationRepository{conversations: []*Conversation{&conversation1, &conversation2}}

	repository.RemoveConversation(conversation2.Id)

	if len(repository.GetConversations()) != 1 || repository.GetConversations()[0].Id != conversation1.Id {
		t.Error("Wrong conversation deleted")
	}
}

func TestRemoveConversation_ShouldDoNothing_WhenNoConversation(t *testing.T) {
	repository := ConversationRepository{conversations: []*Conversation{&conversation1}}

	repository.RemoveConversation(conversation2.Id)

	if len(repository.GetConversations()) != 1 {
		t.Error("Wrong conversation deleted")
	}
}

func TestGetConversations_ShouldReturnConversations(t *testing.T) {
	repository := ConversationRepository{conversations: []*Conversation{&conversation1, &conversation2}}

	if len(repository.GetConversations()) != 2 {
		t.Error("Wrong conversations")
	}
}

func TestCreateConversation_ShouldCreateCorrectConversationAndAddToRepository(t *testing.T) {
	repository := ConversationRepository{}

	conversation := repository.CreateConversation(user1.Id, user2.Id)

	if !reflect.DeepEqual(
		*conversation,
		Conversation{Id: conversation.Id, Participants: []uuid.UUID{user1.Id, user2.Id}, ConversationState: UserOneStartsConversation, RecentReply: nil},
	) {
		t.Error("Invalid conversation")
	}
	if len(repository.GetConversations()) != 1 {
		t.Error("Conversation not added to repository")
	}
}
