package service

import (
	"errors"
	"mygogin/model"
)

type EventService interface {
	CreateEvent(event *model.Event) error
	GetEventByID(id int64) (*model.Event, error)
	GetAllEvents() ([]*model.Event, error)
	UpdateEvent(event *model.Event) error
	DeleteEvent(id int64) error
}

type eventService struct {
	repo model.EventRepository
}

func NewEventService(repo model.EventRepository) EventService {
	return &eventService{repo: repo}
}

func (s *eventService) CreateEvent(event *model.Event) error {
	// 可以做业务校验
	if event.Content == "" {
		return errors.New("事件内容不能为空")
	}
	return s.repo.CreateEvent(event)
}

func (s *eventService) GetEventByID(id int64) (*model.Event, error) {
	return s.repo.GetEventByID(id)
}
func (s *eventService) GetAllEvents() ([]*model.Event, error) {
	return s.repo.FindAllEvents()
}

func (s *eventService) UpdateEvent(event *model.Event) error {
	return s.repo.UpdateEvent(event)
}

func (s *eventService) DeleteEvent(id int64) error {
	return s.repo.DeleteEvent(id)
}
