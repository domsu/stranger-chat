package model

type Config struct {
	UserActiveWhenLastSeenMinutesAgo int    `yaml:"userActiveWhenLastSeenMinutesAgo"`
	MusicOnHoldUrl                   string `yaml:"musicOnHoldUrl"`
	ServiceDomain                    string `yaml:"serviceDomain"`
}

type Messages struct {
	Welcome                   string   `yaml:"welcome"`
	WaitingForParticipant     string   `yaml:"waitingForParticipant"`
	ParticipantFoundInitiator string   `yaml:"participantFoundInitiator"`
	MessageSent               string   `yaml:"messageSent"`
	WaitForReply              string   `yaml:"waitForReply"`
	ParticipantLeft           string   `yaml:"participantLeft"`
	ParticipantSay            string   `yaml:"participantSay"`
	ParticipantReplyPrefix    string   `yaml:"participantReplyPrefix"`
	ParticipantReplySufix     string   `yaml:"participantReplySufix"`
	RespondToParticipant      string   `yaml:"respondToParticipant"`
	SeeYouNextTime            string   `yaml:"seeYouNextTime"`
	PlayManually              string   `yaml:"playManually"`
	MusicTitle                string   `yaml:"musicTitle"`
	MusicDescription          string   `yaml:"musicDescription"`
	CheckStatus               string   `yaml:"checkStatus"`
	NextParticipantWords      []string `yaml:"nextParticipantWords"`
	YesWords                  []string `yaml:"yesWords"`
}
