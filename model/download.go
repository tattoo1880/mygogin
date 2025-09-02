package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DonwLoad struct {
	ID          string    `gorm:"type:char(36);primaryKey;" json:"id"`
	DonwLoadUrl string    `gorm:"type:text;not null" json:"donwload_url"`
	TimeStamp   time.Time `gorm:"autoCreateTime" json:"timestamp"`
}

type DonwLoadRepository interface {
	//! CRUD
	CreateDonwLoad(donwload *DonwLoad) error
	FindAllDonwLoads() ([]*DonwLoad, error)
	DeleteDonwLoad(id string) error
}

type DonwLoadRepo struct {
	db *gorm.DB
}

func NewDonwLoadRepo(db *gorm.DB) DonwLoadRepository {
	return &DonwLoadRepo{db: db}
}

func (r *DonwLoadRepo) CreateDonwLoad(donwload *DonwLoad) error {
	if donwload.TimeStamp.IsZero() {
		donwload.TimeStamp = time.Now()
	}
	if donwload.ID == "" {
		donwload.ID = uuid.New().String()
	}
	return r.db.Create(donwload).Error
}

func (r *DonwLoadRepo) FindAllDonwLoads() ([]*DonwLoad, error) {
	var donwloads []*DonwLoad
	if err := r.db.Find(&donwloads).Error; err != nil {
		return nil, err
	}
	return donwloads, nil
}

func (r *DonwLoadRepo) DeleteDonwLoad(id string) error {
	return r.db.Delete(&DonwLoad{}, "id = ?", id).Error
}
