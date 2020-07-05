package app

import (
	"github.com/gin-gonic/gin"

	NewsInterface "bareksa/app/news"
	TopicInterface "bareksa/app/topic"

	NWHandler "bareksa/app/news/handler"
	TPHandler "bareksa/app/topic/handler"
)

const (
	VERSION = "test"
	PROJECT = "bareksa"
)

func NewsHTTPHandler(r *gin.Engine, nu NewsInterface.INewsUsecase) {
	handler := &NWHandler.NewsHandler{
		NewsUsecase: nu,
	}

	version := r.Group(VERSION)
	project := version.Group(PROJECT)
	route := project.Group("/news")

	route.GET("/", handler.GetAllNews)
	route.GET("/:identifier", handler.GetNews)
	route.POST("/", handler.CreateNews)
	route.DELETE("/:identifier", handler.DeleteNews)
	route.PUT("/:identifier", handler.UpdateNews)
	route.PATCH("/:identifier", handler.PublishNews)
}

func TopicHTTPHandler(r *gin.Engine, tu TopicInterface.ITopicUsecase) {
	handler := &TPHandler.TopicHandler{
		TopicUsecase: tu,
	}

	version := r.Group(VERSION)
	project := version.Group(PROJECT)
	route := project.Group("/topics")

	route.GET("/", handler.GetAllTopics)
	route.GET("/:identifier", handler.GetTopic)
	route.POST("/", handler.CreateTopic)
	route.DELETE("/:identifier", handler.DeleteTopic)
	route.PUT("/:identifier", handler.UpdateTopic)
}
