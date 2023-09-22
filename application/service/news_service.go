package service

import (
	"fmt"

	"github.com/odanaraujo/golang/hexagonal-news-api/application/domain"
	"github.com/odanaraujo/golang/hexagonal-news-api/application/port/output"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/logger"
)

type newsService struct {
	newsPort output.NewsPort
}

func NewNewsService(newsPort output.NewsPort) *newsService {
	return &newsService{newsPort}
}

func (ns *newsService) GetNewsService(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *exception.Exception) {
	logger.Info(fmt.Sprintf("Init GetNewsService function, subject=%s, from=%s", newsDomain.Subject, newsDomain.From))

	newsDomainResponse, err := ns.newsPort.GetNewsPort(newsDomain)

	return newsDomainResponse, err
}
