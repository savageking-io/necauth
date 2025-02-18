package endpoint

import (
	"reflect"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/savageking-io/necauth/service"
)

func Test_makeAuthenticateCredentialsEndpoint(t *testing.T) {
	type args struct {
		s service.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAuthenticateCredentialsEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeAuthenticateCredentialsEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeAuthenticateTokenEndpoint(t *testing.T) {
	type args struct {
		s service.Service
	}
	tests := []struct {
		name string
		args args
		want endpoint.Endpoint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeAuthenticateTokenEndpoint(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeAuthenticateTokenEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
