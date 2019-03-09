package middleware

import (
	. "github.com/domsu/stranger-chat/model"
	. "github.com/domsu/stranger-chat/repository"
	"sort"
	"strings"
)

type SetResponse struct {
	userRepository         IUserRepository
	conversationRepository IConversationRepository
}

func NewSetResponse(userRepository IUserRepository, conversationRepository IConversationRepository) SetResponse {
	return SetResponse{userRepository: userRepository, conversationRepository: conversationRepository}
}

func (action SetResponse) Process(ctx *Context) bool {
	ctx.PartialResponse.ExpectUserResponse = true
	action.handleOnboarding(ctx)

	conversation := action.conversationRepository.GetConversationForUser(ctx.User.Id)

	var activeConversation *Conversation = nil
	if action.hasParticipantSkippedConversation(*ctx.User, conversation) {
		if ctx.User.IsActive(ctx.Config) {
			ctx.PartialResponse.TextResponse += ctx.Messages.ParticipantLeft
		}
		ctx.User.LastKnownConversationId = nil
	} else if action.hasParticipantAbandonedConversation(ctx, conversation) {
		if ctx.User.IsActive(ctx.Config) {
			ctx.PartialResponse.TextResponse += ctx.Messages.ParticipantLeft
		}
		ctx.User.LastKnownConversationId = nil
		action.conversationRepository.RemoveConversation(conversation.Id)
	} else {
		activeConversation = conversation
	}

	if activeConversation == nil {
		action.handleNoConversation(ctx)
	} else {
		action.handleActiveConversation(ctx, activeConversation)
	}

	return true
}

func (action SetResponse) hasParticipantAbandonedConversation(ctx *Context, conversation *Conversation) bool {
	if conversation == nil {
		return false
	}

	participantId := conversation.GetParticipantId(ctx.User.Id)
	participant := action.userRepository.GetUser(participantId)
	return !participant.IsActive(ctx.Config)
}

func (action SetResponse) hasParticipantSkippedConversation(user User, conversationForUser *Conversation) bool {
	return user.LastKnownConversationId != nil && conversationForUser == nil
}

func (action SetResponse) handleOnboarding(ctx *Context) {
	if !ctx.UserStorage.OnboardingFinished {
		ctx.PartialResponse.TextResponse = ctx.Messages.Welcome
		ctx.UserStorage.OnboardingFinished = true
	}
}

func (action SetResponse) handleNoConversation(ctx *Context) {
	waitingUsers := action.getPotentialParticipants(*ctx.User)

	if len(waitingUsers) < 1 {
		ctx.PartialResponse.TextResponse += ctx.Messages.WaitingForParticipant
		if ctx.PartialRequest.InputType == Keyboard {
			ctx.PartialResponse.TextResponse += ctx.Messages.PlayManually
		}
		ctx.PartialResponse.MultiMediaResponse = WaitingForParticipantSound
	} else {
		participant := waitingUsers[0]
		conversation := action.conversationRepository.CreateConversation(ctx.User.Id, participant.Id)
		participant.LastKnownConversationId = &conversation.Id
		ctx.User.LastKnownConversationId = &conversation.Id
		ctx.PartialResponse.TextResponse += ctx.Messages.ParticipantFoundInitiator
	}
}

func (action SetResponse) handleActiveConversation(ctx *Context, conversation *Conversation) {
	var replying = (conversation.Participants[0] == ctx.User.Id &&
		(conversation.ConversationState == UserOneStartsConversation || conversation.ConversationState == UserOneReplies)) ||
		(conversation.Participants[1] == ctx.User.Id && conversation.ConversationState == UserTwoReplies)

	if replying {
		action.handleReplyingInConversation(ctx, conversation)
	} else {
		action.handleWaitingInConversation(conversation, ctx)
	}
}

func (action SetResponse) handleReplyingInConversation(ctx *Context, conversation *Conversation) {
	if ctx.PartialRequest.Type == Text {
		if isWordInArray(ctx.PartialRequest.Text, ctx.Messages.NextParticipantWords) {
			ctx.User.BlockedUsersIds = append(ctx.User.BlockedUsersIds, conversation.GetParticipantId(ctx.User.Id))
			action.conversationRepository.RemoveConversation(conversation.Id)
			ctx.User.LastKnownConversationId = nil
			action.handleNoConversation(ctx)
		} else {
			action.handleUserReplyWithoutKeywords(conversation, ctx)
		}
	} else {
		action.deliverLastReplyIfExists(conversation, ctx)
	}
}

func (action SetResponse) handleUserReplyWithoutKeywords(conversation *Conversation, ctx *Context) {
	conversation.RecentReply = &ctx.PartialRequest.Text
	if conversation.ConversationState == UserTwoReplies {
		conversation.ConversationState = UserOneReplies
	} else {
		conversation.ConversationState = UserTwoReplies
	}
	ctx.PartialResponse.TextResponse += ctx.Messages.MessageSent
	if ctx.PartialRequest.InputType == Keyboard {
		ctx.PartialResponse.TextResponse += ctx.Messages.PlayManually
	}
	ctx.PartialResponse.MultiMediaResponse = WaitingForParticipantSound
}

func (action SetResponse) deliverLastReplyIfExists(conversation *Conversation, ctx *Context) {
	if conversation.RecentReply != nil {
		ctx.PartialResponse.TextResponse += ctx.Messages.ParticipantSay +
			ctx.Messages.ParticipantReplyPrefix +
			*conversation.RecentReply +
			ctx.Messages.ParticipantReplySufix

		ctx.PartialResponse.TextResponse += ctx.Messages.RespondToParticipant
	} else {
		ctx.PartialResponse.TextResponse += ctx.Messages.ParticipantFoundInitiator
	}
}

func (action SetResponse) handleWaitingInConversation(conversation *Conversation, ctx *Context) {
	if conversation.ConversationState == UserOneStartsConversation {
		ctx.PartialResponse.TextResponse += ctx.Messages.WaitingForParticipant
	} else {
		ctx.PartialResponse.TextResponse += ctx.Messages.WaitForReply
	}
	if ctx.PartialRequest.InputType == Keyboard {
		ctx.PartialResponse.TextResponse += ctx.Messages.PlayManually
	}
	ctx.PartialResponse.MultiMediaResponse = WaitingForParticipantSound
}

func isWordInArray(text string, words []string) bool {
	for i := range words {
		if strings.ToLower(words[i]) == strings.ToLower(text) {
			return true
		}
	}

	return false
}

func (action SetResponse) getPotentialParticipants(currentUser User) []*User {
	var waitingUsers []*User
	activeUsers := action.userRepository.GetActiveUsers()
	for i, v := range activeUsers {
		if v.Id != currentUser.Id {
			if action.conversationRepository.GetConversationForUser(v.Id) == nil {
				var blocked = false
				for _, blockedUserId := range currentUser.BlockedUsersIds {
					if blockedUserId == v.Id {
						blocked = true
					}
				}

				for _, blockedUserId := range v.BlockedUsersIds {
					if blockedUserId == currentUser.Id {
						blocked = true
					}
				}

				if !blocked {
					waitingUsers = append(waitingUsers, activeUsers[i])
				}
			}
		}
	}

	sort.Slice(waitingUsers, func(i, j int) bool {
		return waitingUsers[i].LastActive.After(waitingUsers[j].LastActive)
	})

	return waitingUsers
}
