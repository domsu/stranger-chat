package model

import (
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestIsActive_ShouldReturnTrue_WhenUserActive(t *testing.T) {
	user := User{LastActive: time.Now().Add(time.Duration(-3) * time.Minute)}

	if !user.IsActive(Config{UserActiveWhenLastSeenMinutesAgo: 5}) {
		t.Fatal()
	}
}

func TestIsActive_ShouldReturnFalse_WhenUserNotActive(t *testing.T) {
	user := User{LastActive: time.Now().Add(time.Duration(-6) * time.Minute)}

	if user.IsActive(Config{UserActiveWhenLastSeenMinutesAgo: 5}) {
		t.Fatal()
	}
}

func TestGetParticipantId_ShouldReturnCorrectUser(t *testing.T) {
	user1 := User{Id: uuid.New()}
	user2 := User{Id: uuid.New()}
	conv := Conversation{Participants: []uuid.UUID{user1.Id, user2.Id}}

	if conv.GetParticipantId(user1.Id) != user2.Id {
		t.Fatal()
	}

	if conv.GetParticipantId(user2.Id) != user1.Id {
		t.Fatal()
	}
}
