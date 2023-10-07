package handler

import (
	"context"
	"log"
	pb "server/proto"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type messengerRequest struct {
	pb.UnimplementedMessengerServer
	requests []*pb.MessageRequest
}

func RegisterMessengerServer(s *grpc.Server) {
	pb.RegisterMessengerServer(s, &messengerRequest{})
}

func (s *messengerRequest) GetMessages(_ *empty.Empty, stream pb.Messenger_GetMessagesServer) error {
	for _, r := range s.requests {
		if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
			return err
		}
	}

	previousCount := len(s.requests)

	for {
		currentCount := len(s.requests)
		if previousCount < currentCount && currentCount > 0 {
			r := s.requests[currentCount-1]
			log.Printf("Sent: %v", r.GetMessage())
			if err := stream.Send(&pb.MessageResponse{Message: r.GetMessage()}); err != nil {
				return err
			}
		}
		previousCount = currentCount
	}
}

func (s *messengerRequest) CreateMessage(ctx context.Context, r *pb.MessageRequest) (*pb.MessageResponse, error) {
	newR := &pb.MessageRequest{Message: r.GetMessage() + ": " + time.Now().Format("2006-01-02 15:04:05")}
	s.requests = append(s.requests, newR)
	return &pb.MessageResponse{Message: r.GetMessage()}, nil
}
