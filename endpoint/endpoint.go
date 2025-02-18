package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/savageking-io/necauth/service"
)

type Endpoints struct {
	AuthenticateCredentialsEndpoint endpoint.Endpoint
	AuthenticateTokenEndpoint       endpoint.Endpoint
}

type AuthenticateCredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticateCredentialsResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

type AuthenticateTokenRequest struct {
	Token    string `json:"token"`
	Platform string `json:"platform"`
}

type AuthenticateTokenResponse struct {
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func NewEndpoints(s service.Service) Endpoints {
	return Endpoints{
		AuthenticateCredentialsEndpoint: makeAuthenticateCredentialsEndpoint(s),
		AuthenticateTokenEndpoint:       makeAuthenticateTokenEndpoint(s),
	}
}

func (s Endpoints) AuthenticateCredentials(ctx context.Context, username, password string) (string, error) {
	request := AuthenticateCredentialsRequest{Username: username, Password: password}
	response, err := s.AuthenticateCredentialsEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	resp := response.(AuthenticateCredentialsResponse)
	return resp.Token, nil
}

func (s Endpoints) AuthenticateToken(ctx context.Context, token, platform string) (string, error) {
	request := AuthenticateTokenRequest{Token: token, Platform: platform}
	response, err := s.AuthenticateTokenEndpoint(ctx, request)
	if err != nil {
		return "", err
	}
	resp := response.(AuthenticateTokenResponse)
	return resp.Token, nil
}

func makeAuthenticateCredentialsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthenticateCredentialsRequest)
		token, err := s.AuthenticateCredentials(ctx, req.Username, req.Password)
		if err != nil {
			return AuthenticateCredentialsResponse{Token: token, Error: err.Error()}, nil
		}
		return AuthenticateCredentialsResponse{Token: token, Error: ""}, nil
	}
}

func makeAuthenticateTokenEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthenticateTokenRequest)
		token, err := s.AuthenticateToken(ctx, req.Token, req.Platform)
		if err != nil {
			return AuthenticateTokenResponse{Token: token, Error: err.Error()}, nil
		}
		return AuthenticateTokenResponse{Token: token, Error: ""}, nil
	}
}
