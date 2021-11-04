package service

import (
	"encoding/json"
	"fmt"
	"gostagram/domain"
	"gostagram/dto"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

var GRAPH_URL string = "https://graph.instagram.com"
var MEIDA_QUERY string = "media?fields=caption,id,media_type,media_url,username,timestamp,permalink,thumbnail_url,children&access_token"
var CHILDREN_QUERY string = "?fields=children,id,media_type,media_url,thumbnail_url&access_token"

type InstaService interface {
	GetPosts(req dto.InstaRequest) (*dto.InstaResponse, error)
	GetChild(req dto.InstaChildRequest) (*dto.InstaChildResponse, error)
}

type DefaultInstaService struct {
}

func (s DefaultInstaService) GetPosts(req dto.InstaRequest) (*dto.InstaResponse, error) {
	var res dto.InstaResponse
	query := fmt.Sprintf("%s/%s/%s=%s", GRAPH_URL, req.UserId, MEIDA_QUERY, req.AccessToken)
	resp, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var insta domain.InstagramData

	if err := json.NewDecoder(resp.Body).Decode(&insta); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	res.Data = insta.Data
	return &res, nil
}

func (s DefaultInstaService) GetChild(req dto.InstaChildRequest) (*dto.InstaChildResponse, error) {
	var child domain.Child

	query := fmt.Sprintf("%s/%s/%s=%s", GRAPH_URL, req.ChildId, CHILDREN_QUERY, req.AccessToken)
	resp, err := http.Get(query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&child); err != nil {
		return nil, err
	}

	var res dto.InstaChildResponse
	res = dto.InstaChildResponse{
		Id:           child.Id,
		MediaType:    child.MediaType,
		MediaURL:     child.MediaURL,
		ThumbnailURL: child.ThumbnailURL,
	}
	return &res, nil
}

func NewInstaService() InstaService {
	return &DefaultInstaService{}
}
