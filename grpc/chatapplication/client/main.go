package main

import (
	pb "chatSystem/client/protos"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterUser(ctx context.Context, client pb.ChatServiceClient, name string, gender string) {
	_, err := client.Register(ctx, &pb.RegisterRequest{Name: name, Gender: gender})
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("User registered")
}

func GetUsers(ctx context.Context, client pb.ChatServiceClient) {
	users, usersError := client.GetUsers(ctx, &pb.EmptyRequest{})
	if usersError != nil {
		fmt.Println("Error", usersError)
	}
	for _, user := range users.Users {
		fmt.Println(user)
	}
}

func SendMessage(ctx context.Context, client pb.ChatServiceClient, senderId int, receiverId int, text string) {
	_, messageError := client.SendMessage(ctx, &pb.MessageRequest{SenderId: int32(senderId), ReceiverId: int32(receiverId), Text: text})
	if messageError != nil {
		fmt.Println("Error sending message", messageError)
	}
	fmt.Println("Message sent successfully")
}

func GetAllMessages(ctx context.Context, client pb.ChatServiceClient) {
	messages, messagesError := client.GetAllMessages(ctx, &pb.EmptyRequest{})
	if messagesError != nil {
		fmt.Println("Error", messagesError)
	}
	fmt.Println("All messages")
	for _, message := range messages.Messages {
		fmt.Println(message)
	}
}

func GetMyMessages(ctx context.Context, client pb.ChatServiceClient, userId int) {
	messages, messagesError := client.GetMyMessages(ctx, &pb.GetMessageRequest{UserId: int32(userId)})
	if messagesError != nil {
		fmt.Println("Error", messagesError)
	}
	fmt.Println("My messages")
	for _, message := range messages.Messages {
		fmt.Println(message)
	}
}

func createClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error creating client", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// RegisterUser(ctx, client, "Ragu", "male")
	// RegisterUser(ctx, client, "Ramu", "male")
	// RegisterUser(ctx, client, "Raju", "male")
	// RegisterUser(ctx, client, "Rani", "female")

	GetUsers(ctx, client)

	SendMessage(ctx, client, 1, 2, "1-2")
	SendMessage(ctx, client, 2, 3, "2-3")
	SendMessage(ctx, client, 1, 3, "1-3")
	SendMessage(ctx, client, 3, 1, "3-1")

	GetAllMessages(ctx, client)
	GetMyMessages(ctx, client, 1)
}

func main() {
	createClient()
}
