package main

import (
	"log"
	"net"
	"os"
	"server/entity"
	"server/handler"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	entity.DBConnect()
	defer entity.Close()

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	handler.RegisterMessengerServer(s)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
