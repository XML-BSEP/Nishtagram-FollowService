package main

import (
	"FollowService/domain"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/repository"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)
	profile4 := domain.Profile{ID: "123454"}
	profile5 := domain.Profile{ID: "123455"}
	following := domain.ProfileFollowing{Following: profile4, User: profile5, Timestamp: time.Now()}
	followingRepo := repository.NewFollowingRepo(mongoCli)

	_, _ = followingRepo.CreateFollowing(&following)

	g := gin.Default()

	g.Run("localhost:8089")
}
