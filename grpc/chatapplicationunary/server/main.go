package main

import (
	models "chatSystem/server/models"
	pb "chatSystem/server/protos"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type ChatServiceServer struct {
	pb.UnimplementedChatServiceServer
	mu sync.Mutex
}

var (
	users         = make(map[int]models.User)
	messages      = make(map[int]models.Message)
	lastUserId    int
	lastMessageId int
)

func getUserResponse(user models.User) *pb.User {
	return &pb.User{
		Id:     int32(user.Id),
		Name:   user.Name,
		Gender: user.Gender,
	}
}

func getMessageResponse(message models.Message) *pb.Message {
	return &pb.Message{
		Id:       int32(message.Id),
		Text:     message.Text,
		Sender:   getUserResponse(message.Sender),
		Receiver: getUserResponse(message.Receiver),
		Time:     message.Time,
	}
}

func main() {
	fmt.Println("Server starting")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
		return
	}
	fmt.Println("Listening")
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, &ChatServiceServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error serving: %v", err)
		return
	}
	fmt.Println("Server is running")
}

func (c *ChatServiceServer) Register(cts context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, user := range users {
		if user.Name == req.Name {
			return &pb.RegisterResponse{Status: 409}, errors.New("User name already exists")
		}
	}
	if req.Name == " " || req.Gender == "" {
		return &pb.RegisterResponse{Status: 409}, errors.New("Field cannot be empty")
	}
	lastUserId++
	user := models.User{
		Id:     lastUserId,
		Name:   req.Name,
		Gender: req.Gender,
	}
	users[lastUserId] = user
	return &pb.RegisterResponse{Status: 200}, nil
}

func (c *ChatServiceServer) GetUsers(cts context.Context, req *pb.EmptyRequest) (*pb.GetUserResponse, error) {
	modifiedUsers := make([]*pb.User, 0)
	for _, user := range users {
		modifiedUsers = append(modifiedUsers, getUserResponse(user))
	}
	return &pb.GetUserResponse{Users: modifiedUsers}, nil
}

func (c *ChatServiceServer) GetAllMessages(cts context.Context, req *pb.EmptyRequest) (*pb.GetMessageResponse, error) {
	modifiedMessages := make([]*pb.Message, 0)
	for _, message := range messages {
		modifiedMessages = append(modifiedMessages, getMessageResponse(message))
	}
	return &pb.GetMessageResponse{Messages: modifiedMessages}, nil
}

func (c *ChatServiceServer) GetMyMessages(cts context.Context, req *pb.GetMessageRequest) (*pb.GetMessageResponse, error) {
	modifiedMessages := make([]*pb.Message, 0)
	for _, message := range messages {
		if message.Sender.Id == int(req.UserId) || message.Receiver.Id == int(req.UserId) {
			modifiedMessages = append(modifiedMessages, getMessageResponse(message))
		}
	}
	return &pb.GetMessageResponse{Messages: modifiedMessages}, nil
}

func (c *ChatServiceServer) SendMessage(cts context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	if req.SenderId == req.ReceiverId {
		return &pb.MessageResponse{}, errors.New("Sender and receiver cannot be the same")
	}
	sender, senderExists := users[int(req.SenderId)]
	receiver, receiverExists := users[int(req.ReceiverId)]
	if !senderExists || !receiverExists {
		return &pb.MessageResponse{}, errors.New("Invalid user details")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	lastMessageId++
	message := models.Message{
		Id:       lastUserId,
		Sender:   sender,
		Receiver: receiver,
		Text:     req.Text,
	}
	messages[lastMessageId] = message
	return &pb.MessageResponse{Status: 200}, nil
}
