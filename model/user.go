package model

import "gorm.io/gorm"

type User struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRepository interface {
	//! CRUD
	CreateUser(user *User) error
	GetUserByID(id int64) (*User, error)
	FindAllUsers() ([]*User, error)
	UpdateUser(user *User) error
	DeleteUser(id int64) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(user *User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetUserByID(id int64) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) FindAllUsers() ([]*User, error) {
	var users []*User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) UpdateUser(user *User) error {
	return r.db.Save(user).Error
}

func (r *UserRepo) DeleteUser(id int64) error {
	return r.db.Delete(&User{}, id).Error
}
