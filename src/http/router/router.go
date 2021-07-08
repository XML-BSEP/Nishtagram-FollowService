package router

import (
	"FollowService/interactor"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewRouter(handler interactor.AppHandler) *gin.Engine {

	router := gin.Default()
	router.Use(CORSMiddleware())


	router.POST("/unfollow", handler.Unfollow)
	router.POST("/usersFollowings", handler.GetAllUsersFollowings)
	router.POST("/usersFollowers", handler.GetAllUsersFollowers)
	router.POST("/usersFollowRequests", handler.GetAllUsersFollowRequests)
	router.POST("/approveRequest", handler.ApproveRequest)
	router.POST("/follow", handler.Follow)
	router.POST("/isAllowedToFollow", handler.IsAllowedToFollow)
	router.POST("/cancelFollowRequest", handler.CancelFollowRequest)
	router.POST("/getAllUsersFollowingFront", handler.GetAllFollowingFront)
	router.POST("/getAllUsersFollowersFront", handler.GetAllUsersFollowersFron)
	router.POST("/addToCloseFriends", handler.AddToCloseFriends)
	router.POST("/removeFromCloseFriends", handler.RemoveFromCloseFriends)
	router.POST("/getAllUsersCloseFriends", handler.GetAllUsersCloseFriends)
	router.POST("/getAllUsersToWhomUserIsCloseFriend", handler.GetAllUsersToWhomUserIsCloseFriend)
	router.POST("/isUserFollowingUser", handler.IsUserFollowingUser)
	router.POST("/approveAllRequests", handler.ApproveAllRequests)
	router.POST("/banUser", handler.BanUser)
	router.GET("/recommend/:userId", handler.Recommend)

	return router
}
