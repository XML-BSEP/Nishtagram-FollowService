package handler

import (
	"FollowService/dto"
	"FollowService/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowingHandler interface {
	Unfollow(ctx *gin.Context)
	GetAllUsersFollowings(ctx *gin.Context)

}

type followingHandler struct {
	FollowingUseCase usecase.FollowingUseCase
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

func NewFollowingHandler(u usecase.FollowingUseCase) FollowingHandler {
	return &followingHandler{u}
}