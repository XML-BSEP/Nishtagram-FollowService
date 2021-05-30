package main

import (
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/repository"
	"FollowService/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)
	//profile4 := domain.Profile{ID: "123454"}
	//profile5 := domain.Profile{ID: "123455"}
	//following := domain.ProfileFollowing{Following: profile4, User: profile5, Timestamp: time.Now()}
	followingRepo := repository.NewFollowingRepo(mongoCli)
	followingService := usecase.NewFollowingService(followingRepo)

	//_, _ = followingRepo.CreateFollowing(&following)

	p,_ := followingService.GetByID("12341")
	fmt.Println(p)
	g := gin.Default()

	g.Run("localhost:8089")
}
