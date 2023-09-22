package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/hexagonal-news-api/adapter/input/model/request"
	"github.com/odanaraujo/golang/hexagonal-news-api/application/domain"
	"github.com/odanaraujo/golang/hexagonal-news-api/application/port/input"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/logger"
	"github.com/odanaraujo/golang/hexagonal-news-api/configurations/validation"
)

type newsController struct {
	NewsUseCase input.NewsUseCase
}

func NewNewsController(newUseCase input.NewsUseCase) *newsController {
	return &newsController{newUseCase}
}

// https://newsapi.org/v2/everything?q=tesla&from=2023-08-22&sortBy=publishedAt&apiKey=ddffc280fade41e482c22e7229b3a6a1
func (nc *newsController) GetNews(c *gin.Context) {

	logger.Info("Init the GetNews controller api")
	request := request.NewsRequest{}

	if err := c.ShouldBindQuery(&request); err != nil {
		logger.Error("Error trying to bind the request", err)
		err := validation.ValidateUserError(err)
		c.JSON(err.Code, err)
		return
	}

	new := domain.NewsReqDomain{Subject: request.Subject, From: request.From.Format("2006-01-02")}

	newsResponseDomain, err := nc.NewsUseCase.GetNewsService(new)

	if err != nil {
		logger.Error("Error in the GetNewsService", err)
		err := validation.ValidateUserError(err)
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, newsResponseDomain)
}
