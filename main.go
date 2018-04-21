package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/BurntSushi/toml"
	"github.com/devishot/grpc-go-time_tracking/infrastructure/database"
	"github.com/devishot/grpc-go-time_tracking/interface/api"
	"github.com/devishot/grpc-go-time_tracking/interface/handler"
	"github.com/devishot/grpc-go-time_tracking/interface/repository"
)

type tomlConfig struct {
	Port     int
	Database database.Config
}

func main() {
	cfg, err := readConfigs("config/dev.toml")
	checkError(err)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	checkError(err)

	// init repositories
	db, err := database.New(cfg.Database)
	checkError(err)
	defer db.Close()

	tr, err := repository.NewTimeRecordRepository(db)
	checkError(err)

	// init gRPC server
	handler := &handler.Server{TimeRecordRepository: tr}
	grpcServer := grpc.NewServer()

	// register gRPC handler
	api.RegisterTimeTrackingServer(grpcServer, handler)

	// start gRPC server
	err = grpcServer.Serve(lis)
	checkError(err)
}

func readConfigs(filepath string) (tomlConfig, error) {
	var cfg tomlConfig
	_, err := toml.DecodeFile(filepath, &cfg)
	return cfg, err
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(-1)
	}
}
