package service

import "mygogin/model"

type ChatMsgService struct {
	chatmsgrepo model.ChatMsgRepository
}

type ChatMsgServiceInterface interface {
	CreateChatMsg(msg *model.ChatMsg) error
	FindChatMsgsByFromUserId(fromUserId string) ([]*model.ChatMsg, error)
	FindChatMsgsByToUserId(toUserId string) ([]*model.ChatMsg, error)
	FindAllChatMsgs() ([]*model.ChatMsg, error)
	DeleteChatMsg(id string) error
}

func NewChatMsgService(chatmsgrepo model.ChatMsgRepository) ChatMsgServiceInterface {
	return &ChatMsgService{chatmsgrepo: chatmsgrepo}
}

func (s *ChatMsgService) CreateChatMsg(msg *model.ChatMsg) error {
	return s.chatmsgrepo.CreateChatMsg(msg)
}

func (s *ChatMsgService) FindChatMsgsByFromUserId(fromUserId string) ([]*model.ChatMsg, error) {
	return s.chatmsgrepo.FindChatMsgsByFromUserId(fromUserId)
}

func (s *ChatMsgService) FindChatMsgsByToUserId(toUserId string) ([]*model.ChatMsg, error) {
	return s.chatmsgrepo.FindChatMsgsByToUserId(toUserId)
}

func (s *ChatMsgService) FindAllChatMsgs() ([]*model.ChatMsg, error) {
	return s.chatmsgrepo.FindAllChatMsgs()
}

func (s *ChatMsgService) DeleteChatMsg(id string) error {
	return s.chatmsgrepo.DeleteChatMsg(id)
}
