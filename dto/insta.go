package dto

import (
	"errors"
	"gostagram/domain"
)

type InstaRequest struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}

type InstaResponse struct {
	Data domain.InstagramData `json:"data"`
}

func (i InstaRequest) Validate() error {
	if i.UserId == "" {
		return errors.New("required media id")
	}
	if i.AccessToken == "" {
		return errors.New("required access token")
	}
	return nil
}
