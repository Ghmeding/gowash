package services

import (
	"fmt"
)

func (s *Service) CreateUser(email string, pwd string, username string, address string) (string, error) {
	id, err := s.repo.CreateUser(email, pwd, username, address)
	if err != nil {
		return "", fmt.Errorf("could not create user")
	}
	return id, nil
}
