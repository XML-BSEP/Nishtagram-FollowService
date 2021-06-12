package handler

import (
	"FollowService/dto"
	"FollowService/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"net/http"
)

type FollowRequestHandler interface {
	GetAllUsersFollowRequests(ctx *gin.Context)
	ApproveRequest(ctx *gin.Context)
	CancelFollowRequest(ctx *gin.Context)
}
type followRequestHandler struct {
	FollowRequestUseCase usecase.FollowRequestUseCase
	logger *logger.Logger
}
func (f followRequestHandler) CancelFollowRequest(ctx *gin.Context){
	f.logger.Logger.Println("Handling CANCELLATION FOLLOW REQUEST")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.FollowRequestDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowRequestUseCase.CancelFollowRequest(ctx, &t)
	if err!=nil{
		f.logger.Logger.Errorf("cancel follor request error, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message":"Succefully removed follow request!"})
	return
}
func (f followRequestHandler) GetAllUsersFollowRequests(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GETTING ALL USERS FOR FOLLOW REQUEST")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.ProfileDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}

	requests, err := f.FollowRequestUseCase.GetAllUsersFollowRequests(t)
	if err!=nil{
		f.logger.Logger.Errorf("get all users for follow requests, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": requests})
	return
}

func (f followRequestHandler) ApproveRequest(ctx *gin.Context){
	f.logger.Logger.Println("Handling APPROVE REQUEST")

	decoder := json.NewDecoder(ctx.Request.Body)
	var t dto.FollowRequestDTO
	decode_err := decoder.Decode(&t)
	if decode_err!=nil{
		f.logger.Logger.Errorf("decoder error, error: %v\n", decode_err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": decode_err.Error()})
		return
	}
	err := f.FollowRequestUseCase.ApprofeFollowRequest(ctx, t)
	if err !=nil{
		f.logger.Logger.Errorf("approve follow request error, error: %v\n", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data":"success"})
	return
}


func NewFollowRequestHandler(u usecase.FollowRequestUseCase, logger *logger.Logger) FollowRequestHandler {
	return &followRequestHandler{u, logger}
} 	