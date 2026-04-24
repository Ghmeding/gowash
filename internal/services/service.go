package services

import "example.com/internal/storage"

type Service struct {
	repo storage.Storage
}

func NewService(repo storage.Storage) *Service {
	return &Service{repo: repo}
}
