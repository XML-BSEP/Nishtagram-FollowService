package gateway

import (
	"FollowService/dto"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	resty2 "github.com/go-resty/resty/v2"
	"os"
)

func IsProfilePrivate(ctx context.Context, userId string) (bool, error) {
	client := resty2.New()
	domain := os.Getenv("USER_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}

	if os.Getenv("DOCKER_ENV") == "" {
		resp, err := client.R().
			SetBody(gin.H{"id" : userId}).
			SetContext(ctx).
			EnableTrace().
			Post("https://" + domain + ":8082/isPrivate")

		if err != nil {
			return false, err
		}

		if resp.StatusCode() != 200 {
			return false, fmt.Errorf("Err")
		}

		var privacyCheckResponseDto dto.PrivacyCheckResponseDto
		if err := json.Unmarshal(resp.Body(), &privacyCheckResponseDto); err != nil {
			return false, err
		}

		return privacyCheckResponseDto.IsPrivate, err
	} else {
		resp, err := client.R().
			SetBody(gin.H{"id" : userId}).
			SetContext(ctx).
			EnableTrace().
			Post("http://" + domain + ":8082/isPrivate")

		if err != nil {
			return false, err
		}

		if resp.StatusCode() != 200 {
			return false, fmt.Errorf("Err")
		}

		var privacyCheckResponseDto dto.PrivacyCheckResponseDto
		if err := json.Unmarshal(resp.Body(), &privacyCheckResponseDto); err != nil {
			return false, err
		}

		return privacyCheckResponseDto.IsPrivate, err
	}

}


func GetUser(ctx context.Context, userId string) (dto.ProfileUsernameImageDTO, error) {
	client := resty2.New()
	domain := os.Getenv("USER_DOMAIN")
	if domain == "" {
		domain = "127.0.0.1"
	}

	if os.Getenv("DOCKER_ENV") == "" {
		resp, _ := client.R().
			EnableTrace().
			Get("https://" + domain + ":8082/getProfileUsernameImageById?userId=" + userId)

		var responseDTO dto.ProfileUsernameImageDTO
		err := json.Unmarshal(resp.Body(), &responseDTO)
		if err != nil {
			fmt.Println(err)
		}

		return responseDTO, nil
	} else {
		resp, _ := client.R().
			EnableTrace().
			Get("http://" + domain + ":8082/getProfileUsernameImageById?userId=" + userId)

		var responseDTO dto.ProfileUsernameImageDTO
		err := json.Unmarshal(resp.Body(), &responseDTO)
		if err != nil {
			fmt.Println(err)
		}

		return responseDTO, nil
	}

}

