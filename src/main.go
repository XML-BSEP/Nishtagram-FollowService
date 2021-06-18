package main

import (
	"FollowService/http/router"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/interactor"
	"context"
	logger "github.com/jelena-vlajkov/logger/logger"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.InitializeLogger("follow-service", context.Background())

	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	i := interactor.NewInteractor(mongoCli, logger)
	appHandler := i.NewAppHandler()

	g := router.NewRouter(appHandler)

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	//g.Run(":8089")
	if os.Getenv("DOCKER_ENV") == "" {
		err := g.RunTLS(":8089", "certificate/cert.pem", "certificate/key.pem")
		if err != nil {
			return 
		}
		
	} else {
		err := g.Run(":8089")
		if err != nil {
			return 
		}
	}
}
