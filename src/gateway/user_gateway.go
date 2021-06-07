package gateway

import (
	"FollowService/dto"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	resty2 "github.com/go-resty/resty/v2"
)

func IsProfilePrivate(ctx context.Context, userId string) (bool, error) {
	client := resty2.New()

	resp, err := client.R().
		SetBody(gin.H{"id" : userId}).
		SetContext(ctx).
		EnableTrace().
		Post("https://localhost:8082/isPrivate")

	if err != nil {
		return false, err
	}


	var privacyCheckResponseDto dto.PrivacyCheckResponseDto
	if err := json.Unmarshal(resp.Body(), &privacyCheckResponseDto); err != nil {
		return false, err
	}

	return privacyCheckResponseDto.IsPrivate, err
}
