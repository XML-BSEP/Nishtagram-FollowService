package main

import (
	"FollowService/http/router"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/interactor"

	"github.com/gin-gonic/gin"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	i := interactor.NewInteractor(mongoCli)
	appHandler := i.NewAppHandler()

	g := router.NewRouter(appHandler)

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	g.Run(":8089")

}
