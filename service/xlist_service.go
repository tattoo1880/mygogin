// xlist_service.go
package service

import "mygogin/myutils"

type XListService interface {
	GetList(id string) []string
}

type xlistServiceImpl struct {
	utils myutils.GetListUtils
}

func NewXListService(utils myutils.GetListUtils) XListService {
	return &xlistServiceImpl{utils: utils}
}

func (s *xlistServiceImpl) GetList(id string) []string {
	// 这里以后可以加缓存、鉴权、组合逻辑
	return s.utils.GetList(id)
}
