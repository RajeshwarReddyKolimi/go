package main

import (
	"context"
	"log"
	"time"

	pb "chatservice/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	ctx := context.Background()

	user1 := registerUser(ctx, client, "Alice")
	user2 := registerUser(ctx, client, "Bob")
	user3 := registerUser(ctx, client, "Charlie")

	room1 := createRoom(ctx, client, "GoChat")
	room2 := createRoom(ctx, client, "DevTalk")

	go joinRoom(ctx, client, user1, room1)
	go joinRoom(ctx, client, user2, room1)
	go joinRoom(ctx, client, user3, room1)

	go joinRoom(ctx, client, user1, room2)
	go joinRoom(ctx, client, user2, room2)
	go joinRoom(ctx, client, user3, room2)

	time.Sleep(1 * time.Second)

	sendPrivateMessage(ctx, client, user1, user2)
	sendPrivateMessage(ctx, client, user2, user3)
	sendPrivateMessage(ctx, client, user3, user1)

	go chatRoom(ctx, client, user1, room1, "Hi everyone in GoChat!")
	go chatRoom(ctx, client, user2, room1, "Hey Alice, how's Go?")
	go chatRoom(ctx, client, user3, room1, "Just joined GoChat, hello!")

	go chatRoom(ctx, client, user1, room2, "Welcome to DevTalk!")
	go chatRoom(ctx, client, user2, room2, "DevTalk rocks!")
	go chatRoom(ctx, client, user3, room2, "Love discussing dev topics here.")

	time.Sleep(5 * time.Second)
}

func registerUser(ctx context.Context, client pb.ChatServiceClient, name string) int32 {
	res, err := client.Register(ctx, &pb.RegisterRequest{Name: name})
	if err != nil {
		log.Fatal("register failed:", err)
	}
	log.Println("Registered:", res.UserId)
	return int32(len(name))
}

func createRoom(ctx context.Context, client pb.ChatServiceClient, name string) int32 {
	res, err := client.CreateRoom(ctx, &pb.CreateRoomRequest{
		Name:        name,
		Description: "Test room",
	})
	if err != nil {
		log.Fatal("create room failed:", err)
	}
	log.Println("Room created:", name, "Room ID:", res.RoomId)
	return res.RoomId
}

func sendPrivateMessage(ctx context.Context, client pb.ChatServiceClient, senderId, receiverId int32) {
	_, err := client.SendMessage(ctx, &pb.MessageRequest{
		SenderId:   senderId,
		ReceiverId: receiverId,
		Text:       "Hey Bob, this is a private message!",
	})
	if err != nil {
		log.Println("Failed to send private message:", err)
		return
	}
	log.Println("Private message sent from", senderId, "to", receiverId)
}

func joinRoom(ctx context.Context, client pb.ChatServiceClient, userId, roomId int32) {
	stream, err := client.JoinRoom(ctx, &pb.JoinRoomRequest{
		UserId: userId,
		RoomId: roomId,
	})
	if err != nil {
		log.Println("JoinRoom failed:", err)
		return
	}
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println("JoinRoom stream closed for user", userId, ":", err)
			return
		}
		log.Println("[Room]", "User", userId, "received:", msg.Text)
	}
}

func chatRoom(ctx context.Context, client pb.ChatServiceClient, userId, roomId int32, message string) {
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Println("Chat stream failed:", err)
		return
	}

	err = stream.Send(&pb.ChatRequest{
		SenderId: userId,
		RoomId:   roomId,
		Text:     message,
	})
	if err != nil {
		log.Println("Send chat failed:", err)
		return
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Println("Chat receive error for user", userId, ":", err)
				return
			}
			log.Println("[Chat]", "User", userId, "sees:", res.Text)
		}
	}()
}
