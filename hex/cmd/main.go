package main

import (
	"log"
	"os"

	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/adapters/app/api"
	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/adapters/core/arithmetic"
	gRPC "github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/adapters/framework/left/grpc"
	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/adapters/framework/right/db"
	"github.com/R014ND-D35CH41N/SWEasyHexArch/hex/cmd/internal/ports"
)

func main() {

	var err error

	//ports
	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
 
	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DB_NAME")

	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDBConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run() // This runs the application

	// Install the database container first and launch it
}
