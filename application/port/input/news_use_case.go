package input

import (
	"github.com/odanaraujo/golang/hexagonal-news-api/application/domain"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
)

type NewsUseCase interface {
	GetNewsService(new domain.NewsReqDomain) (*domain.NewsDomain, *exception.Exception)
}
