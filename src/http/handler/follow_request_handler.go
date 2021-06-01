package handler

import (
	"FollowService/dto"
	"FollowService/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowRequestHandler interface {
	GetAllUsersFollowRequests(ctx *gin.Context)
}
type followRequestHandler struct {
	FollowRequestUseCase usecase.FollowRequestUseCase
}

func (f followRequestHandler) GetAllUsersFollowRequests(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
	}

	requests, err := f.FollowRequestUseCase.GetAllUsersFollowRequests(t)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}
	ctx.JSON(http.StatusOK, gin.H{"data": requests})
}

func NewFollowRequestHandler(u usecase.FollowRequestUseCase) FollowRequestHandler {
	return &followRequestHandler{u}
} 	