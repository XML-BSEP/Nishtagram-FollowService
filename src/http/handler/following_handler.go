package handler

import (
	"FollowService/dto"
	"FollowService/infrastructure/mapper"
	"FollowService/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowingHandler interface {
	Unfollow(ctx *gin.Context)
	GetAllUsersFollowings(ctx *gin.Context)
	Follow(ctx *gin.Context)
	IsAllowedToFollow(ctx *gin.Context)

}

type followingHandler struct {
	FollowingUseCase usecase.FollowingUseCase
	FollowingRequestUsecase usecase.FollowRequestUseCase
}

func (f *followingHandler) IsAllowedToFollow(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)

	var followDto dto.FollowDTO
	if err := decoder.Decode(&followDto); err != nil {
		ctx.JSON(400, gin.H{"message": "Decoding error"})
		return
	}

	if followDto.Follower.ID==followDto.User.ID{
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
	decoder := json.NewDecoder(ctx.Request.Body)

	var followDto dto.FollowDTO
	if err := decoder.Decode(&followDto); err != nil {
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
		ctx.JSON(400, gin.H{"message" : "Error"})
		return
	}

	ctx.JSON(200, gin.H{"message" : "Success"})
}

func (f followingHandler) GetAllUsersFollowings(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)

	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	followers, err := f.FollowingUseCase.GetAllUsersFollowings(t)
	if err!=nil{
		//TODO: HANDLE RESPONSE ERROR
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": followers})
	return
}

func (f followingHandler) Unfollow(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)

	var t dto.Unfollow
	decode_err := decoder.Decode(&t)

	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	err := f.FollowingUseCase.Unfollow(ctx,t)
	if err!= nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
	return
}

func NewFollowingHandler(u usecase.FollowingUseCase, f usecase.FollowRequestUseCase) FollowingHandler {
	return &followingHandler{u, f}
}