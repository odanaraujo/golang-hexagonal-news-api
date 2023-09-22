package news_http

import (
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/odanaraujo/golang/hexagonal-news-api/adapter/output/response"
	"github.com/odanaraujo/golang/hexagonal-news-api/application/domain"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/env"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/exception"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/logger"
)

var (
	client *resty.Client
)

type newsClient struct{}

func NewNewsClient() *newsClient {
	client = resty.New().SetBaseURL(env.GetBaseURL())
	return &newsClient{}
}

func (nc *newsClient) GetNewsPort(newsDomain domain.NewsReqDomain) (*domain.NewsDomain, *exception.Exception) {

	newsResponse := &response.NewsClientResponse{}

	_, err := client.R().SetQueryParams(map[string]string{
		"q":      newsDomain.Subject,
		"from":   newsDomain.From,
		"apiKey": env.GetNewsTokenAPI(),
	}).SetResult(newsResponse).Get("/everything")

	if err != nil {
		logger.Error("Error on get news from API", err)
		return nil, exception.NewInternalServerError("Error trying to call NewsAPI with params")
	}

	newsResponseDomain := &domain.NewsDomain{}
	copier.Copy(newsResponseDomain, newsResponse)

	return newsResponseDomain, nil
}
