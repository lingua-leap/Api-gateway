package main

import (
	"api-gateway/api"
	"api-gateway/api/handler"
	"api-gateway/config"
	"api-gateway/generated/user"
	"api-gateway/pkg/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	logger := logs.InitLogger()
	cnf := config.Load()

	client, err := grpc.NewClient(cnf.HOST+cnf.GRPC_USER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("error in connecting to user service", "error", err)
		log.Fatal(err)
	}

	userClient := user.NewUserServiceClient(client)

	hd := handler.NewHandler(userClient, logger)

	router := api.NewRouter(hd)

	router.Init()
	log.Fatal(router.Run())
}
