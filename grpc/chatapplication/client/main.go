package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	pb "chatservice/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect:", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)
	ctx := context.Background()

	stream, err := client.ChatRoom(ctx)
	if err != nil {
		log.Fatal("failed to connect to chat room:", err)
	}

	var userId int32
	var roomId int32
	fmt.Print("Enter your user ID: ")
	fmt.Scanln(&userId)
	fmt.Print("Enter room ID to join: ")
	fmt.Scanln(&roomId)

	err = stream.Send(&pb.ChatRequest{
		SenderId: userId,
		RoomId:   roomId,
		Text:     "[joined the room]",
	})
	if err != nil {
		log.Fatal("failed to send join message:", err)
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				log.Println("receive error:", err)
				return
			}
			fmt.Printf("\nUser %d: %s\n> ", res.SenderId, res.Text)
		}
	}()

	fmt.Println("Use /private <receiverId> <message> or /leave to exit.")

	for {
		fmt.Print("> ")
		var input1, input2, input3 string
		fmt.Scanln(&input1, &input2, &input3)

		if input1 == "/leave" {
			_, err := client.LeaveRoom(ctx, &pb.LeaveRoomRequest{
				UserId: userId,
				RoomId: roomId,
			})
			if err != nil {
				log.Println("failed to leave room:", err)
			} else {
				fmt.Println("You left the room.")
			}
			stream.CloseSend()
			return
		}

		if input1 == "/private" {
			var receiverId int32
			var message string

			_, err := fmt.Sscan(input2, &receiverId)
			if err != nil {
				fmt.Println("Invalid receiver ID")
				continue
			}

			message = input3
			resp, err := client.SendPrivateMessage(ctx, &pb.MessageRequest{
				SenderId:   userId,
				ReceiverId: receiverId,
				Text:       message,
			})
			if err != nil {
				log.Println("Private message error:", err)
			} else {
				fmt.Printf("Private message status: %s\n", resp.Status)
			}
			continue
		}

		fullMessage := strings.Join([]string{input1, input2, input3}, " ")
		err := stream.Send(&pb.ChatRequest{
			SenderId: userId,
			RoomId:   roomId,
			Text:     fullMessage,
		})
		if err != nil {
			log.Println("send error:", err)
			break
		}
	}
}
