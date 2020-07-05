package main

import (
	routes "bareksa/app"
	"bareksa/config"
	gorm "bareksa/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	NWRepository "bareksa/app/news/repository"
	TPRepository "bareksa/app/topic/repository"

	NWUsecase "bareksa/app/news/usecase"
	TPUsecase "bareksa/app/topic/usecase"
)

var appConfig = config.Config.App

func main() {
	r := gin.New()

	db := gorm.MysqlConn()

	r.Use(gin.Recovery())

	//Repository Initialize
	nwr := NWRepository.NewNewsRepository(db)
	tpr := TPRepository.NewTopicRepository(db)

	//Usecase Initialize
	nwu := NWUsecase.NewNewsUsecase(nwr)
	tpu := TPUsecase.NewTopicUsecase(tpr)

	//Routes
	routes.NewsHTTPHandler(r, nwu)
	routes.TopicHTTPHandler(r, tpu)

	if err := r.Run(fmt.Sprintf(":%s", appConfig.HTTPPort)); err != nil {
		log.Fatal(err)
	}
}
