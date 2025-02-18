package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/savageking-io/necauth/endpoint"
	"github.com/savageking-io/necauth/pb"
	"github.com/savageking-io/necauth/service"
	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedAuthServiceServer
	credentials grpctransport.Handler
	token       grpctransport.Handler
}

func NewRPCServer(endpoints endpoint.Endpoints) pb.AuthServiceServer {
	return &grpcServer{
		credentials: grpctransport.NewServer(
			endpoints.AuthenticateCredentialsEndpoint,
			decodeGRPCCredentialsRequest,
			encodeGRPCCredentialsResponse,
		),
		token: grpctransport.NewServer(
			endpoints.AuthenticateTokenEndpoint,
			decodeGRPCTokenRequest,
			encodeGRPCTokenResponse,
		),
	}
}

func NewClient(conn *grpc.ClientConn) service.Service {
	var credentialsEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"AuthenticateCredentials",
		encodeClientGRPCCredentialsRequest,
		decodeClientGRPCCredentialsResponse,
		pb.AuthenticateCredentialsResponse{},
	).Endpoint()

	var tokenEndpoint = grpctransport.NewClient(
		conn,
		"pb.AuthService",
		"AuthenticateToken",
		encodeClientGRPCTokenRequest,
		decodeClientGRPCTokenResponse,
		pb.AuthenticateTokenResponse{},
	).Endpoint()

	return endpoint.Endpoints{
		AuthenticateCredentialsEndpoint: credentialsEndpoint,
		AuthenticateTokenEndpoint:       tokenEndpoint,
	}
}

func (s *grpcServer) AuthenticateCredentials(ctx context.Context, req *pb.AuthenticateCredentialsRequest) (*pb.AuthenticateCredentialsResponse, error) {
	_, rep, err := s.credentials.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AuthenticateCredentialsResponse), nil
}

func (s *grpcServer) AuthenticateToken(ctx context.Context, req *pb.AuthenticateTokenRequest) (*pb.AuthenticateTokenResponse, error) {
	_, rep, err := s.token.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AuthenticateTokenResponse), nil
}

func decodeGRPCCredentialsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AuthenticateCredentialsRequest)
	return endpoint.AuthenticateCredentialsRequest{Username: req.Username, Password: req.Password}, nil
}

func decodeGRPCTokenRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.AuthenticateTokenRequest)
	return endpoint.AuthenticateTokenRequest{Token: req.Token, Platform: req.Platform}, nil
}

func encodeGRPCCredentialsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.AuthenticateCredentialsResponse)
	return &pb.AuthenticateCredentialsResponse{Token: resp.Token, Error: resp.Error}, nil
}

func encodeGRPCTokenResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.AuthenticateTokenResponse)
	return &pb.AuthenticateTokenResponse{Token: resp.Token, Error: resp.Error}, nil
}

func encodeClientGRPCCredentialsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.AuthenticateCredentialsRequest)
	return &pb.AuthenticateCredentialsRequest{Username: req.Username, Password: req.Password}, nil
}

func decodeClientGRPCCredentialsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AuthenticateCredentialsResponse)
	return endpoint.AuthenticateCredentialsResponse{Token: reply.Token, Error: reply.Error}, nil
}

func encodeClientGRPCTokenRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint.AuthenticateTokenRequest)
	return &pb.AuthenticateTokenRequest{Token: req.Token, Platform: req.Platform}, nil
}

func decodeClientGRPCTokenResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.AuthenticateTokenResponse)
	return endpoint.AuthenticateTokenResponse{Token: reply.Token, Error: reply.Error}, nil
}
