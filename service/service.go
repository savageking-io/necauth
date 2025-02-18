package service

import (
	"context"
	"errors"
)

type Service interface {
	AuthenticateCredentials(ctx context.Context, username, password string) (string, error)
	AuthenticateToken(ctx context.Context, token, platform string) (string, error)
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidToken       = errors.New("invalid token")
)

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) AuthenticateCredentials(ctx context.Context, username, password string) (string, error) {
	if username == "admin" && password == "admin" {
		return "token", nil
	}
	return "", ErrInvalidCredentials
}

func (s *service) AuthenticateToken(ctx context.Context, token, platform string) (string, error) {
	if token == "token" && platform == "platform" {
		return "user", nil
	}
	return "", ErrInvalidToken
}
