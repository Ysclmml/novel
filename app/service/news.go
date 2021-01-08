package service

import (
	"github.com/pkg/errors"
	"novel/app/dao"
	"novel/app/model"
)

var newsDao = dao.News{}

type NewsService struct {
}

func (ns *NewsService) ListIndexNews() *[]model.News {
	return newsDao.ListIndexNews()
}

func (ns *NewsService) ListByPage(page int, pageSize int) (*[]model.News, int64) {
	return newsDao.ListByPage(page, pageSize)
}

func (ns *NewsService) QueryNewsInfo(newsId int64) (*model.News, error) {
	news := newsDao.QueryNewsInfo(newsId)
	if news.Id <= 0{
		return nil, errors.New("新闻不存在")
	}
	return news, nil
}


