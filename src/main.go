package main

import (
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"github.com/gin-gonic/gin"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	g := gin.Default()

	g.Run("localhost:8089")
}
