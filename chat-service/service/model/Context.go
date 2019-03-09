package model

import "github.com/google/uuid"

type Context struct {
	User            *User
	UserStorage     UserStorage
	PartialRequest  PartialRequest
	PartialResponse PartialResponse
	Config          Config
	Messages        Messages
}

type UserStorage struct {
	UserId             *uuid.UUID `json:"userId"`
	OnboardingFinished bool       `json:"onboardingFinished"`
}

type PartialResponse struct {
	TextResponse       string
	FinalResponse      string
	MultiMediaResponse MultiMediaResponse
	ExpectUserResponse bool
}

type PartialRequest struct {
	Type        RequestType
	Text        string
	Storage     *string
	InputType   InputType
	HealthCheck bool
}
type RequestType int

type MultiMediaResponse int

type InputType int

const (
	Action RequestType = iota
	Text
)

const (
	Empty MultiMediaResponse = iota
	WaitingForParticipantSound
)

const (
	Voice InputType = iota
	Keyboard
)
