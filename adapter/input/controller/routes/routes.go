package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/golang/hexagonal-news-api/adapter/input/controller"
	"github.com/odanaraujo/golang/hexagonal-news-api/adapter/output/news_http"
	"github.com/odanaraujo/golang/hexagonal-news-api/application/service"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)
	r.GET("/news", newsController.GetNews)
}
