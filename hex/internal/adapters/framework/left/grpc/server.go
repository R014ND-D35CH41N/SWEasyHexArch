package rpc

import (
	"log"
	"net"

	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/adapters/framework/left/grpc/pb"
	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to lusten on port 9000: %v", err)
	}

	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
