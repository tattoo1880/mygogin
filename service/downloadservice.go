package service

import "mygogin/model"

type DownloadService struct {
	downloadrepo model.DonwLoadRepository
}

type DownloadServiceInterface interface {
	CreateDonwLoad(donwload *model.DonwLoad) error
	FindAllDonwLoads() ([]*model.DonwLoad, error)
	DeleteDonwLoad(id string) error
}

func NewDownloadService(downloadrepo model.DonwLoadRepository) DownloadServiceInterface {
	return &DownloadService{downloadrepo: downloadrepo}
}

func (s *DownloadService) CreateDonwLoad(donwload *model.DonwLoad) error {
	return s.downloadrepo.CreateDonwLoad(donwload)
}
func (s *DownloadService) FindAllDonwLoads() ([]*model.DonwLoad, error) {
	return s.downloadrepo.FindAllDonwLoads()
}

func (s *DownloadService) DeleteDonwLoad(id string) error {
	return s.downloadrepo.DeleteDonwLoad(id)
}
