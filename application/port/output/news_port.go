package output

import (
	"github.com/odanaraujo/golang/hexagonal-news-api/application/domain"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
)

type NewsPort interface {
	GetNewsPort(domain.NewsReqDomain) (*domain.NewsDomain, *exception.Exception)
}
