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
	Data []domain.Instagram `json:"data"`
}

type InstaChildRequest struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
	ChildId     string `json:"child_id"`
}

type InstaChildResponse struct {
	Id           string `json:"id"`
	MediaType    string `json:"media_type"`
	MediaURL     string `json:"media_url"`
	Permalink    string `json:"permalink"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type Children struct {
	Data []ChildrenObj `json:"data"`
}

type ChildrenObj struct {
	Id string `json:"id"`
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
