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
}
type followerHandler struct {
	FollowerUseCase usecase.FollowerUseCase
	logger *logger.Logger
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