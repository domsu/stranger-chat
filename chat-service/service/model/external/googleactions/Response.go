package googleactions

type Response struct {
	ConversationToken  string          `json:"conversationToken,omitempty"`
	UserStorage        string          `json:"userStorage,omitempty"`
	ExpectUserResponse bool            `json:"expectUserResponse"`
	ExpectedInputs     []ExpectedInput `json:"expectedInputs"`
	FinalResponse      *FinalResponse  `json:"finalResponse,omitempty"`
}

type ExpectedInput struct {
	InputPrompt        InputPrompt      `json:"inputPrompt"`
	PossibleIntents    []ExpectedIntent `json:"possibleIntents"`
	SpeechBiasingHints []string         `json:"speechBiasingHints"`
}

type InputPrompt struct {
	RichInitialPrompt RichResponse `json:"richInitialPrompt"`
}

type ExpectedIntent struct {
	Intent string `json:"intent"`
}

type FinalResponse struct {
	RichResponse RichResponse `json:"richResponse"`
}

type RichResponse struct {
	Items       []Item        `json:"items"`
	Suggestions *[]Suggestion `json:"suggestions,omitempty"`
}

type Item struct {
	SimpleResponse *SimpleResponse `json:"simpleResponse,omitempty"`
	MediaResponse  *MediaResponse  `json:"mediaResponse,omitempty"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	Ssml         string `json:"ssml,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

type MediaResponse struct {
	MediaType    string        `json:"mediaType"`
	MediaObjects []MediaObject `json:"mediaObjects"`
}

type MediaObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ContentUrl  string `json:"contentUrl"`
}

type Suggestion struct {
	Title string `json:"title"`
}
