package service

import "mygogin/myutils"

type XVideoService interface {
	GetUrl(url string) (string, error)
}

type xvideoServiceImpl struct {
	utils myutils.GetXVideoDownload
}

func NewXVideoService(utils myutils.GetXVideoDownload) XVideoService {
	return &xvideoServiceImpl{utils: utils}
}

func (s *xvideoServiceImpl) GetUrl(url string) (string, error) {
	// 这里以后可以加缓存、鉴权、组合逻辑
	return s.utils.GetXVideoDownloadUrl(url)
}
