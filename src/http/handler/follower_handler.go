package handler

import (
	"FollowService/dto"
	"FollowService/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowerHandler interface {
	GetAllUsersFollowers(ctx *gin.Context)
}
type followerHandler struct {
	FollowerUseCase usecase.FollowerUseCase
}

func (f followerHandler) GetAllUsersFollowers(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	followers, err := f.FollowerUseCase.GetAllUsersFollowers(t)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"data": followers})
	return
}



func NewFollowerHandler(u usecase.FollowerUseCase) FollowerHandler {
	return &followerHandler{u}
}