package pb

//go:generate protoc auth.proto --go_out=. --go-grpc_out=. --go_opt=Mauth.proto=github.com/savageking-io/necauth/pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_opt=Mauth.proto=github.com/savageking-io/necauth/pb
