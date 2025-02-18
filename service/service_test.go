package service

import (
	"context"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AuthenticateCredentials(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.AuthenticateCredentials(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.AuthenticateCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.AuthenticateCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AuthenticateToken(t *testing.T) {
	type args struct {
		ctx      context.Context
		token    string
		platform string
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.AuthenticateToken(tt.args.ctx, tt.args.token, tt.args.platform)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.AuthenticateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.AuthenticateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
