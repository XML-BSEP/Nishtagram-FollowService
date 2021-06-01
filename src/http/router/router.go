package router

import (
	"FollowService/interactor"
	"github.com/gin-gonic/gin"
)


func NewRouter(handler interactor.AppHandler) *gin.Engine{
	router := gin.Default()

	router.POST("/unfollow", handler.Unfollow)
	router.POST("/usersFollowings", handler.GetAllUsersFollowings)
	router.POST("/usersFollowers", handler.GetAllUsersFollowers)
	router.POST("/usersFollowRequests", handler.GetAllUsersFollowRequests)

	return router
}