package router

import (
	"FollowService/interactor"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func NewRouter(handler interactor.AppHandler) *gin.Engine{

	router := gin.New()
	router.Use(CORSMiddleware())

	router.POST("/unfollow", handler.Unfollow)
	router.POST("/usersFollowings", handler.GetAllUsersFollowings)
	router.POST("/usersFollowers", handler.GetAllUsersFollowers)
	router.POST("/usersFollowRequests", handler.GetAllUsersFollowRequests)
	router.POST("/approveRequest", handler.ApproveRequest)

	return router
}