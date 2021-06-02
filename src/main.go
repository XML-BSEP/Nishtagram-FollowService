package main

import (
	"FollowService/http/router"
	"FollowService/infrastructure/mongo"
	"FollowService/infrastructure/seeder"
	"FollowService/interactor"
)

func main() {
	mongoCli, ctx := mongo.NewMongoClient()
	db := mongo.GetDbName()

	seeder.SeedData(db, mongoCli, ctx)

	i := interactor.NewInteractor(mongoCli)
	appHandler := i.NewAppHandler()

	g := router.NewRouter(appHandler)

	//g := gin.Default()
	//g.Use(gin.Logger())
	//g.Use(gin.Recovery())

	//TODO: check about possible changes in middleware
	//g.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "HEAD"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "*"
	//	},
	//	MaxAge: 12 * time.Hour,
	//}))
	//g.Use(cors.Default())
	//g.Use(CORSMiddleware())
	//g.Use(func(ctx *gin.Context) {
	//	ctx.Header("Access-Control-Allow-Origin", "*")
	//})

	//g.Use(func(ctx *gin.Context) {
	//	ctx.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	//})

	//corsConfig := cors.DefaultConfig()
	//
	//corsConfig.AllowOrigins = []string{"https://example.com"}
	//// To be able to send tokens to the server.
	//corsConfig.AllowCredentials = true
	//
	//// OPTIONS method for ReactJS
	//corsConfig.AddAllowMethods("OPTIONS", "POST")
	//
	//// Register the middleware
	//g.Use(cors.New(corsConfig))

	//todo: check this middleware thingy out
	//g.Use(static.Serve("/static", static.LocalFile("/assets", false)))

	g.Run("localhost:8089")
}
