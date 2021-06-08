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
	ApproveRequest(ctx *gin.Context)
	CancelFollowRequest(ctx *gin.Context)
}
type followRequestHandler struct {
	FollowRequestUseCase usecase.FollowRequestUseCase
}
func (f followRequestHandler) CancelFollowRequest(ctx *gin.Context){
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.FollowRequestDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowRequestUseCase.CancelFollowRequest(ctx, &t)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Succefully removed follow request!"})
	return
}
func (f followRequestHandler) GetAllUsersFollowRequests(ctx *gin.Context) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	requests, err := f.FollowRequestUseCase.GetAllUsersFollowRequests(t)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": requests})
	return
}

func (f followRequestHandler) ApproveRequest(ctx *gin.Context){
	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.FollowRequestDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowRequestUseCase.ApprofeFollowRequest(ctx, t)
	if err !=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data":"success"})
	return
}


func NewFollowRequestHandler(u usecase.FollowRequestUseCase) FollowRequestHandler {
	return &followRequestHandler{u}
} 	