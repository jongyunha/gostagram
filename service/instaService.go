package service

import (
	"gostagram/dto"
)

type InstaService interface {
	GetPosts(req dto.InstaRequest) (*dto.InstaResponse, error)
}

type DefaultInstaService struct {
}

func (s DefaultInstaService) GetPosts(req dto.InstaRequest) (*dto.InstaResponse, error) {
	var res dto.InstaResponse
	return &res, nil
}

func NewInstaService() InstaService {
	return &DefaultInstaService{}
}
