package model

type StatsResponse struct {
	UsersCount         int `json:"usersCount"`
	ActiveUsersCount   int `json:"activeUsersCount"`
	ConversationsCount int `json:"conversationsCount"`
}
