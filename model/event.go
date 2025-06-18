package model

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type EventRepository interface {
	CreateEvent(event *Event) error
	GetEventByID(id int64) (*Event, error)
	FindAllEvents() ([]*Event, error)
	UpdateEvent(event *Event) error
	DeleteEvent(id int64) error
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepository {
	return &EventRepo{db: db}
}

func (r *EventRepo) CreateEvent(event *Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepo) GetEventByID(id int64) (*Event, error) {
	var event Event
	if err := r.db.First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepo) FindAllEvents() ([]*Event, error) {
	var events []*Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepo) UpdateEvent(event *Event) error {
	//! 如果 event.CreateAt 为零值，GORM 会自动设置为当前时间
	if event.CreatedAt.IsZero() {
		event.CreatedAt = time.Now()
	}
	return r.db.Save(event).Error
}
func (r *EventRepo) DeleteEvent(id int64) error {
	return r.db.Delete(&Event{}, id).Error
}
