package handler

import (
	"FollowService/dto"
	"FollowService/infrastructure/mapper"
	"FollowService/usecase"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"net/http"
)

type FollowingHandler interface {
	Unfollow(ctx *gin.Context)
	//Unfollow1(ctx *gin.Context)

	GetAllUsersFollowings(ctx *gin.Context)
	Follow(ctx *gin.Context)
	IsAllowedToFollow(ctx *gin.Context)
	GetAllFollowingFront(ctx *gin.Context)
}

type followingHandler struct {
	FollowingUseCase usecase.FollowingUseCase
	FollowingRequestUsecase usecase.FollowRequestUseCase
	logger *logger.Logger
}

func (f *followingHandler) GetAllFollowingFront(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)

	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	followers, err := f.FollowingUseCase.GetUserFollowingsForFrontend(context.Background(), t.ID)
	if err!=nil{
		//TODO: HANDLE RESPONSE ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, followers)
	return
}

func (f *followingHandler) IsAllowedToFollow(ctx *gin.Context) {
	f.logger.Logger.Println("Handling IS ALLOW TO FOLLOW ")

	decoder := json.NewDecoder(ctx.Request.Body)

	var followDto dto.FollowDTO
	if err := decoder.Decode(&followDto); err != nil {
		f.logger.Logger.Errorf("decoder error, error: %v\n", err)
		ctx.JSON(400, gin.H{"message": "Decoding error"})
		return
	}

	if followDto.Follower.ID==followDto.User.ID{
		f.logger.Logger.Errorf("same user follow error")
		ctx.JSON(400, gin.H{"message" : "Its you, you moron!"})
		return
	}

	profileFollowing := mapper.FollowDtoToProfileFollowing(followDto)
	if f.FollowingUseCase.AlreadyFollowing(ctx, profileFollowing) {
		ctx.JSON(400, gin.H{"message" : "You are already following user"})
		return
	}

	followingRequest := mapper.FollowDtoToFollowRequest(followDto)
	if f.FollowingRequestUsecase.IsCreated(ctx, followingRequest) {
		ctx.JSON(400 , gin.H{"message" : "Request already sent"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Allowed to follow"})
}

func (f *followingHandler) Follow(ctx *gin.Context) {
	f.logger.Logger.Println("Handling FOLLOW")

	decoder := json.NewDecoder(ctx.Request.Body)

	var followDto dto.FollowDTO
	if err := decoder.Decode(&followDto); err != nil {
		f.logger.Logger.Errorf("decoder error, error: %v\n", err)
		ctx.JSON(400, gin.H{"message" : "Decoding error"})
		return
	}

	profileFollowing := mapper.FollowDtoToProfileFollowing(followDto)
	if f.FollowingUseCase.AlreadyFollowing(ctx, profileFollowing) {
		ctx.JSON(400, gin.H{"message" : "You are already following user"})
		return
	}

	followingRequest := mapper.FollowDtoToFollowRequest(followDto)
	if f.FollowingRequestUsecase.IsCreated(ctx, followingRequest) {
		ctx.JSON(400, gin.H{"message" : "Request already sent"})
		return
	}

	_, err := f.FollowingUseCase.CreateFollowing(ctx, profileFollowing)
	if err != nil {
		f.logger.Logger.Errorf("failed to create following, error: %v\n", err)
		ctx.JSON(400, gin.H{"message" : "Error"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Success"})
}

func (f followingHandler) GetAllUsersFollowings(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GETTING ALL USERS FOLLOWINGS")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)

	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	followers, err := f.FollowingUseCase.GetAllUsersFollowings(t)
	if err!=nil{
		//TODO: HANDLE RESPONSE ERROR
		f.logger.Logger.Errorf("failed getting all users following, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": followers})
	return
}
//
//func (f followingHandler) Unfollow(ctx *gin.Context) {
//	f.logger.Logger.Println("Handling UNFOLLOW")
//
//	decoder := json.NewDecoder(ctx.Request.Body)
//
//	var t dto.Unfollow1
//	decode_err := decoder.Decode(&t)
//
//	if decode_err!=nil{
//		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
//		return
//	}
//
//	err := f.FollowingUseCase.Unfollow(ctx,t)
//	if err!= nil{
//		f.logger.Logger.Errorf("unfollow error, error: %v\n", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{})
//	return
//}

func (f followingHandler) Unfollow(ctx *gin.Context) {
	f.logger.Logger.Println("Handling UNFOLLOW")

	decoder := json.NewDecoder(ctx.Request.Body)

	var t dto.Unfollow
	decode_err := decoder.Decode(&t)

	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	err := f.FollowingUseCase.Unfollow(ctx,t.UserToUnfollow, t.UserUnfollowing)
	if err!= nil{
		f.logger.Logger.Errorf("unfollow error, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}
func NewFollowingHandler(u usecase.FollowingUseCase, f usecase.FollowRequestUseCase, logger *logger.Logger) FollowingHandler {
	return &followingHandler{u, f, logger}
}