package googleactions

import "time"

type AppRequest struct {
	User              User         `json:"user"`
	Surface           Surface      `json:"surface"`
	Conversation      Conversation `json:"conversation"`
	Input             []Input      `json:"inputs"`
	IsInSandbox       bool         `json:"isInSandbox"`
	AvailableSurfaces []Surface    `json:"availableSurfaces"`
}

type User struct {
	UserID      string     `json:"userId"`
	Locale      string     `json:"locale"`
	LastSeen    *time.Time `json:"lastSeen"`
	UserStorage *string    `json:"userStorage"`
}

type Surface struct {
	Capabilities []Capability `json:"capabilities"`
}

type Capability struct {
	Name string `json:"name"`
}

type Conversation struct {
	ConversationID    string `json:"conversationId"`
	Type              string `json:"type"`
	ConversationToken string `json:"conversationToken"`
}

type Input struct {
	RawInputs []RawInput `json:"rawInputs"`
	Intent    string     `json:"intent"`
	Arguments []Argument `json:"arguments"`
}

type RawInput struct {
	InputType string `json:"inputType"`
	Query     string `json:"query"`
	Url       string `json:"url"`
}

type Argument struct {
	Name      string    `json:"name"`
	Extension Extension `json:"extension"`
}

type Extension struct {
	Type   string `json:"@type"`
	Status string `json:"status"`
}
