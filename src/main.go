package main

import (
	"FollowService/http/router"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/interactor"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	i := interactor.NewInteractor(mongoCli)
	appHandler := i.NewAppHandler()
	g := router.NewRouter(appHandler)

	//g := gin.Default()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	//TODO: check about possible changes in middleware
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	}))

	//todo: check this middleware thingy out
	//g.Use(static.Serve("/static", static.LocalFile("/assets", false)))

	g.Run("localhost:8089")
}
