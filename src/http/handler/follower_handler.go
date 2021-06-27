package handler

import (
	"FollowService/dto"
	"FollowService/usecase"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"net/http"
)

type FollowerHandler interface {
	GetAllUsersFollowers(ctx *gin.Context)
	GetAllUsersFollowersFron(ctx *gin.Context)
	AddToCloseFriends(ctx *gin.Context)
	RemoveFromCloseFriends(ctx *gin.Context)
	GetAllUsersCloseFriends(ctx *gin.Context)
	GetAllUsersToWhomUserIsCloseFriend(ctx *gin.Context)
	IsUserFollowingUser(ctx *gin.Context)
}
type followerHandler struct {
	FollowerUseCase usecase.FollowerUseCase
	logger *logger.Logger
}


func (f followerHandler) IsUserFollowingUser(ctx *gin.Context) {
	f.logger.Logger.Println("Handling IS USER FOLLOWING USER")
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.FollowDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	isFollowing, err := f.FollowerUseCase.IsUserFollowingUser(ctx, t.Follower.ID, t.User.ID)
	if err!=nil{
		f.logger.Logger.Errorf("add to close friends, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, isFollowing)
	return
}

func (f followerHandler) GetAllUsersCloseFriends(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GET ALL USERS CLOSE FRIENDS")
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	closefriends, err := f.FollowerUseCase.GetAllUsersCloseFriends(ctx, t.ID)
	if err!=nil{
		f.logger.Logger.Errorf("get all users closefriends, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, closefriends)
	return
}

func (f followerHandler) GetAllUsersToWhomUserIsCloseFriend(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GET ALL USERS TO WHOM USER IS CLOSE FRIEND")
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	closefriends, err := f.FollowerUseCase.GetAllUsersToWhomUserIsCloseFriend(ctx, t.ID)
	if err!=nil{
		f.logger.Logger.Errorf("get all users closefriends, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, closefriends)
	return
}
func (f followerHandler) AddToCloseFriends(ctx *gin.Context){
	f.logger.Logger.Println("Handling ADD TO CLOSE FRIENDS")
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.CloseFriendDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowerUseCase.AddToCloseFriends(ctx, t.CloseFriend, t.User)
	if err!=nil{
		f.logger.Logger.Errorf("add to close friends, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
	return

}

func (f followerHandler) RemoveFromCloseFriends(ctx *gin.Context){
	f.logger.Logger.Println("Handling REMOVE FROM CLOSE FRIENDS")
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.CloseFriendDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowerUseCase.RemoveFromCloseFriends(ctx, t.CloseFriend, t.User)
	if err!=nil{
		f.logger.Logger.Errorf("remove from close friends, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"message":"success"})
	return
}


func (f followerHandler) GetAllUsersFollowersFron(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GET ALL USERS FOLLOWERS FOR FRONT")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	followers, err := f.FollowerUseCase.GetFollowersForFront(context.Background(), t.ID)
	if err!=nil{
		f.logger.Logger.Errorf("get all followers for front, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, followers)
	return
}

func (f followerHandler) GetAllUsersFollowers(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GET ALL USERS FOLLOWERS")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	followers, err := f.FollowerUseCase.GetAllUsersFollowers(t)
	if err!=nil{
		f.logger.Logger.Errorf("get all followers for front, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": followers})
	return
}



func NewFollowerHandler(u usecase.FollowerUseCase, logger *logger.Logger) FollowerHandler {
	return &followerHandler{u, logger}
}