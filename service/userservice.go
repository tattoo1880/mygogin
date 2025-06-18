package service

import (
	"errors"
	"mygogin/model"
)

type UserService interface {
	//! CRUD
	CreateUser(user *model.User) error
	GetUserByID(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userService struct {
	repo model.UserRepository
}

func NewUserService(repo model.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	// 可以做业务校验
	if user.Name == "" {
		return errors.New("用户名不能为空")
	}
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id int64) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.repo.FindAllUsers()
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.DeleteUser(id)
}
