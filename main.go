package main

import (
	"fmt"
	"log/slog"

	"github.com/Exam4/ApiGatawey/api"
	"github.com/Exam4/ApiGatawey/api/handler"
	"github.com/Exam4/ApiGatawey/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := config.Load()

	healthconn, err := grpc.NewClient(fmt.Sprintf("medical%s", config.MEDICAL_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		slog.Error("Error while connecting to grpc server" + err.Error())
		return
	}
	defer healthconn.Close()

	handler := handler.NewHandler(healthconn, nil)

	r := api.NewGin(handler)

	slog.Info("Starting server on port " + config.API_GETEWAY + "...")
	if err := r.Run(config.API_GETEWAY); err != nil {
		slog.Error(err.Error())
		return
	}
}
