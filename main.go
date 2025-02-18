package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	log "github.com/go-kit/log"
	grun "github.com/oklog/run"
	"github.com/savageking-io/necauth/endpoint"
	"github.com/savageking-io/necauth/pb"
	"github.com/savageking-io/necauth/service"
	"github.com/savageking-io/necauth/transport"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var (
		service    = service.NewService()
		endpoints  = endpoint.NewEndpoints(service)
		grpcServer = transport.NewRPCServer(endpoints)
	)

	var g grun.Group
	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}

	g.Add(func() error {
		logger.Log("transport", "gRPC", "addr", ":8080")
		baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
		pb.RegisterAuthServiceServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		grpcListener.Close()
	})

	interrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-interrupt:
			return nil
		}
	}, func(err error) {
		close(interrupt)
	})
	logger.Log("shutdown", g.Run())
}
