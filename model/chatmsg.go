package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ChatMsg struct {
	ID         string    `gorm:"type:char(36);primaryKey;" json:"id"`
	FromUserID string    `gorm:"type:varchar(100);not null" json:"from_userid"`
	ToUserID   string    `gorm:"type:varchar(100);not null" json:"to_userid"`
	Msg        string    `gorm:"type:text;not null" json:"msg"`
	TimeStamp  time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

type ChatMsgRepository interface {
	//! CRUD
	CreateChatMsg(msg *ChatMsg) error
	FindChatMsgsByFromUserId(fromUserId string) ([]*ChatMsg, error)
	FindChatMsgsByToUserId(toUserId string) ([]*ChatMsg, error)
	FindAllChatMsgs() ([]*ChatMsg, error)
	DeleteChatMsg(id string) error
}

type ChatMsgRepo struct {
	db *gorm.DB
}

func NewChatMsgRepo(db *gorm.DB) ChatMsgRepository {
	return &ChatMsgRepo{db: db}
}

func (r *ChatMsgRepo) CreateChatMsg(msg *ChatMsg) error {
	if msg.TimeStamp.IsZero() {
		msg.TimeStamp = time.Now()
	}
	if msg.ID == "" {
		msg.ID = uuid.New().String()
	}
	return r.db.Create(msg).Error
}

func (r *ChatMsgRepo) FindChatMsgsByFromUserId(fromUserId string) ([]*ChatMsg, error) {
	var msgs []*ChatMsg
	if err := r.db.Where("from_user_id = ?", fromUserId).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *ChatMsgRepo) FindChatMsgsByToUserId(toUserId string) ([]*ChatMsg, error) {
	var msgs []*ChatMsg
	if err := r.db.Where("to_user_id = ?", toUserId).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *ChatMsgRepo) FindAllChatMsgs() ([]*ChatMsg, error) {
	var msgs []*ChatMsg
	if err := r.db.Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func (r *ChatMsgRepo) DeleteChatMsg(id string) error {
	return r.db.Delete(&ChatMsg{}, id).Error
}
